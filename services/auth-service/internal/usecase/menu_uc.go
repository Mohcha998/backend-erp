package usecase

import (
	"auth-service/internal/domain"
	"auth-service/internal/repository"
)

type MenuUsecase struct {
	repo repository.MenuRepository
}

func NewMenuUsecase(r repository.MenuRepository) *MenuUsecase {
	return &MenuUsecase{repo: r}
}

func (u *MenuUsecase) Create(menu *domain.Menu) error {
	return u.repo.Create(menu)
}

func (u *MenuUsecase) GetAll() ([]domain.Menu, error) {
	return u.repo.FindAll()
}

func (u *MenuUsecase) GetByID(id uint) (*domain.Menu, error) {
	return u.repo.FindByID(id)
}

func (u *MenuUsecase) Update(menu *domain.Menu) error {
	return u.repo.Update(menu)
}

func (u *MenuUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
