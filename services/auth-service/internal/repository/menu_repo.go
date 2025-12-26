package repository

import (
	"errors"

	"auth-service/internal/domain"

	"gorm.io/gorm"
)

type MenuRepository interface {
	Create(*domain.Menu) error
	FindAll() ([]domain.Menu, error)
	FindByID(uint) (*domain.Menu, error)
	Update(*domain.Menu) error
	Delete(uint) error
}

type menuRepo struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepo{db}
}

func (r *menuRepo) Create(m *domain.Menu) error {
	return r.db.Create(m).Error
}

func (r *menuRepo) FindAll() ([]domain.Menu, error) {
	var data []domain.Menu
	err := r.db.Find(&data).Error
	return data, err
}

func (r *menuRepo) FindByID(id uint) (*domain.Menu, error) {
	var data domain.Menu
	err := r.db.First(&data, id).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *menuRepo) Update(m *domain.Menu) error {
	if m.ID == 0 {
		return errors.New("menu id required")
	}

	return r.db.Model(&domain.Menu{}).
		Where("id = ?", m.ID).
		Updates(map[string]interface{}{
			"code": m.Code,
			"name": m.Name,
			"path": m.Path,
		}).Error
}

func (r *menuRepo) Delete(id uint) error {
	return r.db.Delete(&domain.Menu{}, id).Error
}
