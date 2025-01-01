package handler

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	"shozai_model1/internal/pkg/errors"
	"shozai_model1/internal/usecase"
)

type product struct {
	uc usecase.Product
}

func NewProduct(u usecase.Usecase) *product {
	return &product{
		uc: u.Product,
	}
}

func (h *product) GetProductByID(c echo.Context) error {
	res, err := h.uc.GetProductByID(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *product) ListProduct(c echo.Context) error {
	res, err := h.uc.ListProduct(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *product) ListProductRaw(c echo.Context) error {
	res, err := h.uc.ListProductRaw(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *product) CreateProduct(c echo.Context) error {
	res, err := h.uc.CreateProduct(c)
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

func (h *product) BulkCreateProduct(c echo.Context) error {
	res, err := h.uc.BulkCreateProduct(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *product) UpdateProduct(c echo.Context) error {
	res, err := h.uc.UpdateProduct(c)
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
