package custom_request

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mutannejs/luof-go/cmd/api/custom/custom_log"
	"github.com/mutannejs/luof-go/cmd/api/interfaces"
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/stretchr/testify/assert"
)

var (
	paramsErr = []lerror.MsgErrors{
		{
			Message: VALIDATE_PARAMS_ERR,
			Errors: []string{"categoryUid: must be a valid UUID"},
		},
	}
	jsonErr = []lerror.MsgErrors{
		{
			Message: VALIDATE_BODY_ERR,
			Errors: []string{"name: is required", "useMarkdown: value is invalid"},
		},
	}
)

func TestRequestParamsOperations_Success(t *testing.T) {
	assert := assert.New(t)

	var cr *CRequest = New(
		&custom_log.CLog{},
		"GET",
		urlBase + ":categoryUid/",
		urlBase + categoryUid + "/",
		&BodySuccess{},
		getGetParam(categoryUid),
		sendJson,
		nil)

	var gc interfaces.GetCategory

	err := cr.RequestParamsOperations(
		&gc,
		&interfaces.GetCategorySchema,
	)

	assert.NoError(
		err,
		"Se passado parâmetros válidos na request, RequestParamsOperations não deveria retornar erro")

	assert.Equal(
		categoryUid,
		gc.CategoryUid,
		"Se passado parâmetros válidos na request, values passado em RequestParamsOperations deveria ser corretamente preenchida")
}

func TestRequestParamsOperations_Error(t *testing.T) {
	assert := assert.New(t)

	var cr *CRequest = New(
		&custom_log.CLog{},
		"GET",
		urlBase + ":categoryUid/",
		urlBase + paramsError["categoryUid"],
		&BodySuccess{},
		getGetParam(paramsError["categoryUid"]),
		sendJson,
		nil)

	var gc interfaces.GetCategory

	err := cr.RequestParamsOperations(
		&gc,
		&interfaces.GetCategorySchema,
	)

	firstMsgErrs := getFirstMsgErrs(err)

	assert.Error(
		err,
		"Se passado parâmetros inválidos na request, RequestParamsOperations deveria retornar erro")

	assert.Equal(
		paramsErr[0].GetMessage(),
		firstMsgErrs.GetMessage(),
		"Se passado parâmetros inválidos na request, RequestParamsOperations deveria retornar a mensagem de erro: " + VALIDATE_PARAMS_ERR)

	assert.ElementsMatch(
		paramsErr[0].GetErrors(),
		firstMsgErrs.GetErrors(),
		"Se passado parâmetros inválidos na request, RequestParamsOperations deveria retornar quais parâmetros são inválidos")
}

func TestRequestJSONOperations_Success(t *testing.T) {
	assert := assert.New(t)

	var cr *CRequest = New(
		&custom_log.CLog{},
		"POST",
		urlBase + ":categoryUid/",
		urlBase + categoryUid + "/",
		&BodySuccess{},
		getGetParam(categoryUid),
		sendJson,
		nil)

	var sc interfaces.SaveCategory

	err := cr.RequestJSONOperations(
		&sc,
		&interfaces.SaveCategorySchema,
	)

	assert.NoError(
		err,
		"Se passado um corpo json válido na request, RequestJSONOperations não deveria retornar erro")

	assert.Equal(
		domain.AlternativeMockCategory.Name,
		sc.Name,
		"Se passado um corpo json válido na request, values passado em RequestJSONOperations deveria ser corretamente preenchida")

	assert.Equal(
		domain.AlternativeMockCategory.Description.Content,
		sc.Description,
		"Se passado um corpo json válido na request, values passado em RequestJSONOperations deveria ser corretamente preenchida")

	assert.Equal(
		domain.AlternativeMockCategory.Description.UseMarkdown,
		sc.UseMarkdown,
		"Se passado um corpo json válido na request, values passado em RequestJSONOperations deveria ser corretamente preenchida")
}

func TestRequestJSONOperations_Error(t *testing.T) {
	assert := assert.New(t)

	var cr *CRequest = New(
		&custom_log.CLog{},
		"POST",
		urlBase + ":categoryUid/",
		urlBase + categoryUid + "/",
		&BodyError{},
		getGetParam(categoryUid),
		sendJson,
		nil)

	var sc interfaces.SaveCategory

	err := cr.RequestJSONOperations(
		&sc,
		&interfaces.SaveCategorySchema,
	)

	firstMsgErrs := getFirstMsgErrs(err)

	assert.Error(
		err,
		"Se passado um corpo json inválido na request, RequestJSONOperations deveria retornar erro")

	assert.Equal(
		jsonErr[0].GetMessage(),
		firstMsgErrs.GetMessage(),
		"Se passado um corpo json inválido na request, RequestJSONOperations deveria retornar a mensagem de erro: " + VALIDATE_BODY_ERR)

	assert.ElementsMatch(
		jsonErr[0].GetErrors(),
		firstMsgErrs.GetErrors(),
		"Se passado um corpo json inválido na request, RequestJSONOperations deveria retornar quais campos são inválidos")
}

