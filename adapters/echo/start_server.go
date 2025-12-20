package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
    DEFAULT_PORT = "8123"
)

func StartServer (env map[string]string) error {
    var address string

    if envPort, exists := env["SERVER_PORT"]; exists {
        address = ":" + envPort
    } else {
        address = ":" + DEFAULT_PORT
    }

    var e *echo.Echo = echo.New()

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

    return e.Start(address)
}
