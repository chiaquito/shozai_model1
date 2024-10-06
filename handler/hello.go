package handler

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

// ここ使ってない
// type Hello interface {
// 	MyMethod(c echo.Context) error
// }

type hello struct {
}

// func NewHello() Hello {
// 	return &hello{}
// }

func (h *hello) MyMethod(c echo.Context) error {

	return c.JSON(http.StatusOK, "Hello World")
}
