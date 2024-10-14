package handler

import "shozai_model1/internal/usecase"

type Handler struct {
	Hello   *hello
	Company *company
}

func NewHandler(uc *usecase.Usecase) *Handler {

	hello := &hello{}
	company := &company{
		uc: uc.Company,
	}

	return &Handler{
		Hello:   hello,
		Company: company,
	}
}
