package handler

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type hello struct {
}

func (h *hello) MyMethod(c echo.Context) error {

	return c.JSON(http.StatusOK, "heathcheck")
}
