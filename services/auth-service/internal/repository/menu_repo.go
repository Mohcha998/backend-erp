package repository

import (
	"auth-service/internal/domain"

	"gorm.io/gorm"
)

type MenuRepository struct{ db *gorm.DB }

func NewMenuRepository(db *gorm.DB) *MenuRepository {
	return &MenuRepository{db}
}

func (r *MenuRepository) Create(d *domain.Menu) error {
	return r.db.Create(d).Error
}
func (r *MenuRepository) FindAll() ([]domain.Menu, error) {
	var data []domain.Menu
	return data, r.db.Find(&data).Error
}
