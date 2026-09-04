package handler

import (
	"net/http"

	"github.com/mutannejs/luof-go/cmd/api/custom"
	"github.com/mutannejs/luof-go/cmd/api/custom/custom_request"
	"github.com/mutannejs/luof-go/cmd/api/interfaces"
	"github.com/mutannejs/luof-go/core/usecase/get_links_by_category"
	"github.com/mutannejs/luof-go/core/usecase/insert_link_in_category"
	"github.com/mutannejs/luof-go/core/usecase/remove_link_from_category"
	"github.com/mutannejs/luof-go/core/usecase/toggle_main_category"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetLinksByCategory(echoContext echo.Context) error {
	var cc = echoContext.(*custom.Context)
	var gc interfaces.GetCategory

	if err := cc.Init().RequestParamsOperations(
		&gc,
		&interfaces.GetCategorySchema,
	); err != nil {
		return err
	}

	categoryUid, err := uuid.Parse(gc.CategoryUid)
	if err != nil {
		return cc.Log.ReturnInternalErr(err)
	}

	glbc := get_links_by_category.New(cc.Repositories.BelongsTo, cc.Repositories.Category)
	links, vErr := glbc.Execute(categoryUid)

	if !vErr.IsNil() {
		return cc.Log.ReturnErr(vErr)
	}

	return cc.JSON(http.StatusOK, links)
}

func InsertLinkInCategory(echoContext echo.Context) error {
	var cc = echoContext.(*custom.Context)
	var gc interfaces.GetCategory
	var cbt interfaces.CreateBelongsTo

	if err := cc.Init().RequestOperations(
		custom_request.RequestValues{ Params: &gc, JsonBody: &cbt },
		custom_request.RequestValidations{
			Params : interfaces.GetCategorySchema,
			JsonBody: interfaces.CreateBelongsToSchema,
		},
	); err != nil {
		return err
	}

	linkUid, err := uuid.Parse(cbt.LinkUid)
	if err != nil {
		return cc.Log.ReturnInternalErr(err)
	}

	categoryUid, err := uuid.Parse(gc.CategoryUid)
	if err != nil {
		return cc.Log.ReturnInternalErr(err)
	}

	ilic := insert_link_in_category.New(
		cc.Repositories.BelongsTo,
		cc.Repositories.Category,
		cc.Repositories.Link)
	vErr := ilic.Execute(
		linkUid,
		categoryUid,
		cbt.IsMain)

	if !vErr.IsNil() {
		return cc.Log.ReturnErr(vErr)
	}

	return cc.NoContent(http.StatusNoContent)
}

func RemoveLinkFromCategory(echoContext echo.Context) error {
	var cc = echoContext.(*custom.Context)
	var gbt interfaces.GetBelongsTo

	if err := cc.Init().RequestParamsOperations(
		&gbt,
		&interfaces.GetBelongsToSchema,
	); err != nil {
		return err
	}

	categoryUid, err := uuid.Parse(gbt.CategoryUid)
	if err != nil {
		return cc.Log.ReturnInternalErr(err)
	}

	linkUid, err := uuid.Parse(gbt.LinkUid)
	if err != nil {
		return cc.Log.ReturnInternalErr(err)
	}

	rlfc := remove_link_from_category.New(
		cc.Repositories.BelongsTo,
		cc.Repositories.Category,
		cc.Repositories.Link)
	vErr := rlfc.Execute(linkUid, categoryUid)

	if !vErr.IsNil() {
		return cc.Log.ReturnErr(vErr)
	}

	return cc.NoContent(http.StatusNoContent)
}

func ToggleMainCategory(echoContext echo.Context) error {
	var cc = echoContext.(*custom.Context)
	var gbt interfaces.GetBelongsTo
	var ubt interfaces.UpdateBelongsTo

	if err := cc.Init().RequestOperations(
		custom_request.RequestValues{
			Params: &gbt,
			JsonBody: &ubt,
		},
		custom_request.RequestValidations{
			Params : interfaces.GetBelongsToSchema,
			JsonBody: interfaces.UpdateBelongsToSchema,
		},
	); err != nil {
		return err
	}

	categoryUid, err := uuid.Parse(gbt.CategoryUid)
	if err != nil {
		return cc.Log.ReturnInternalErr(err)
	}

	linkUid, err := uuid.Parse(gbt.LinkUid)
	if err != nil {
		return cc.Log.ReturnInternalErr(err)
	}

	tmc := toggle_main_category.New(
		cc.Repositories.BelongsTo,
		cc.Repositories.Category,
		cc.Repositories.Link)
	vErr := tmc.Execute(
		linkUid,
		categoryUid,
		ubt.IsMain)

	if !vErr.IsNil() {
		return cc.Log.ReturnErr(vErr)
	}

	return cc.NoContent(http.StatusNoContent)
}
