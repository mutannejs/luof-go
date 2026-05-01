package handlers

import (
	"net/http"

	"github.com/mutannejs/luof-go/cmd/api/custom"
	"github.com/mutannejs/luof-go/cmd/api/types"
	"github.com/mutannejs/luof-go/core/usecase/get_links_by_category"
	"github.com/mutannejs/luof-go/core/usecase/insert_link_in_category"
	"github.com/mutannejs/luof-go/core/usecase/remove_link_from_category"
	"github.com/mutannejs/luof-go/core/usecase/toggle_main_category"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetLinksByCategory(echoContext echo.Context) error {
	var cc = echoContext.(*custom.Context)
	var gc types.GetCategory

	if err := cc.ExecRequetParamsOperations(
		&gc,
		&types.GetCategorySchema,
	); err != nil {
		return err
	}

	categoryUid, err := uuid.Parse(gc.CategoryUid)
	if err != nil {
		return cc.LogAndReturnErr(err)
	}

	glbc := get_links_by_category.New(cc.Repositories.BelongsTo, cc.Repositories.Category)
	links, err := glbc.Execute(categoryUid)

	if err != nil {
		return cc.LogAndReturnErr(err)
	}

	/*if len(links) == 0 {
		return cc.JSON(http.StatusOK, make([]any, 0, 0))
	}*/

	return cc.JSON(http.StatusOK, links)
}

func InsertLinkInCategory(echoContext echo.Context) error {
	var cc = echoContext.(*custom.Context)
	var gc types.GetCategory
	var cbt types.CreateBelongsTo

	if err := cc.ExecRequetOperations(
		custom.RequestValues{ Params: &gc, JsonBody: &cbt },
		custom.RequestValidations{
			Params : types.GetCategorySchema,
			JsonBody: types.CreateBelongsToSchema,
		},
	); err != nil {
		return err
	}

	linkUid, err := uuid.Parse(cbt.LinkUid)
	if err != nil {
		return cc.LogAndReturnErr(err)
	}

	categoryUid, err := uuid.Parse(gc.CategoryUid)
	if err != nil {
		return cc.LogAndReturnErr(err)
	}

	ilic := insert_link_in_category.New(cc.Repositories.BelongsTo)
	err = ilic.Execute(
		linkUid,
		categoryUid,
		cbt.IsMain)

	if err != nil {
		return cc.LogAndReturnErr(err)
	}

	return cc.NoContent(http.StatusOK)
}

func RemoveLinkFromCategory(echoContext echo.Context) error {
	var cc = echoContext.(*custom.Context)
	var gbt types.GetBelongsTo

	if err := cc.ExecRequetParamsOperations(
		&gbt,
		&types.GetBelongsToSchema,
	); err != nil {
		return err
	}

	categoryUid, err := uuid.Parse(gbt.CategoryUid)
	if err != nil {
		return cc.LogAndReturnErr(err)
	}

	linkUid, err := uuid.Parse(gbt.LinkUid)
	if err != nil {
		return cc.LogAndReturnErr(err)
	}

	rlfc := remove_link_from_category.New(cc.Repositories.BelongsTo)
	err = rlfc.Execute(linkUid, categoryUid)

	if err != nil {
		return cc.LogAndReturnErr(err)
	}

	return cc.NoContent(http.StatusOK)
}

func ToggleMainCategory(echoContext echo.Context) error {
	var cc = echoContext.(*custom.Context)
	var gbt types.GetBelongsTo
	var ubt types.UpdateBelongsTo

	if err := cc.ExecRequetOperations(
		custom.RequestValues{
			Params: &gbt,
			JsonBody: &ubt,
		},
		custom.RequestValidations{
			Params : types.GetBelongsToSchema,
			JsonBody: types.UpdateBelongsToSchema,
		},
	); err != nil {
		return err
	}

	categoryUid, err := uuid.Parse(gbt.CategoryUid)
	if err != nil {
		return cc.LogAndReturnErr(err)
	}

	linkUid, err := uuid.Parse(gbt.LinkUid)
	if err != nil {
		return cc.LogAndReturnErr(err)
	}

	tmc := toggle_main_category.New(cc.Repositories.BelongsTo)
	err = tmc.Execute(
		linkUid,
		categoryUid,
		ubt.IsMain)

	if err != nil {
		return cc.LogAndReturnErr(err)
	}

	return cc.NoContent(http.StatusOK)
}
