package middleware

import (
	"github.com/mutannejs/luof-go/cmd/api/custom"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/labstack/echo/v4"
)

func contextMiddleware(repositories repository.Repositories) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &custom.Context{Context: c, Repositories: repositories}
			return next(cc)
		}
	}
}
