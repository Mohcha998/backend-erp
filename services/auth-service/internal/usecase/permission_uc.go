package usecase

import "auth-service/internal/repository"

type PermissionUsecase struct {
	repo repository.RolePermissionRepository
}

func NewPermissionUsecase(r repository.RolePermissionRepository) *PermissionUsecase {
	return &PermissionUsecase{repo: r}
}

func (u *PermissionUsecase) Assign(roleID uint, permissionIDs []uint) error {
	return u.repo.Assign(roleID, permissionIDs)
}
