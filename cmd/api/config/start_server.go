package config

import (
	"github.com/mutannejs/luof-go/cmd/api/types"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

const (
    DEFAULT_PORT = "8123"
)

func StartServer(env map[string]string, repositories repository.Repositories) error {
    var address string

    if envPort, exists := env["SERVER_PORT"]; exists {
        address = ":" + envPort
    } else {
        address = ":" + DEFAULT_PORT
    }

    var e *echo.Echo = echo.New()

    e.Validator = &types.CustomValidator{Validator: validator.New()}
    e.HideBanner = true

    setRootRoutes(e)
    setMiddleware(e, repositories)

    return e.Start(address)
}
