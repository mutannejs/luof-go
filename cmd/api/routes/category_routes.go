package routes

import (
	"github.com/mutannejs/luof-go/cmd/api/handlers"

	"github.com/labstack/echo/v4"
)

func setApiCategoriesRoutes(categories *echo.Group) {
	categories.GET("/", handlers.GetAllRootCategories)
	categories.GET("/:categoryUid/", handlers.GetCategoryByUid)
	categories.POST("/", handlers.CreateCategory)
	categories.DELETE("/:categoryUid/", handlers.DeleteCategory)
	categories.PUT("/:categoryUid/", handlers.UpdateCategory)

	var subcategories *echo.Group = categories.Group("/:categoryUid/subcategories")
	setApiSubcategoriesRoutes(subcategories)

	var belongsTo *echo.Group = categories.Group("/:categoryUid/links")
	setApiBelongsToRoutes(belongsTo)
}

func setApiSubcategoriesRoutes(subcategories *echo.Group) {
	subcategories.GET("/", handlers.GetSubcategories)
	subcategories.POST("/", handlers.InsertSubcategory)
	subcategories.DELETE("/:childUid/", handlers.RemoveSubcategory)
}
