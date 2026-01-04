package config

import (
	"github.com/mutannejs/luof-go/cmd/api/types"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setMiddleware(e *echo.Echo, repositories repository.Repositories) {
    e.Pre(middleware.AddTrailingSlash())

    e.Use(customContextMiddleware(repositories))
}

func customContextMiddleware(repositories repository.Repositories) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            cc := &types.CustomContext{Context: c, Repositories: repositories}
            return next(cc)
        }
    }
}
