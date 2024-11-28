package handler

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	"shozai_model1/internal/pkg/context"
	"shozai_model1/internal/pkg/errors"
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
	ctx := context.NewContext(c)
	res, err := h.uc.GetCompanyByID(ctx)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *company) ListCompany(c echo.Context) error {
	ctx := context.NewContext(c)
	res, err := h.uc.ListCompany(ctx)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *company) CreateCompany(c echo.Context) error {
	ctx := context.NewContext(c)
	res, err := h.uc.CreateCompany(ctx)
	if err != nil {
		if errors.IsBusinessError(err) {
			return c.JSON(http.StatusInternalServerError, err)
		}
		if errors.IsSystemError(err) {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}
	return c.JSON(http.StatusOK, res)
}

func (h *company) BulkCreateCompany(c echo.Context) error {
	ctx := context.NewContext(c)
	res, err := h.uc.BulkCreateCompany(ctx)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *company) UpdateCompany(c echo.Context) error {
	ctx := context.NewContext(c)
	res, err := h.uc.UpdateCompany(ctx)
	if err != nil {
		if errors.IsOptimisticLockingError(err) {
			return c.JSON(http.StatusConflict, err)
		}
		if errors.IsBusinessError(err) {
			return c.JSON(http.StatusInternalServerError, err)
		}
		if errors.IsSystemError(err) {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}
	return c.JSON(http.StatusOK, res)
}
