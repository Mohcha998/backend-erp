package usecase

import (
	"auth-service/internal/repository"
	"auth-service/internal/pkg/apperror"
)

type PermissionUsecase struct {
	repo repository.RolePermissionRepository
}

func NewPermissionUsecase(r repository.RolePermissionRepository) *PermissionUsecase {
	return &PermissionUsecase{repo: r}
}

func (u *PermissionUsecase) Assign(roleID uint, permissionIDs []uint) error {
	// Assign permissions and handle any errors that may occur
	if err := u.repo.Assign(roleID, permissionIDs); err != nil {
		return apperror.ErrInternal // Handle error during permission assignment
	}
	return nil
}
