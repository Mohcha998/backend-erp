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

func (u *UserUsecase) Create(user *domain.User) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		return u.repo.Create(tx, user)
	})
}

func (u *UserUsecase) GetAll() ([]domain.User, error) {
	return u.repo.FindAll()
}

func (u *UserUsecase) GetByID(id uint) (*domain.User, error) {
	return u.repo.FindByID(id)
}

func (u *UserUsecase) Update(user *domain.User) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		return u.repo.Update(tx, user)
	})
}

func (u *UserUsecase) Delete(id uint) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		return u.repo.Delete(tx, id)
	})
}
