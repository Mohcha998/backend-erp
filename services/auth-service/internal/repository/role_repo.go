package repository

import (
	"auth-service/internal/domain"

	"gorm.io/gorm"
)

type RoleRepository struct{ db *gorm.DB }

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db}
}

func (r *RoleRepository) Create(d *domain.Role) error {
	return r.db.Create(d).Error
}
func (r *RoleRepository) FindAll() ([]domain.Role, error) {
	var data []domain.Role
	return data, r.db.Find(&data).Error
}
