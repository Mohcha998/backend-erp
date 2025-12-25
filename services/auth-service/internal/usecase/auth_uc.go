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

/* ================= INTERFACE ================= */

type AuthUsecase interface {
	Login(email, password string) (string, error)
	Register(name, email, password string, divisionID, roleID uint) error
	Logout(token string)
	ForgotPassword(email string) (string, error)
}

/* ================= STRUCT ================= */

type authUsecase struct {
	db           *gorm.DB
	userRepo     repository.UserRepository
	userRoleRepo repository.UserRoleRepository
	jwtKey       string
}

/* ================= TOKEN BLACKLIST ================= */
var tokenBlacklist sync.Map

/* ================= CONSTRUCTOR ================= */

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

func (u *authUsecase) Login(email, password string) (string, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// ✅ bcrypt compare
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	); err != nil {
		return "", errors.New("invalid email or password")
	}

	// ambil role
	roles := []string{}
	for _, r := range user.Roles {
		roles = append(roles, r.Name)
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"roles":   roles,
		"exp":     time.Now().Add(8 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(u.jwtKey))
}

func (u *authUsecase) Register(
	name, email, password string,
	divisionID, roleID uint,
) error {

	// cek email
	if _, err := u.userRepo.FindByEmail(email); err == nil {
		return errors.New("email already registered")
	}

	// bcrypt hash
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	return u.db.Transaction(func(tx *gorm.DB) error {
		user := &domain.User{
			Name:       name,
			Email:      email,
			Password:   string(hashedPassword),
			DivisionID: divisionID,
		}

		if err := u.userRepo.Create(tx, user); err != nil {
			return err
		}

		return u.userRoleRepo.Create(tx, &domain.UserRole{
			UserID: user.ID,
			RoleID: roleID,
		})
	})
}

func (u *authUsecase) Logout(token string) {
	// blacklist token
	tokenBlacklist.Store(token, true)
}

func (u *authUsecase) ForgotPassword(email string) (string, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("email not found")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(u.jwtKey))
}
