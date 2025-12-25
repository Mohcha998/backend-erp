package repository

import (
	"auth-service/internal/domain"

	"gorm.io/gorm"
)

type UserRoleRepository interface {
	Create(tx *gorm.DB, ur *domain.UserRole) error
}

type userRoleRepo struct{}

func NewUserRoleRepository() UserRoleRepository {
	return &userRoleRepo{}
}

func (r *userRoleRepo) Create(tx *gorm.DB, ur *domain.UserRole) error {
	return tx.Create(ur).Error
}
