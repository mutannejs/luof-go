package handler

import (
	"net/http"

	"github.com/mutannejs/luof-go/cmd/api/custom"
	"github.com/mutannejs/luof-go/cmd/api/interfaces"
	"github.com/mutannejs/luof-go/core/usecase/create_link"
	"github.com/mutannejs/luof-go/core/usecase/delete_link"
	"github.com/mutannejs/luof-go/core/usecase/get_link_by_uid"
	"github.com/mutannejs/luof-go/core/usecase/update_link"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetLinkByUid(c echo.Context) error {
	var cc = c.(*custom.Context)
	var gl interfaces.GetLink

	if err := cc.ExecRequetParamsOperations(
		&gl,
		&interfaces.GetLinkSchema,
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
	var l = interfaces.SaveLink{}

	if err := cc.ExecRequetJSONOperations(
		&l,
		&interfaces.SaveLinkSchema,
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

	return cc.String(http.StatusCreated, uid.String())
}

func DeleteLink(c echo.Context) error {
	var cc = c.(*custom.Context)
	var gl interfaces.GetLink

	if err := cc.ExecRequetParamsOperations(
		&gl,
		&interfaces.GetLinkSchema,
	); err != nil {
		return err
	}

	uid, err := uuid.Parse(gl.LinkUid)
	if err != nil {
		return cc.LogAndReturnErr(err)
	}

	dl := delete_link.New(cc.Repositories.Link)
	_, err = dl.Execute(uid)

	if err != nil {
		return cc.LogAndReturnErr(err)
	}

	return cc.NoContent(http.StatusNoContent)
}

func UpdateLink(c echo.Context) error {
	var cc = c.(*custom.Context)
	var l = interfaces.SaveLink{}
	var gl = interfaces.GetLink{}

	if err := cc.ExecRequetOperations(
		custom.RequestValues{ JsonBody: &l, Params: &gl },
		custom.RequestValidations{ JsonBody: interfaces.SaveLinkSchema, Params: interfaces.GetLinkSchema },
	); err != nil {
		return err
	}

	uid, err := uuid.Parse(gl.LinkUid)
	if err != nil {
		return cc.LogAndReturnErr(err)
	}

	ul := update_link.New(cc.Repositories.Link)
	_, err = ul.Execute(
		uid,
		l.Url,
		l.Name,
		l.Description,
		l.UseMarkdown)

	if err != nil {
		return cc.LogAndReturnErr(err)
	}

	return cc.NoContent(http.StatusNoContent)
}
