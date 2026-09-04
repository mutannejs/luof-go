package handler

import (
	"net/http"

	"github.com/mutannejs/luof-go/cmd/api/custom"
	"github.com/mutannejs/luof-go/cmd/api/custom/custom_request"
	"github.com/mutannejs/luof-go/cmd/api/interfaces"
	"github.com/mutannejs/luof-go/core/usecase/create_category"
	"github.com/mutannejs/luof-go/core/usecase/delete_category"
	"github.com/mutannejs/luof-go/core/usecase/get_all_root_categories"
	"github.com/mutannejs/luof-go/core/usecase/get_category_by_uid"
	"github.com/mutannejs/luof-go/core/usecase/get_subcategories"
	"github.com/mutannejs/luof-go/core/usecase/insert_subcategory"
	"github.com/mutannejs/luof-go/core/usecase/remove_subcategory"
	"github.com/mutannejs/luof-go/core/usecase/update_category"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetAllRootCategories(echoContext echo.Context) error {
	var cc = echoContext.(*custom.Context)

	garc := get_all_root_categories.New(cc.Repositories.Category)
	categories, vErr := garc.Execute()

	if !vErr.IsNil() {
		return cc.Log.ReturnErr(vErr)
	}

	return cc.JSON(http.StatusOK, categories)
}

func GetCategoryByUid(echoContext echo.Context) error {
	var cc = echoContext.(*custom.Context)
	var gc interfaces.GetCategory

	if err := cc.Init().RequestParamsOperations(
		&gc,
		&interfaces.GetCategorySchema,
	); err != nil {
		return err
	}

	uid, err := uuid.Parse(gc.CategoryUid)
	if err != nil {
		return cc.Log.ReturnInternalErr(err)
	}

	gcbu := get_category_by_uid.New(cc.Repositories.Category)
	c, vErr := gcbu.Execute(uid)

	if !vErr.IsNil() {
		return cc.Log.ReturnErr(vErr)
	}

	return cc.JSON(http.StatusOK, c)
}

func GetSubcategories(echoContext echo.Context) error {
	var cc = echoContext.(*custom.Context)
	var gc interfaces.GetCategory

	if err := cc.Init().RequestParamsOperations(
		&gc,
		&interfaces.GetCategorySchema,
	); err != nil {
		return err
	}

	uid, err := uuid.Parse(gc.CategoryUid)
	if err != nil {
		return cc.Log.ReturnInternalErr(err)
	}

	gs := get_subcategories.New(cc.Repositories.Category)
	subcategories, vErr := gs.Execute(uid)

	if !vErr.IsNil() {
		return cc.Log.ReturnErr(vErr)
	}

	return cc.JSON(http.StatusOK, subcategories)
}

func CreateCategory(echoContext echo.Context) error {
	var cc = echoContext.(*custom.Context)
	var c = interfaces.SaveCategory{}

	if err := cc.Init().RequestJSONOperations(
		&c,
		&interfaces.SaveCategorySchema,
	); err != nil {
		return err
	}

	cct := create_category.New(cc.Repositories.Category)
	uid, vErr := cct.Execute(
		c.Name,
		c.Description,
		c.UseMarkdown)

	if !vErr.IsNil() {
		return cc.Log.ReturnErr(vErr)
	}

	return cc.String(http.StatusCreated, uid.String())
}

func DeleteCategory(echoContext echo.Context) error {
	var cc = echoContext.(*custom.Context)
	var gc interfaces.GetCategory

	if err := cc.Init().RequestParamsOperations(
		&gc,
		&interfaces.GetCategorySchema,
	); err != nil {
		return err
	}

	uid, err := uuid.Parse(gc.CategoryUid)
	if err != nil {
		return cc.Log.ReturnInternalErr(err)
	}

	dc := delete_category.New(cc.Repositories.BelongsTo, cc.Repositories.Category)
	vErr := dc.Execute(uid)

	if !vErr.IsNil() {
		return cc.Log.ReturnErr(vErr)
	}

	return cc.NoContent(http.StatusNoContent)
}

func InsertSubcategory(echoContext echo.Context) error {
	var cc = echoContext.(*custom.Context)
	var c = interfaces.GetCategory{}
	var s = interfaces.SaveSubcategory{}

	if err := cc.Init().RequestOperations(
		custom_request.RequestValues{ JsonBody: &s, Params: &c },
		custom_request.RequestValidations{ JsonBody: interfaces.SaveSubcategorySchema, Params: interfaces.GetCategorySchema },
	); err != nil {
		return err
	}

	fatherUid, err := uuid.Parse(c.CategoryUid)
	if err != nil {
		return cc.Log.ReturnInternalErr(err)
	}

	childUid, err := uuid.Parse(s.ChildUid)
	if err != nil {
		return cc.Log.ReturnInternalErr(err)
	}

	ic := insert_subcategory.New(cc.Repositories.Category)
	vErr := ic.Execute(
		fatherUid,
		childUid)

	if !vErr.IsNil() {
		return cc.Log.ReturnErr(vErr)
	}

	return cc.NoContent(http.StatusNoContent)
}

func RemoveSubcategory(echoContext echo.Context) error {
	var cc = echoContext.(*custom.Context)
	var rs interfaces.RemoveSubcategory

	if err := cc.Init().RequestParamsOperations(
		&rs,
		&interfaces.RemoveSubcategorySchema,
	); err != nil {
		return err
	}

	categoryUid, err := uuid.Parse(rs.CategoryUid)
	if err != nil {
		return cc.Log.ReturnInternalErr(err)
	}

	childUid, err := uuid.Parse(rs.ChildUid)
	if err != nil {
		return cc.Log.ReturnInternalErr(err)
	}

	rsuc := remove_subcategory.New(cc.Repositories.Category)
	vErr := rsuc.Execute(categoryUid, childUid)

	if !vErr.IsNil() {
		return cc.Log.ReturnErr(vErr)
	}

	return cc.NoContent(http.StatusNoContent)
}

func UpdateCategory(echoContext echo.Context) error {
	var cc = echoContext.(*custom.Context)
	var c = interfaces.SaveCategory{}
	var gc = interfaces.GetCategory{}

	if err := cc.Init().RequestOperations(
		custom_request.RequestValues{ JsonBody: &c, Params: &gc },
		custom_request.RequestValidations{ JsonBody: interfaces.SaveCategorySchema, Params: interfaces.GetCategorySchema },
	); err != nil {
		return err
	}

	uid, err := uuid.Parse(gc.CategoryUid)
	if err != nil {
		return cc.Log.ReturnInternalErr(err)
	}

	uc := update_category.New(cc.Repositories.Category)
	_, vErr := uc.Execute(
		uid,
		c.Name,
		c.Description,
		c.UseMarkdown)

	if !vErr.IsNil() {
		return cc.Log.ReturnErr(vErr)
	}

	return cc.NoContent(http.StatusNoContent)
}
