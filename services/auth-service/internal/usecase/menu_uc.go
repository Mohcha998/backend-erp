package usecase

import (
	"auth-service/internal/domain"
	"auth-service/internal/repository"
)

type MenuUsecase struct{ repo *repository.MenuRepository }

func NewMenuUsecase(r *repository.MenuRepository) *MenuUsecase {
	return &MenuUsecase{r}
}

func (u *MenuUsecase) Create(d *domain.Menu) error {
	return u.repo.Create(d)
}
func (u *MenuUsecase) GetAll() ([]domain.Menu, error) {
	return u.repo.FindAll()
}
