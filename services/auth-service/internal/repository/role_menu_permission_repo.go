package repository

import (
	"auth-service/internal/domain"

	"gorm.io/gorm"
)

type RoleMenuPermissionRepository struct{ db *gorm.DB }

func NewRoleMenuPermissionRepository(db *gorm.DB) *RoleMenuPermissionRepository {
	return &RoleMenuPermissionRepository{db}
}

func (r *RoleMenuPermissionRepository) Create(p *domain.RoleMenuPermission) error {
	return r.db.Create(p).Error
}

func (r *RoleMenuPermissionRepository) FindAll() ([]domain.RoleMenuPermission, error) {
	var data []domain.RoleMenuPermission
	return data, r.db.Find(&data).Error
}
