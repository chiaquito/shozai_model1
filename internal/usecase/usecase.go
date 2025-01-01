package usecase

import "shozai_model1/internal/repository"

type Usecase struct {
	Company *company
	Product *product
}

func New(r *repository.Repository) *Usecase {
	company := NewCompany(r.Company)
	product := NewProduct(r.Product)
	return &Usecase{
		Company: company,
		Product: product,
	}
}
