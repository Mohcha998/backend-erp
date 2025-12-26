package usecase

import (
	"auth-service/internal/domain"
	"auth-service/internal/pkg/apperror"
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
		if err := u.repo.Create(tx, user); err != nil {
			return apperror.ErrInternal
		}
		return nil
	})
}

func (u *UserUsecase) GetAll() ([]domain.User, error) {
	users, err := u.repo.FindAll()
	if err != nil {
		return nil, apperror.ErrInternal
	}
	return users, nil
}

func (u *UserUsecase) GetByID(id uint) (*domain.User, error) {
	user, err := u.repo.FindByID(id)
	if err != nil {
		return nil, apperror.ErrNotFound
	}
	return user, nil
}

func (u *UserUsecase) Update(user *domain.User) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		if err := u.repo.Update(tx, user); err != nil {
			return apperror.ErrNotFound
		}
		return nil
	})
}

func (u *UserUsecase) Delete(id uint) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		if err := u.repo.Delete(tx, id); err != nil {
			return apperror.ErrNotFound
		}
		return nil
	})
}
