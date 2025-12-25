package handlers

import (
	"net/http"

	"github.com/mutannejs/luof-go/adapters/echo/types"
	"github.com/mutannejs/luof-go/core/usecase/create_link"

	"github.com/labstack/echo/v4"
)

func GetLink(c echo.Context) error {
    return c.String(http.StatusOK, "OK")
}

func CreateLink(c echo.Context) (err error) {
    var l types.SaveLink

    if err = c.Bind(l); err != nil {
      return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    if err = c.Validate(l); err != nil {
      return err
    }

    cc := c.(*types.CustomContext)

    cl := create_link.New(cc.Repositories.Link)
    uid, err := cl.Execute(
        l.Url,
        l.Name,
        l.Description,
        l.UseMarkdown)

    return c.String(http.StatusOK, uid.String())
}

func DeleteLink(c echo.Context) error {
    return c.String(http.StatusOK, "OK")
}

func UpdateLink(c echo.Context) error {
    return c.String(http.StatusOK, "OK")
}

func PartialUpdateLink(c echo.Context) error {
    return c.String(http.StatusOK, "OK")
}
