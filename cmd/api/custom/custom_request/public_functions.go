package custom_request

import (
	"net/http"

	"github.com/Oudwins/zog"
	"github.com/labstack/echo/v4"
	"github.com/mutannejs/luof-go/pkg/lerror"
)

// Valida os parâmetros enviados na URL com base no schema
// passado em "validation", e atribui os parâmetros à
// variável passada em "values". Retorna erro descritivo
// se a validação ou conversão dos dados falhou
func (cr *CRequest) RequestParamsOperations(
	paramsValue any,
	validation **zog.StructSchema,
) error {
	return cr.RequestOperations(
		RequestValues{ Params: paramsValue },
		RequestValidations{ Params: *validation },
	)
}

// Valida o corpo json enviado na request com base no schema
// passado em "validation", e atribui o corpo da request à
// estrutura passada em "values". Retorna erro descritivo se
// a validação ou conversão dos dados falhou
func (cr *CRequest) RequestJSONOperations(
	jsonValue any,
	validation **zog.StructSchema,
) error {
	return cr.RequestOperations(
		RequestValues{ JsonBody: jsonValue },
		RequestValidations{ JsonBody: *validation },
	)
}

// Valida os campos enviados na request com base nos schemas
// passados em "validations", e atribui os campos da request
// às estrutura passadas em "values". Retorna erro
// descritivo se a validação ou conversão dos dados falhou
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
