package handler

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type health struct {
}

func (h *health) HealthCheck(c echo.Context) error {

	return c.JSON(http.StatusOK, "healthcheck")
}
