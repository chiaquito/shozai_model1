package handler

import "shozai_model1/internal/usecase"

type Handler struct {
	Health  *health
	Company *company
	Product *product
}

func NewHandler(uc *usecase.Usecase) *Handler {

	health := &health{}
	company := &company{
		uc: uc.Company,
	}
	product := &product{
		uc: uc.Product,
	}

	return &Handler{
		Health:  health,
		Company: company,
		Product: product,
	}
}
