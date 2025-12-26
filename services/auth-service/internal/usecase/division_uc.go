package usecase

import (
	"auth-service/internal/domain"
	"auth-service/internal/repository"
)

type DivisionUsecase struct {
	repo repository.DivisionRepository
}

func NewDivisionUsecase(r repository.DivisionRepository) *DivisionUsecase {
	return &DivisionUsecase{repo: r}
}

func (u *DivisionUsecase) Create(d *domain.Division) error {
	return u.repo.Create(d)
}

func (u *DivisionUsecase) GetAll() ([]domain.Division, error) {
	return u.repo.FindAll()
}

func (u *DivisionUsecase) GetByID(id uint) (*domain.Division, error) {
	return u.repo.FindByID(id)
}

func (u *DivisionUsecase) Update(d *domain.Division) error {
	return u.repo.Update(d)
}

func (u *DivisionUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
