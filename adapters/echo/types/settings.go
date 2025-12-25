package types

import (
	"net/http"

	"github.com/mutannejs/luof-go/core/repository"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
    echo.Context
    Repositories repository.Repositories
}

type CustomValidator struct {
    Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
    if err := cv.Validator.Struct(i); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return nil
}
