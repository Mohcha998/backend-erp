package repository

import (
	"errors"

	"auth-service/internal/domain"

	"gorm.io/gorm"
)

type DivisionRepository interface {
	Create(*domain.Division) error
	FindAll() ([]domain.Division, error)
	FindByID(uint) (*domain.Division, error)
	Update(*domain.Division) error
	Delete(uint) error
}

type divisionRepo struct {
	db *gorm.DB
}

func NewDivisionRepository(db *gorm.DB) DivisionRepository {
	return &divisionRepo{db}
}

func (r *divisionRepo) Create(d *domain.Division) error {
	return r.db.Create(d).Error
}

func (r *divisionRepo) FindAll() ([]domain.Division, error) {
	var data []domain.Division
	err := r.db.
		Preload("Roles").
		Preload("Roles.Menus").
		Find(&data).Error
	return data, err
}

func (r *divisionRepo) FindByID(id uint) (*domain.Division, error) {
	var data domain.Division
	err := r.db.
		Preload("Roles").
		Preload("Roles.Menus").
		First(&data, id).Error
	return &data, err
}

func (r *divisionRepo) Update(d *domain.Division) error {
	if d.ID == 0 {
		return errors.New("division id required")
	}

	return r.db.Model(&domain.Division{}).
		Where("id = ?", d.ID).
		Updates(map[string]interface{}{
			"name": d.Name,
		}).Error
}

func (r *divisionRepo) Delete(id uint) error {
	if id == 0 {
		return errors.New("division id required")
	}

	// soft delete
	return r.db.Delete(&domain.Division{}, id).Error
}
