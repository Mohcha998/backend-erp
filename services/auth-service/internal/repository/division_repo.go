package repository

import (
	"auth-service/internal/domain"

	"gorm.io/gorm"
)

type DivisionRepository struct{ db *gorm.DB }

func NewDivisionRepository(db *gorm.DB) *DivisionRepository {
	return &DivisionRepository{db}
}

func (r *DivisionRepository) Create(d *domain.Division) error {
	return r.db.Create(d).Error
}
func (r *DivisionRepository) FindAll() ([]domain.Division, error) {
	var data []domain.Division
	return data, r.db.Find(&data).Error
}
func (r *DivisionRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Division{}, id).Error
}
