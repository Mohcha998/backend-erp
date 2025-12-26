package usecase

import (
	"auth-service/internal/domain"
	"auth-service/internal/repository"
	"auth-service/internal/pkg/apperror"
)

type DivisionUsecase struct {
	repo repository.DivisionRepository
}

func NewDivisionUsecase(r repository.DivisionRepository) *DivisionUsecase {
	return &DivisionUsecase{repo: r}
}

func (u *DivisionUsecase) Create(d *domain.Division) error {
	if err := u.repo.Create(d); err != nil {
		return apperror.ErrInternal // Handle error if division creation fails
	}
	return nil
}

func (u *DivisionUsecase) GetAll() ([]domain.Division, error) {
	divisions, err := u.repo.FindAll()
	if err != nil {
		return nil, apperror.ErrInternal // Handle error if fetching all divisions fails
	}
	return divisions, nil
}

func (u *DivisionUsecase) GetByID(id uint) (*domain.Division, error) {
	division, err := u.repo.FindByID(id)
	if err != nil {
		return nil, apperror.ErrNotFound // Handle error if division not found
	}
	return division, nil
}

func (u *DivisionUsecase) Update(d *domain.Division) error {
	if err := u.repo.Update(d); err != nil {
		return apperror.ErrNotFound // Handle error if updating division fails (e.g., division not found)
	}
	return nil
}

func (u *DivisionUsecase) Delete(id uint) error {
	if err := u.repo.Delete(id); err != nil {
		return apperror.ErrNotFound // Handle error if deleting division fails (e.g., division not found)
	}
	return nil
}
