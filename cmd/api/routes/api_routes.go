package routes

import (
	"github.com/mutannejs/luof-go/cmd/api/handlers"

	"github.com/labstack/echo/v4"
)

func SetRootRoutes(e *echo.Echo) {
	var api *echo.Group = e.Group("/api")
	setApiRoutes(api)
}

func setApiRoutes(api *echo.Group) {
	var links *echo.Group = api.Group("/links")
	setApiLinksRoutes(links)

	var categories *echo.Group = api.Group("/categories")
	setApiCategoriesRoutes(categories)

	api.GET("/", handlers.GetApi)
}
