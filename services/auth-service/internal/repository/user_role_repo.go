package repository

import (
	"auth-service/internal/domain"

	"gorm.io/gorm"
)

type UserRoleRepository interface {
	Create(tx *gorm.DB, userRole *domain.UserRole) error
}

type userRoleRepo struct {
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) UserRoleRepository {
	return &userRoleRepo{db: db}
}

// Create method untuk menambahkan user role
func (r *userRoleRepo) Create(tx *gorm.DB, userRole *domain.UserRole) error {
	return tx.Create(userRole).Error
}
