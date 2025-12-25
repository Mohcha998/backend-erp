package usecase

import (
	"errors"
	"sync"
	"time"

	"auth-service/internal/domain"
	"auth-service/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthUsecase interface {
	Login(email, password string) (string, error)
	Register(name, email, password string, divisionID, roleID uint) error
	Logout(token string)
	ForgotPassword(email string) (string, error)
}

type authUsecase struct {
	db           *gorm.DB
	userRepo     repository.UserRepository
	userRoleRepo repository.UserRoleRepository
	jwtKey       string
}

var tokenBlacklist sync.Map

func NewAuthUsecase(
	db *gorm.DB,
	userRepo repository.UserRepository,
	userRoleRepo repository.UserRoleRepository,
	jwtKey string,
) AuthUsecase {
	return &authUsecase{
		db:           db,
		userRepo:     userRepo,
		userRoleRepo: userRoleRepo,
		jwtKey:       jwtKey,
	}
}

/* ================= LOGIN ================= */
func (u *authUsecase) Login(email, password string) (string, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	); err != nil {
		return "", errors.New("invalid password")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID, // ✅ uint
		"email":   user.Email,
		"exp":     time.Now().Add(8 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(u.jwtKey))
}

/* ================= REGISTER ================= */
func (u *authUsecase) Register(
	name, email, password string,
	divisionID, roleID uint,
) error {
	_, err := u.userRepo.FindByEmail(email)
	if err == nil {
		return errors.New("email already registered")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return u.db.Transaction(func(tx *gorm.DB) error {
		user := &domain.User{
			Name:       name,
			Email:      email,
			Password:   string(hash),
			DivisionID: divisionID,
		}

		if err := u.userRepo.Create(tx, user); err != nil {
			return err
		}

		userRole := &domain.UserRole{
			UserID: user.ID, // ✅ uint
			RoleID: roleID,
		}

		if err := u.userRoleRepo.Create(tx, userRole); err != nil {
			return err
		}

		return nil
	})
}

/* ================= LOGOUT ================= */
func (u *authUsecase) Logout(token string) {
	tokenBlacklist.Store(token, true)
}

/* ================= FORGOT PASSWORD ================= */
func (u *authUsecase) ForgotPassword(email string) (string, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("email not found")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID, // ✅ uint
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(u.jwtKey))
}
