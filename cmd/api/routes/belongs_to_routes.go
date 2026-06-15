package routes

import (
	"github.com/mutannejs/luof-go/cmd/api/handlers"

	"github.com/labstack/echo/v4"
)

func setApiBelongsToRoutes(belongsTo *echo.Group) {
	belongsTo.GET("/", handlers.GetLinksByCategory)
	belongsTo.POST("/", handlers.InsertLinkInCategory)
	belongsTo.DELETE("/:linkUid/", handlers.RemoveLinkFromCategory)
	belongsTo.PATCH("/:linkUid/", handlers.ToggleMainCategory)
}
