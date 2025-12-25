package echo

import (
	"github.com/mutannejs/luof-go/adapters/echo/types"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

func setMiddleware(e *echo.Echo, repositories repository.Repositories) {
    e.Pre(middleware.AddTrailingSlash())

    e.Use(customContextMiddleware(repositories))

    e.Use(middleware.BodyDump(logRequestMiddleware))
}

func customContextMiddleware(repositories repository.Repositories) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            cc := &types.CustomContext{c, repositories}
            return next(cc)
        }
    }
}

func logRequestMiddleware(c echo.Context, reqBody, resBody []byte) {
    log.Info().
        Str("path", c.Request().URL.Path).
        Str("method", c.Request().Method).
        RawJSON("reqBody", reqBody).
        RawJSON("resBody", resBody).
        Send()
}
