package handlers

import (
	"net/http"

    "github.com/mutannejs/luof-go/cmd/api/custom"
	"github.com/mutannejs/luof-go/cmd/api/types"
	"github.com/mutannejs/luof-go/core/usecase/create_link"
	"github.com/mutannejs/luof-go/core/usecase/get_link_by_uid"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetLink(c echo.Context) error {
    var cc = c.(*custom.Context)
    var gl types.GetLink

    if err := cc.ExecRequetParamsOperations(
        &gl,
        &types.GetLinkSchema,
    ); err != nil {
        return err
    }

    uid, err := uuid.Parse(gl.LinkUid)
    if err != nil {
        return cc.LogAndReturnErr(err)
    }

    glbu := get_link_by_uid.New(cc.Repositories.Link)
    l, err := glbu.Execute(uid)
    if err != nil {
        return cc.LogAndReturnErr(err)
    }

    return cc.JSON(http.StatusOK, l)
}

func CreateLink(c echo.Context) error {
    var cc = c.(*custom.Context)
    var l = types.SaveLink{}

    if err := cc.ExecRequetJSONOperations(
        &l,
        &types.SaveLinkSchema,
    ); err != nil {
        return err
    }

    cl := create_link.New(cc.Repositories.Link)
    uid, err := cl.Execute(
        l.Url,
        l.Name,
        l.Description,
        l.UseMarkdown)

    if err != nil {
        return cc.LogAndReturnErr(err)
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
