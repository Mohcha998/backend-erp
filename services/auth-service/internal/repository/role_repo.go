package repository

import (
	"errors"

	"auth-service/internal/domain"

	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(*domain.Role) error
	FindAll() ([]domain.Role, error)
	FindByID(id uint) (*domain.Role, error)
	Update(*domain.Role) error
	Delete(id uint) error
}

type roleRepo struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepo{db}
}

func (r *roleRepo) Create(role *domain.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepo) FindAll() ([]domain.Role, error) {
	var roles []domain.Role
	err := r.db.
		Preload("Menus").
		Preload("Permissions").
		Find(&roles).Error
	return roles, err
}

func (r *roleRepo) FindByID(id uint) (*domain.Role, error) {
	var role domain.Role
	err := r.db.
		Preload("Menus").
		Preload("Permissions").
		First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepo) Update(role *domain.Role) error {
	if role.ID == 0 {
		return errors.New("role id is required")
	}

	return r.db.Model(&domain.Role{}).
		Where("id = ?", role.ID).
		Updates(map[string]interface{}{
			"name": role.Name,
		}).Error
}

func (r *roleRepo) Delete(id uint) error {
	if id == 0 {
		return errors.New("role id is required")
	}

	return r.db.Delete(&domain.Role{}, id).Error
}
