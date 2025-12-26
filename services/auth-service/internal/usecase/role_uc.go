package usecase

import (
	"auth-service/internal/domain"
	"auth-service/internal/repository"
	"auth-service/internal/pkg/apperror"
)

type RoleUsecase struct {
	repo repository.RoleRepository
}

func NewRoleUsecase(r repository.RoleRepository) *RoleUsecase {
	return &RoleUsecase{repo: r}
}

func (u *RoleUsecase) Create(role *domain.Role) error {
	if err := u.repo.Create(role); err != nil {
		return apperror.ErrInternal // Handle error if role creation fails
	}
	return nil
}

func (u *RoleUsecase) GetAll() ([]domain.Role, error) {
	roles, err := u.repo.FindAll()
	if err != nil {
		return nil, apperror.ErrInternal // Handle error if fetching all roles fails
	}
	return roles, nil
}

func (u *RoleUsecase) GetByID(id uint) (*domain.Role, error) {
	role, err := u.repo.FindByID(id)
	if err != nil {
		return nil, apperror.ErrNotFound // Handle error if role not found
	}
	return role, nil
}

func (u *RoleUsecase) Update(role *domain.Role) error {
	if err := u.repo.Update(role); err != nil {
		return apperror.ErrNotFound // Handle error if updating role fails (e.g., role not found)
	}
	return nil
}

func (u *RoleUsecase) Delete(id uint) error {
	if err := u.repo.Delete(id); err != nil {
		return apperror.ErrNotFound // Handle error if deleting role fails (e.g., role not found)
	}
	return nil
}
