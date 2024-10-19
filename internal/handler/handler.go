package handler

import "shozai_model1/internal/usecase"

type Handler struct {
	Health  *health
	Company *company
}

func NewHandler(uc *usecase.Usecase) *Handler {

	health := &health{}
	company := &company{
		uc: uc.Company,
	}

	return &Handler{
		Health:  health,
		Company: company,
	}
}
