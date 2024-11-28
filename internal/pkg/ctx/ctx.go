package ctx

import (
	echo "github.com/labstack/echo/v4"
)

type context struct {
	echoCtx echo.Context
}

func NewContext(echoCtx echo.Context) *context {
	return &context{
		echoCtx: echoCtx,
	}
}

func (ctx *context) Validate(i interface{}) error {
	return ctx.echoCtx.Validate(i)
}

func (ctx *context) Bind(i interface{}) error {
	return ctx.echoCtx.Bind(i)
}

func (ctx *context) JSON(code int, i interface{}) error {
	return ctx.echoCtx.JSON(code, i)
}
