package route

import (
	"github.com/mutannejs/luof-go/cmd/api/handler"

	"github.com/labstack/echo/v4"
)

func setApiLinksRoutes(links *echo.Group) {
	links.GET("/:linkUid/", handler.GetLinkByUid)
	links.POST("/", handler.CreateLink)
	links.DELETE("/:linkUid/", handler.DeleteLink)
	links.PUT("/:linkUid/", handler.UpdateLink)
}
