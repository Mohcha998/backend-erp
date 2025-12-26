package usecase

import (
	// "errors"
	"time"

	"auth-service/internal/domain"
	"auth-service/internal/repository"
	"auth-service/internal/pkg/apperror"
	"auth-service/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ================= INTERFACE =================
type AuthUsecase interface {
	Login(email, password string) (accessToken string, refreshToken string, err error)
	RefreshToken(refreshToken string) (string, string, error)
	Register(name, email, password string, divisionID, roleID uint) error
	Logout(refreshToken string)
	ForgotPassword(email string) (string, error)

	// ✅ wajib untuk Auto Refresh Middleware
	ValidateAccessToken(token string) (*jwt.Token, error)
}

// ================= STRUCT =================
type authUsecase struct {
	db           *gorm.DB
	userRepo     repository.UserRepository
	userRoleRepo repository.UserRoleRepository
	refreshRepo  repository.RefreshTokenRepository
	jwtKey       string
}

// ================= CONSTRUCTOR =================
func NewAuthUsecase(
	db *gorm.DB,
	userRepo repository.UserRepository,
	userRoleRepo repository.UserRoleRepository,
	refreshRepo repository.RefreshTokenRepository,
	jwtKey string,
) AuthUsecase {
	return &authUsecase{
		db:           db,
		userRepo:     userRepo,
		userRoleRepo: userRoleRepo,
		refreshRepo:  refreshRepo,
		jwtKey:       jwtKey,
	}
}

// =====================================================
// LOGIN
// =====================================================
func (u *authUsecase) Login(email, password string) (string, string, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return "", "", apperror.ErrNotFound // Using apperror for consistent error handling
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	); err != nil {
		return "", "", apperror.ErrUnauthorized // Unauthorized access error
	}

	// ===== ACCESS TOKEN =====
	accessClaims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	accessToken, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		accessClaims,
	).SignedString([]byte(u.jwtKey))
	if err != nil {
		return "", "", apperror.ErrInternal // Internal error for failed token generation
	}

	// ===== REFRESH TOKEN =====
	rawRefresh := uuid.NewString()
	hashed := utils.HashToken(rawRefresh)

	err = u.refreshRepo.Create(&domain.RefreshToken{
		UserID:    user.ID,
		TokenHash: hashed,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	})
	if err != nil {
		return "", "", apperror.ErrInternal // Error while storing refresh token
	}

	return accessToken, rawRefresh, nil
}

// =====================================================
// REFRESH TOKEN
// =====================================================
func (u *authUsecase) RefreshToken(oldRefresh string) (string, string, error) {
	hashed := utils.HashToken(oldRefresh)

	rt, err := u.refreshRepo.FindValid(hashed)
	if err != nil {
		return "", "", apperror.ErrUnauthorized // Invalid refresh token error
	}

	// revoke old token
	_ = u.refreshRepo.Revoke(hashed)

	user, err := u.userRepo.FindByID(rt.UserID)
	if err != nil {
		return "", "", apperror.ErrNotFound // User not found
	}

	// ===== NEW ACCESS TOKEN =====
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	newAccess, _ := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	).SignedString([]byte(u.jwtKey))

	// ===== NEW REFRESH TOKEN =====
	newRefresh := uuid.NewString()
	newHash := utils.HashToken(newRefresh)

	_ = u.refreshRepo.Create(&domain.RefreshToken{
		UserID:    user.ID,
		TokenHash: newHash,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	})

	return newAccess, newRefresh, nil
}

// =====================================================
// REGISTER
// =====================================================
func (u *authUsecase) Register(
	name, email, password string,
	divisionID, roleID uint,
) error {

	if _, err := u.userRepo.FindByEmail(email); err == nil {
		return apperror.ErrBadRequest // Bad request for duplicate email
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return apperror.ErrInternal // Internal error for password hashing failure
	}

	return u.db.Transaction(func(tx *gorm.DB) error {
		user := &domain.User{
			Name:       name,
			Email:      email,
			Password:   string(hash),
			DivisionID: divisionID,
			IsActive:   true,
		}

		if err := u.userRepo.Create(tx, user); err != nil {
			return apperror.ErrInternal // Error during user creation
		}

		return u.userRoleRepo.Create(tx, &domain.UserRole{
			UserID: user.ID,
			RoleID: roleID,
		})
	})
}

// =====================================================
// LOGOUT
// =====================================================
func (u *authUsecase) Logout(refreshToken string) {
	hashed := utils.HashToken(refreshToken)
	_ = u.refreshRepo.Revoke(hashed)
}

// =====================================================
// FORGOT PASSWORD
// =====================================================
func (u *authUsecase) ForgotPassword(email string) (string, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return "", apperror.ErrNotFound // Email not found error
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	return jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	).SignedString([]byte(u.jwtKey))
}

// =====================================================
// VALIDATE ACCESS TOKEN (AUTO REFRESH)
// =====================================================
func (u *authUsecase) ValidateAccessToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(u.jwtKey), nil
	})
}
