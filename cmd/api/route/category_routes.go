package route

import (
	"github.com/mutannejs/luof-go/cmd/api/handler"

	"github.com/labstack/echo/v4"
)

func setApiCategoriesRoutes(categories *echo.Group) {
	categories.GET("/", handler.GetAllRootCategories)
	categories.GET("/:categoryUid/", handler.GetCategoryByUid)
	categories.POST("/", handler.CreateCategory)
	categories.DELETE("/:categoryUid/", handler.DeleteCategory)
	categories.PUT("/:categoryUid/", handler.UpdateCategory)

	var subcategories *echo.Group = categories.Group("/:categoryUid/subcategories")
	setApiSubcategoriesRoutes(subcategories)

	var belongsTo *echo.Group = categories.Group("/:categoryUid/links")
	setApiBelongsToRoutes(belongsTo)
}

func setApiSubcategoriesRoutes(subcategories *echo.Group) {
	subcategories.GET("/", handler.GetSubcategories)
	subcategories.POST("/", handler.InsertSubcategory)
	subcategories.DELETE("/:childUid/", handler.RemoveSubcategory)
}
