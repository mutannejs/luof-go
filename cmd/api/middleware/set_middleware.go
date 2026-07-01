package middleware

import (
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetMiddleware(e *echo.Echo, repositories repository.Repositories) {
	e.Pre(middleware.AddTrailingSlash())

	e.Use(contextMiddleware(repositories))
}
