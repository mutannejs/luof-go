package echo

import (
	"github.com/mutannejs/luof-go/adapters/echo/handlers"

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
    links.GET("/:linkId", handlers.GetLink)
    links.POST("/", handlers.CreateLink)
    links.DELETE("/:linkId", handlers.DeleteLink)
    links.PUT("/:linkId", handlers.UpdateLink)
    links.PATCH("/:linkId", handlers.PartialUpdateLink)
}

func setApiCategoriesRoutes(categories *echo.Group) {
    categories.GET("/:categoryId", handlers.GetCategory)
    categories.POST("/", handlers.CreateCategory)
    categories.DELETE("/:categoryId", handlers.DeleteCategory)
    categories.PUT("/:categoryId", handlers.UpdateCategory)
    categories.PATCH("/:categoryId", handlers.PartialUpdateCategory)

    var belongsTo *echo.Group = categories.Group("/:categoryId/links")
    setApiBelongsToRoutes(belongsTo)
}

func setApiBelongsToRoutes(belongsTo *echo.Group) {
    belongsTo.GET("/", handlers.ListBelongsTo)
    belongsTo.POST("/", handlers.CreateBelongsTo)
    belongsTo.DELETE("/:linkId", handlers.DeleteBelongsTo)
    belongsTo.PATCH("/:linkId", handlers.PartialUpdateBelongsTo)
}
