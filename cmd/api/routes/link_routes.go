package routes

import (
	"github.com/mutannejs/luof-go/cmd/api/handlers"

	"github.com/labstack/echo/v4"
)

func setApiLinksRoutes(links *echo.Group) {
	links.GET("/:linkUid/", handlers.GetLinkByUid)
	links.POST("/", handlers.CreateLink)
	links.DELETE("/:linkUid/", handlers.DeleteLink)
	links.PUT("/:linkUid/", handlers.UpdateLink)
}
