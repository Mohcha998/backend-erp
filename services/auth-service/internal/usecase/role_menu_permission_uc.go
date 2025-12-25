package usecase

import "auth-service/internal/repository"

type RoleMenuUsecase struct {
	repo repository.RoleMenuRepository
}

func NewRoleMenuUsecase(r repository.RoleMenuRepository) *RoleMenuUsecase {
	return &RoleMenuUsecase{repo: r}
}

func (u *RoleMenuUsecase) Assign(roleID uint, menuIDs []uint) error {
	return u.repo.Assign(roleID, menuIDs)
}
