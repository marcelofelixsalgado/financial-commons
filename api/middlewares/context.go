package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/marcelofelixsalgado/financial-commons/api/context"
)

func CustomContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &context.CustomContext{
			Context: c,
		}
		return next(cc)
	}
}
