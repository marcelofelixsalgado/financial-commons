package middlewares

import (
	"github.com/marcelofelixsalgado/financial-commons/api/responses"
	"github.com/marcelofelixsalgado/financial-commons/api/responses/faults"
	"github.com/marcelofelixsalgado/financial-commons/pkg/commons/logger"

	"github.com/labstack/echo/v4"
	"github.com/marcelofelixsalgado/financial-commons/pkg/auth"
)

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := auth.ValidateToken(c.Request()); err != nil {
			logger.GetLogger().Infof("Token validation error: %v", err)
			responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.NotAuthorized)
			return c.JSON(responseMessage.HttpStatusCode, responseMessage)
		}
		return next(c)
	}
}
