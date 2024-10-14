package usecase

import "shozai_model1/internal/repository"

type Usecase struct {
	Company *company
}

func New(r *repository.Repository) *Usecase {
	company := NewCompany(r.Company)
	return &Usecase{
		Company: company,
	}
}
