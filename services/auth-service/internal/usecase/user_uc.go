package usecase

import (
	"auth-service/internal/domain"
	"auth-service/internal/repository"

	"gorm.io/gorm"
)

type UserUsecase struct {
	db   *gorm.DB
	repo repository.UserRepository
}

func NewUserUsecase(db *gorm.DB, repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		db:   db,
		repo: repo,
	}
}

// CREATE USER (ADMIN)
func (u *UserUsecase) Create(user *domain.User) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		return u.repo.Create(tx, user)
	})
}

// GET ALL USER
func (u *UserUsecase) GetAll() ([]domain.User, error) {
	var users []domain.User
	err := u.db.Find(&users).Error
	return users, err
}
