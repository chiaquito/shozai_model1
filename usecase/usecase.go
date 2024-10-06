package usecase

import "shozai_model1/repository"

type Usecase struct {
	Company *company
}

func New(r *repository.Repository) *Usecase {
	company := NewCompany(r.Company)
	return &Usecase{
		Company: company,
	}
}
