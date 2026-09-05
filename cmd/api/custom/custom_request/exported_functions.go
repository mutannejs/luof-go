package custom_request

import (
	"net/http"

	"github.com/Oudwins/zog"
	"github.com/labstack/echo/v4"
	"github.com/mutannejs/luof-go/pkg/lerror"
)

func (cr *CRequest) RequestParamsOperations(
	paramsValue any,
	validation **zog.StructSchema,
) error {
	return cr.RequestOperations(
		RequestValues{ Params: paramsValue },
		RequestValidations{ Params: *validation },
	)
}

func (cr *CRequest) RequestJSONOperations(
	jsonValue any,
	validation **zog.StructSchema,
) error {
	return cr.RequestOperations(
		RequestValues{ JsonBody: jsonValue },
		RequestValidations{ JsonBody: *validation },
	)
}

func (cr *CRequest) RequestOperations(
	values RequestValues,
	validations RequestValidations,
) error {
	// Ocorreu erro ao criar um uuid para o log
	// Retorna imediatamente
	if cr.err != nil {
		return cr.err
	}

	var vErr lerror.ValueError

	body := cr.getEncodedJsonBody(values, validations, &vErr)
	params := cr.getEncodedParams(values, validations, &vErr)

	cr.log.LogRequest(body, params, cr.method, cr.path, vErr)

	if !vErr.IsNil() {
		return echo.NewHTTPError(http.StatusBadRequest, vErr.GetErrors())
	}

	return nil
}
