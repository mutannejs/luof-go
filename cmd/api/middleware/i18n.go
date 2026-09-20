package middleware

import (
	"github.com/mutannejs/luof-go/cmd/api/custom"
	"github.com/mutannejs/luof-go/pkg/li18n"

	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func i18nMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var language = c.Request().Header.Get("Accept-Language")
			var bundle = li18n.GetBundle()

			cc := c.(*custom.Context)
			cc.Localizer = i18n.NewLocalizer(bundle, language)

			return next(cc)
		}
	}
}
