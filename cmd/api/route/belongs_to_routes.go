package route

import (
	"github.com/mutannejs/luof-go/cmd/api/handler"

	"github.com/labstack/echo/v4"
)

func setApiBelongsToRoutes(belongsTo *echo.Group) {
	belongsTo.GET("/", handler.GetLinksByCategory)
	belongsTo.POST("/", handler.InsertLinkInCategory)
	belongsTo.DELETE("/:linkUid/", handler.RemoveLinkFromCategory)
	belongsTo.PATCH("/:linkUid/", handler.ToggleMainCategory)
}
