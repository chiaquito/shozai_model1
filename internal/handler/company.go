package handler

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	"shozai_model1/internal/usecase"
)

type company struct {
	uc usecase.Company
}

func NewCompany(u usecase.Usecase) *company {
	return &company{
		uc: u.Company,
	}
}

func (h *company) GetCompanyByID(c echo.Context) error {
	res, err := h.uc.GetCompanyByID(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *company) ListCompany(c echo.Context) error {
	res, err := h.uc.ListCompany(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *company) CreateCompany(c echo.Context) error {
	res, err := h.uc.CreateCompany(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *company) BulkCreateCompany(c echo.Context) error {
	res, err := h.uc.BulkCreateCompany(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
