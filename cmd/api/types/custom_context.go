package types

import (
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/mutannejs/luof-go/core/repository"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type CustomContext struct {
    echo.Context
    Repositories repository.Repositories
}

const (
    PARAM = "param"
    FORM = "form"
)

func (cc *CustomContext) InitReqStruct(v any) (err error) {
    if err = cc.LogAndValidate(v); err != nil {
      return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    if err = cc.Bind(v); err != nil {
      return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return
}

func (cc *CustomContext) LogAndValidate(obj any) error {
    logReq := log.Info().
        Str("path", cc.Request().URL.Path).
        Str("method", cc.Request().Method)

    formValues, err := cc.FormParams()
    if err != nil {
        return err
    }

    validate := validator.New(validator.WithRequiredStructEnabled())

    var objType reflect.Type = reflect.TypeOf(obj).Elem()

    var structField reflect.StructField
    var value string
    var valueExists bool
    var bindTagKey string
    var bindTagValue string
    var validateTagValue string
    var validateTagExists bool

    for i := 0; i < objType.NumField(); i++ {
        structField = objType.Field(i)

        if exists := cc.setBindTag(structField, &bindTagKey, &bindTagValue); !exists {
            return errors.New("errror binding request")
        }

        valueExists = cc.setValue(formValues, bindTagKey, bindTagValue, &value)

        cc.logAttribute(logReq, bindTagValue, value, valueExists)

        validateTagValue, validateTagExists = structField.Tag.Lookup("validate")

        var isAbsentAndNotRequired bool =
            validateTagExists &&
            !strings.Contains(validateTagValue, "required") &&
            !valueExists

        if !isAbsentAndNotRequired {
           err = errors.Join(err, validate.VarWithKey(bindTagValue, value, validateTagValue))
        }
    }

    logReq.Send()
    return err
}

func (cc *CustomContext) setBindTag(
    structField reflect.StructField,
    bindTagKey *string,
    bindTagValue *string,
) bool {
    for _, tag := range []string{FORM, PARAM} {
        if value, exists := structField.Tag.Lookup(tag); !exists {
            continue
        } else {
            *bindTagKey = tag
            *bindTagValue = value
            return true
        }
    }
    return false
}

func (cc *CustomContext) setValue(
    formValues url.Values,
    bindTagKey string,
    bindTagValue string,
    value *string,
) bool {
    switch bindTagKey {
        case FORM:
            if formValues.Has(bindTagValue) {
                *value = formValues.Get(bindTagValue)
                return true
            }
        case PARAM:
            *value = cc.Param(bindTagValue)
            return true
    }
    return false
}

func (cc *CustomContext) logAttribute(
    logReq *zerolog.Event,
    bindTagValue string,
    value string,
    valueExists bool,
) {
    if valueExists {
        logReq = logReq.Str(bindTagValue, value)
    } else {
        logReq = logReq.Str(bindTagValue, "<nil>")
    }
}
