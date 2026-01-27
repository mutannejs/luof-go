package config

import (
	"github.com/mutannejs/luof-go/cmd/api/handlers"

	"github.com/labstack/echo/v4"
)

func setRootRoutes(e *echo.Echo) {
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

func setApiLinksRoutes(links *echo.Group) {
	links.GET("/:linkUid/", handlers.GetLinkByUid)
	links.POST("/", handlers.CreateLink)
	links.DELETE("/:linkUid/", handlers.DeleteLink)
	links.PUT("/:linkUid/", handlers.UpdateLink)
}

func setApiCategoriesRoutes(categories *echo.Group) {
	categories.GET("/:categoryUid/", handlers.GetCategoryByUid)
	categories.POST("/", handlers.CreateCategory)
	categories.DELETE("/:categoryUid/", handlers.DeleteCategory)
	categories.PUT("/:categoryUid/", handlers.UpdateCategory)

	var belongsTo *echo.Group = categories.Group("/:categoryUid/links")
	setApiBelongsToRoutes(belongsTo)
}

func setApiBelongsToRoutes(belongsTo *echo.Group) {
	belongsTo.GET("/", handlers.GetLinksByCategory)
	belongsTo.POST("/", handlers.InsertLinkInCategory)
	belongsTo.DELETE("/:linkUid/", handlers.RemoveLinkFromCategory)
	belongsTo.PATCH("/:linkUid/", handlers.ToggleMainCategory)
}