func TestRequestOperations_Success(t *testing.T) {
	assert := assert.New(t)

	var cr *CRequest = New(
		&custom_log.CLog{},
		"POST",
		urlBase + ":categoryUid/",
		urlBase + categoryUid + "/",
		&BodySuccess{},
		getGetParam(categoryUid),
		sendJson,
		nil)

	var sc interfaces.SaveCategory
	var gc interfaces.GetCategory

	err := cr.RequestOperations(
		RequestValues{ JsonBody: &sc, Params: &gc },
		RequestValidations{
			JsonBody: interfaces.SaveCategorySchema,
			Params: interfaces.GetCategorySchema,
		},
	)

	assert.NoError(
		err,
		"Se passado campos válidos na request, RequestOperations não deveria retornar erro")

	assert.Equal(
		categoryUid,
		gc.CategoryUid,
		"Se passado campos válidos na request, values passado em RequestOperations deveria ser corretamente preenchida")

	assert.Equal(
		domain.AlternativeMockCategory.Name,
		sc.Name,
		"Se passado campos válidos na request, values passado em RequestOperations deveria ser corretamente preenchida")

	assert.Equal(
		domain.AlternativeMockCategory.Description.Content,
		sc.Description,
		"Se passado campos válidos na request, values passado em RequestOperations deveria ser corretamente preenchida")

	assert.Equal(
		domain.AlternativeMockCategory.Description.UseMarkdown,
		sc.UseMarkdown,
		"Se passado campos válidos na request, values passado em RequestOperations deveria ser corretamente preenchida")
}

func TestRequestOperations_ParamsError(t *testing.T) {
	assert := assert.New(t)

	var cr *CRequest = New(
		&custom_log.CLog{},
		"POST",
		urlBase + ":categoryUid/",
		urlBase + paramsError["categoryUid"],
		&BodySuccess{},
		getGetParam(paramsError["categoryUid"]),
		sendJson,
		nil)

	var sc interfaces.SaveCategory
	var gc interfaces.GetCategory

	err := cr.RequestOperations(
		RequestValues{ JsonBody: &sc, Params: &gc },
		RequestValidations{
			JsonBody: interfaces.SaveCategorySchema,
			Params: interfaces.GetCategorySchema,
		},
	)

	firstMsgErrs := getFirstMsgErrs(err)

	assert.Error(
		err,
		"Se passado parâmetros inválidos na request, RequestJSONOperations deveria retornar erro")

	assert.Equal(
		paramsErr[0].GetMessage(),
		firstMsgErrs.GetMessage(),
		"Se passado parâmetros inválidos na request, RequestJSONOperations deveria retornar a mensagem de erro: " + VALIDATE_PARAMS_ERR)

	assert.ElementsMatch(
		paramsErr[0].GetErrors(),
		firstMsgErrs.GetErrors(),
		"Se passado parâmetros inválidos na request, RequestJSONOperations deveria retornar quais parâmetros são inválidos")
}

func TestRequestOperations_JsonError(t *testing.T) {
	assert := assert.New(t)

	var cr *CRequest = New(
		&custom_log.CLog{},
		"POST",
		urlBase + ":categoryUid/",
		urlBase + categoryUid + "/",
		&BodyError{},
		getGetParam(categoryUid),
		sendJson,
		nil)

	var sc interfaces.SaveCategory
	var gc interfaces.GetCategory

	err := cr.RequestOperations(
		RequestValues{ JsonBody: &sc, Params: &gc },
		RequestValidations{
			JsonBody: interfaces.SaveCategorySchema,
			Params: interfaces.GetCategorySchema,
		},
	)

	firstMsgErrs := getFirstMsgErrs(err)

	assert.Error(
		err,
		"Se passado um corpo json inválido na request, RequestOperations deveria retornar erro")

	assert.Equal(
		jsonErr[0].GetMessage(),
		firstMsgErrs.GetMessage(),
		"Se passado um corpo json inválido na request, RequestOperations deveria retornar a mensagem de erro: " + VALIDATE_BODY_ERR)

	assert.ElementsMatch(
		jsonErr[0].GetErrors(),
		firstMsgErrs.GetErrors(),
		"Se passado um corpo json inválido na request, RequestOperations deveria retornar quais campos são inválidos")
}

/*
Helper
*/

func sendJson(_ int, i any) error {
	json, _ := json.Marshal(i)
	return errors.New(string(json))
}

func getFirstMsgErrs(err error) lerror.MsgErrors {
	var response *echo.HTTPError = err.(*echo.HTTPError)
	return response.Message.([]lerror.MsgErrors)[0]
}
