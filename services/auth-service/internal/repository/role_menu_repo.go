package repository

import (
	"auth-service/internal/domain"

	"gorm.io/gorm"
)

type RoleMenuRepository interface {
	Assign(roleID uint, menuIDs []uint) error
}

type roleMenuRepo struct {
	db *gorm.DB
}

func NewRoleMenuRepository(db *gorm.DB) RoleMenuRepository {
	return &roleMenuRepo{db}
}

func (r *roleMenuRepo) Assign(roleID uint, menuIDs []uint) error {
	var menus []domain.Menu
	r.db.Where("id IN ?", menuIDs).Find(&menus)

	return r.db.Model(&domain.Role{ID: roleID}).
		Association("Menus").
		Replace(&menus)
}
