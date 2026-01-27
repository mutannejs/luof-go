package handlers

import (
	"net/http"

	"github.com/mutannejs/luof-go/cmd/api/custom"
	"github.com/mutannejs/luof-go/cmd/api/types"
	"github.com/mutannejs/luof-go/core/usecase/create_category"
	"github.com/mutannejs/luof-go/core/usecase/delete_category"
	"github.com/mutannejs/luof-go/core/usecase/get_category_by_uid"
	"github.com/mutannejs/luof-go/core/usecase/update_category"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetCategoryByUid(echoContext echo.Context) error {
	var cc = echoContext.(*custom.Context)
	var gc types.GetCategory

	if err := cc.ExecRequetParamsOperations(
		&gc,
		&types.GetCategorySchema,
	); err != nil {
		return err
	}

	uid, err := uuid.Parse(gc.CategoryUid)
	if err != nil {
		return cc.LogAndReturnErr(err)
	}

	gcbu := get_category_by_uid.New(cc.Repositories.Category)
	c, err := gcbu.Execute(uid)

	if err != nil {
		return cc.LogAndReturnErr(err)
	}

	return cc.JSON(http.StatusOK, c)
}

func CreateCategory(echoContext echo.Context) error {
	var cc = echoContext.(*custom.Context)
	var c = types.SaveCategory{}

	if err := cc.ExecRequetJSONOperations(
		&c,
		&types.SaveCategorySchema,
	); err != nil {
		return err
	}

	cct := create_category.New(cc.Repositories.Category)
	uid, err := cct.Execute(
		c.Name,
		c.Description,
		c.UseMarkdown)

	if err != nil {
		return cc.LogAndReturnErr(err)
	}

	return cc.String(http.StatusOK, uid.String())
}

func DeleteCategory(echoContext echo.Context) error {
	var cc = echoContext.(*custom.Context)
	var gc types.GetCategory

	if err := cc.ExecRequetParamsOperations(
		&gc,
		&types.GetCategorySchema,
	); err != nil {
		return err
	}

	uid, err := uuid.Parse(gc.CategoryUid)
	if err != nil {
		return cc.LogAndReturnErr(err)
	}

	dc := delete_category.New(cc.Repositories.Category)
	_, err = dc.Execute(uid)

	if err != nil {
		return cc.LogAndReturnErr(err)
	}

	return cc.NoContent(http.StatusOK)
}

func UpdateCategory(echoContext echo.Context) error {
	var cc = echoContext.(*custom.Context)
	var c = types.SaveCategory{}
	var gc = types.GetCategory{}

	if err := cc.ExecRequetOperations(
		custom.RequestValues{ JsonBody: &c, Params: &gc },
		custom.RequestValidations{ JsonBody: types.SaveCategorySchema, Params: types.GetCategorySchema },
	); err != nil {
		return err
	}

	uid, err := uuid.Parse(gc.CategoryUid)
	if err != nil {
		return cc.LogAndReturnErr(err)
	}

	uc := update_category.New(cc.Repositories.Category)
	_, err = uc.Execute(
		uid,
		c.Name,
		c.Description,
		c.UseMarkdown)

	if err != nil {
		return cc.LogAndReturnErr(err)
	}

	return cc.NoContent(http.StatusOK)
}
