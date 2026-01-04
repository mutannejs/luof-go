package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/mutannejs/luof-go/cmd/api/types"
	"github.com/mutannejs/luof-go/core/usecase/create_link"
	"github.com/mutannejs/luof-go/core/usecase/get_link_by_uid"
	"github.com/rs/zerolog/log"

	"github.com/labstack/echo/v4"
)

func GetLink(c echo.Context) (err error) {
    var cc = c.(*types.CustomContext)
    var gl types.GetLink

    if err = cc.InitReqStruct(&gl); err != nil {
        return err
    }

    uid, err := uuid.Parse(gl.Uid)
    if err != nil {
        log.Error().Err(err).Send()
        return err
    }

    glbu := get_link_by_uid.New(cc.Repositories.Link)
    l, err := glbu.Execute(uid)

    if err != nil {
        log.Error().Err(err).Send()
        return err
    }

    return cc.JSON(http.StatusOK, l)
}

func CreateLink(c echo.Context) (err error) {
    var cc = c.(*types.CustomContext)
    var l types.SaveLink

    if err = cc.InitReqStruct(&l); err != nil {
        return err
    }

    cl := create_link.New(cc.Repositories.Link)
    uid, err := cl.Execute(
        l.Url,
        l.Name,
        l.Description,
        l.UseMarkdown)

    if err != nil {
        return err
    }

    return cc.String(http.StatusOK, uid.String())
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
