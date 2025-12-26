package repository

import (
	"auth-service/internal/domain"

	"gorm.io/gorm"
)

type RolePermissionRepository interface {
	Assign(roleID uint, permissionIDs []uint) error
}

type rolePermissionRepo struct {
	db *gorm.DB
}

func NewRolePermissionRepository(db *gorm.DB) RolePermissionRepository {
	return &rolePermissionRepo{db}
}

func (r *rolePermissionRepo) Assign(roleID uint, permissionIDs []uint) error {
	var perms []domain.Permission
	r.db.Where("id IN ?", permissionIDs).Find(&perms)

	return r.db.Model(&domain.Role{ID: roleID}).
		Association("Permissions").
		Replace(&perms)
}
