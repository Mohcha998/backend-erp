package usecase

import (
	"auth-service/internal/domain"
	"auth-service/internal/repository"
)

type RoleUsecase struct{ repo *repository.RoleRepository }

func NewRoleUsecase(r *repository.RoleRepository) *RoleUsecase {
	return &RoleUsecase{r}
}

func (u *RoleUsecase) Create(d *domain.Role) error {
	return u.repo.Create(d)
}
func (u *RoleUsecase) GetAll() ([]domain.Role, error) {
	return u.repo.FindAll()
}
