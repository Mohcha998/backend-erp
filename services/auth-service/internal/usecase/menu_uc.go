package usecase

import (
	"auth-service/internal/domain"
	"auth-service/internal/repository"
	"auth-service/internal/pkg/apperror"
)

type MenuUsecase struct {
	repo repository.MenuRepository
}

func NewMenuUsecase(r repository.MenuRepository) *MenuUsecase {
	return &MenuUsecase{repo: r}
}

func (u *MenuUsecase) Create(menu *domain.Menu) error {
	if err := u.repo.Create(menu); err != nil {
		return apperror.ErrInternal // Handle error if menu creation fails
	}
	return nil
}

func (u *MenuUsecase) GetAll() ([]domain.Menu, error) {
	menus, err := u.repo.FindAll()
	if err != nil {
		return nil, apperror.ErrInternal // Handle error if fetching all menus fails
	}
	return menus, nil
}

func (u *MenuUsecase) GetByID(id uint) (*domain.Menu, error) {
	menu, err := u.repo.FindByID(id)
	if err != nil {
		return nil, apperror.ErrNotFound // Handle error if menu not found
	}
	return menu, nil
}

func (u *MenuUsecase) Update(menu *domain.Menu) error {
	if err := u.repo.Update(menu); err != nil {
		return apperror.ErrNotFound // Handle error if updating menu fails (e.g., menu not found)
	}
	return nil
}

func (u *MenuUsecase) Delete(id uint) error {
	if err := u.repo.Delete(id); err != nil {
		return apperror.ErrNotFound // Handle error if deleting menu fails (e.g., menu not found)
	}
	return nil
}
