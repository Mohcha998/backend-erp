package repository

import (
	"auth-service/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*domain.User, error)
	Create(tx *gorm.DB, user *domain.User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepo) Create(tx *gorm.DB, user *domain.User) error {
	return tx.Create(user).Error
}
