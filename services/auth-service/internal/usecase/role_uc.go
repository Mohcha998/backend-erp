package usecase

import (
	"auth-service/internal/domain"
	"auth-service/internal/repository"
)

type RoleUsecase struct {
	repo repository.RoleRepository
}

func NewRoleUsecase(r repository.RoleRepository) *RoleUsecase {
	return &RoleUsecase{repo: r}
}

func (u *RoleUsecase) Create(role *domain.Role) error {
	return u.repo.Create(role)
}

func (u *RoleUsecase) GetAll() ([]domain.Role, error) {
	return u.repo.FindAll()
}

func (u *RoleUsecase) GetByID(id uint) (*domain.Role, error) {
	return u.repo.FindByID(id)
}

func (u *RoleUsecase) Update(role *domain.Role) error {
	return u.repo.Update(role)
}

func (u *RoleUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
