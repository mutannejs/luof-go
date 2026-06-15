package config

import (
	"github.com/mutannejs/luof-go/cmd/api/routes"
	"github.com/mutannejs/luof-go/core/repository"

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

	e.HideBanner = true

	setMiddleware(e, repositories)
	routes.SetRootRoutes(e)

	return e.Start(address)
}
