package handlers

import (
    "net/http"

    "github.com/mutannejs/luof-go/cmd/api/custom"
    "github.com/mutannejs/luof-go/cmd/api/types"
    "github.com/mutannejs/luof-go/core/usecase/insert_link_in_category"
    // "github.com/mutannejs/luof-go/core/usecase/remove_link_from_category"
    // "github.com/mutannejs/luof-go/core/usecase/toggle_main_category"
    // "github.com/mutannejs/luof-go/core/usecase/update_category"

    "github.com/google/uuid"
    "github.com/labstack/echo/v4"
)

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

// func ToggleMainCategory(echoContext echo.Context) error {
//     var cc = echoContext.(*custom.Context)
//     var bt = types.SaveBelongsTo{}

//     if err := cc.ExecRequetJSONOperations(
//         &bt,
//         &types.SaveBelongsToSchema,
//     ); err != nil {
//         return err
//     }

//     cct := create_category.New(cc.Repositories.Category)
//     uid, err := cct.Execute(
//         c.Name,
//         c.Description,
//         c.UseMarkdown)

//     if err != nil {
//         return cc.LogAndReturnErr(err)
//     }

//     return cc.String(http.StatusOK, uid.String())
// }

// func RemoveLinkFromCategory(echoContext echo.Context) error {
//     var cc = echoContext.(*custom.Context)
//     var gc types.GetCategory

//     if err := cc.ExecRequetParamsOperations(
//         &gc,
//         &types.GetCategorySchema,
//     ); err != nil {
//         return err
//     }

//     uid, err := uuid.Parse(gc.CategoryUid)
//     if err != nil {
//         return cc.LogAndReturnErr(err)
//     }

//     dc := delete_category.New(cc.Repositories.Category)
//     _, err = dc.Execute(uid)

//     if err != nil {
//         return cc.LogAndReturnErr(err)
//     }

//     return cc.NoContent(http.StatusOK)
// }
