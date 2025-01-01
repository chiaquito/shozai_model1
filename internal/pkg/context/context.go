package context

import (
	echo "github.com/labstack/echo/v4"
)

type Context struct {
	echoCtx echo.Context
}

func NewContext(echoCtx echo.Context) *Context {
	return &Context{
		echoCtx: echoCtx,
	}
}

func (ctx *Context) Validate(i interface{}) error {
	return ctx.echoCtx.Validate(i)
}

func (ctx *Context) Bind(i interface{}) error {
	return ctx.echoCtx.Bind(i)
}

func (ctx *Context) JSON(code int, i interface{}) error {
	return ctx.echoCtx.JSON(code, i)
}
