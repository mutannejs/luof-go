package types

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/mutannejs/luof-go/core/repository"
	"github.com/rs/zerolog/log"

	"github.com/labstack/echo/v4"
)

type CustomContext struct {
    echo.Context
    Repositories repository.Repositories
    NilValues []string
}

func (cc *CustomContext) BindAndValidate(v any) (err error) {
    if err = cc.ProcessNoRequiredFields(v); err != nil {
      return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    if err = cc.Bind(v); err != nil {
      return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    if err = cc.Validate(v); err != nil {
      return err
    }
    return
}

func (cc *CustomContext) ProcessNoRequiredFields(value any) error {
    tValue := reflect.TypeOf(value).Elem()
    vValue := reflect.ValueOf(value).Elem()
    formValues, err := cc.FormParams()

    logRequest := log.Info().
        Str("path", cc.Request().URL.Path).
        Str("method", cc.Request().Method)

    if err != nil {
        return err
    }

    for i := 0; i < tValue.NumField(); i++ {
        vField := vValue.Field(i)
        tField := tValue.Field(i)

        formTag, _ := tField.Tag.Lookup("form")
        validateTag, exists := tField.Tag.Lookup("validate")

        if !formValues.Has(formTag) {
            cc.NilValues = append(cc.NilValues, formTag)
            logRequest = logRequest.Str(formTag, "<nil>")
        } else {
            logRequest = logRequest.Str(formTag, formValues.Get(formTag))
        }

        var isAbsentAndNotRequired bool =
            exists &&
            !strings.Contains(validateTag, "required") &&
            !formValues.Has(formTag)

        if !isAbsentAndNotRequired {
           continue
        }

        props := strings.Split(validateTag, ",")

        switch props[0] {
            case "url": vField.SetString("http://localhost:8123")
        }
    }

    logRequest.Send()

    return nil
}
