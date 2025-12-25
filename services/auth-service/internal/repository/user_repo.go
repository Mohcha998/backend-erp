package repository

import (
	"errors"

	"auth-service/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*domain.User, error)
	FindAll() ([]domain.User, error)
	FindByID(id uint) (*domain.User, error)
	Create(tx *gorm.DB, user *domain.User) error
	Update(tx *gorm.DB, user *domain.User) error
	Delete(tx *gorm.DB, id uint) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.
		Preload("Roles").
		Preload("Division").
		Where("email = ?", email).
		First(&user).Error
	return &user, err
}

func (r *userRepo) FindAll() ([]domain.User, error) {
	var users []domain.User
	err := r.db.
		Preload("Roles").
		Preload("Division").
		Find(&users).Error
	return users, err
}

func (r *userRepo) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.
		Preload("Roles").
		Preload("Division").
		First(&user, id).Error
	return &user, err
}

func (r *userRepo) Create(tx *gorm.DB, user *domain.User) error {
	return tx.Create(user).Error
}

func (r *userRepo) Update(tx *gorm.DB, user *domain.User) error {
	if user.ID == 0 {
		return errors.New("user id required")
	}

	return tx.Model(&domain.User{}).
		Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"name":        user.Name,
			"email":       user.Email,
			"is_active":   user.IsActive,
			"division_id": user.DivisionID,
		}).Error
}

func (r *userRepo) Delete(tx *gorm.DB, id uint) error {
	if id == 0 {
		return errors.New("user id required")
	}

	// soft delete
	return tx.Delete(&domain.User{}, id).Error
}
