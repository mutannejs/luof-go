package handlers

import (
	"cmp"
	"net/http"
	"slices"

	"github.com/labstack/echo/v4"
)

func GetApi(c echo.Context) error {
	var orderedRoutes = c.Echo().Routes()
	slices.SortFunc(orderedRoutes, cmpRoutesByPath)

	var respStr = "Available Routes:\n\n"
	for _, route := range orderedRoutes {
		respStr += route.Method + "\t" + route.Path + "\n"
	}

	return c.String(http.StatusOK, respStr)
}

func cmpRoutesByPath(a, b *echo.Route) int {
	if res := cmp.Compare(a.Path, b.Path); res == 0 {
		return cmp.Compare(a.Method, b.Method)
	} else {
		return res
	}
}
