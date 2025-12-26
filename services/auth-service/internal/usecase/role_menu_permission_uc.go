package usecase

import (
	"auth-service/internal/repository"
	"auth-service/internal/pkg/apperror"
)

type RoleMenuUsecase struct {
	repo repository.RoleMenuRepository
}

func NewRoleMenuUsecase(r repository.RoleMenuRepository) *RoleMenuUsecase {
	return &RoleMenuUsecase{repo: r}
}

func (u *RoleMenuUsecase) Assign(roleID uint, menuIDs []uint) error {
	// Assign menu to the role and handle errors if any occur
	if err := u.repo.Assign(roleID, menuIDs); err != nil {
		return apperror.ErrInternal // Return a generic internal error if assignment fails
	}
	return nil
}
