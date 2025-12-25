package usecase

import (
	"auth-service/internal/domain"
	"auth-service/internal/repository"
)

type PermissionUsecase struct {
	repo *repository.RoleMenuPermissionRepository
}

func NewPermissionUsecase(r *repository.RoleMenuPermissionRepository) *PermissionUsecase {
	return &PermissionUsecase{r}
}

func (u *PermissionUsecase) Create(p *domain.RoleMenuPermission) error {
	return u.repo.Create(p)
}
func (u *PermissionUsecase) GetAll() ([]domain.RoleMenuPermission, error) {
	return u.repo.FindAll()
}
