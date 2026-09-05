package custom_request

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/mutannejs/luof-go/cmd/api/custom/custom_log"
	"github.com/mutannejs/luof-go/cmd/api/interfaces"
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lerror"
	"github.com/stretchr/testify/assert"
)

const (
	urlBase = "http://localhost:8123/api/categories/:categoryUid"
)

var (
	categoryUid = domain.MockUidCategory.String()
	jsonSuccess, _ = json.Marshal(map[string]string{
		"name": domain.AlternativeMockCategory.Name,
		"description": domain.AlternativeMockCategory.Description.Content,
		"useMarkdown": "false",
	})
	jsonError, _ = json.Marshal(map[string]string{
		"description": domain.AlternativeMockCategory.Description.Content,
		"useMarkdown": "teste",
	})
	paramsSuccess = map[string]string{
		"categoryUid": domain.MockUidCategory.String()}
	paramsError = map[string]string{
		"categoryUid": "teste"}
)

func TestGetPathParams(t *testing.T) {
	assert := assert.New(t)

	var cr *CRequest = New(
		&custom_log.CLog{},
		"GET",
		urlBase,
		&BodySuccess{},
		getGetParam(categoryUid),
		nil)

	params := cr.getPathParams()
	response, _ := json.Marshal(params)
	expected, _ := json.Marshal(paramsSuccess)

	assert.Equal(
		expected,
		response,
		"getPathPArams deveria retornar um map válido se passado parâmetros válidos ao CRequest")
}

func TestSetValidateErr_Success(t *testing.T) {
	assert := assert.New(t)

	var cr *CRequest = New(
		&custom_log.CLog{},
		"GET",
		urlBase,
		&BodySuccess{},
		getGetParam(categoryUid),
		nil)
		
	var validations = interfaces.GetCategorySchema
	var gc = interfaces.GetCategory{}
	var vErr = lerror.ValueError{}

	var issues = validations.Parse(paramsSuccess, &gc)
	cr.setValidateErr(issues, &vErr, VALIDATE_PARAMS_ERR)

	assert.True(
		vErr.IsNil(),
		"setValidateErr não deveria adicionar nenhum erro a vErr se passado parâmetros válidos ao CRequest")
}

func TestSetValidateErr_Errors(t *testing.T) {
	assert := assert.New(t)

	var cr *CRequest = New(
		&custom_log.CLog{},
		"GET",
		urlBase,
		&BodySuccess{},
		getGetParam(categoryUid),
		nil)

	var validations = interfaces.GetCategorySchema
	var gc = interfaces.GetCategory{}
	var vErr = lerror.ValueError{}

	var issues = validations.Parse(paramsError, &gc)
	cr.setValidateErr(issues, &vErr, VALIDATE_PARAMS_ERR)

	var expectedErrors = []error{errors.New("categoryUid: must be a valid UUID")}
	var expectedMessage = VALIDATE_PARAMS_ERR

	assert.ElementsMatch(
		expectedErrors,
		vErr.GetErrors()[0].GetErrors(),
		"Se passado um uuid inválido ao setValidateErr, deveria ser retornado o seguinte errors \"categoryUid: must be a valid UUID\"")

	assert.Equal(
		expectedMessage,
		vErr.GetErrors()[0].GetMessage(),
		"Se passado um uuid inválido ao setValidateErr, deveria ser retornado a seguinte mensagem: " + VALIDATE_PARAMS_ERR)
}

func TestGetEncodedParams_Success(t *testing.T) {
	assert := assert.New(t)

	var cr *CRequest = New(
		&custom_log.CLog{},
		"GET",
		urlBase,
		&BodySuccess{},
		getGetParam(categoryUid),
		nil)

	var validations = interfaces.GetCategorySchema
	var gc = interfaces.GetCategory{}
	var vErr = lerror.ValueError{}

	cr.getEncodedParams(
		RequestValues{ Params: &gc },
		RequestValidations{ Params: validations },
		&vErr)

	assert.True(
		vErr.IsNil(),
		"getEncodedParams não deveria adicionar nenhum erro a vErr se passado parâmetros válidos ao CRequest")
}

func TestGetEncodedParams_Error(t *testing.T) {
	assert := assert.New(t)

	var cr *CRequest = New(
		&custom_log.CLog{},
		"GET",
		urlBase,
		&BodySuccess{},
		getGetParam(paramsError["categoryUid"]),
		nil)

	var validations = interfaces.GetCategorySchema
	var gc = interfaces.GetCategory{}
	var vErr = lerror.ValueError{}

	paramsByteSlice := cr.getEncodedParams(
		RequestValues{ Params: &gc },
		RequestValidations{ Params: validations },
		&vErr)

	var expectedParams, _ = json.Marshal(paramsError)
	var expectedErrors = []error{errors.New("categoryUid: must be a valid UUID")}
	var expectedMessage = VALIDATE_PARAMS_ERR

	assert.Equal(
		expectedParams,
		paramsByteSlice,
		"getEncodedParams deveria retornar um map contendo os mesmos parâmetros passados na request")

	assert.ElementsMatch(
		expectedErrors,
		vErr.GetErrors()[0].GetErrors(),
		"Se passado um uuid inválido ao CRequest, getEncodedParams deveria retornar o seguinte errors \"categoryUid: must be a valid UUID\"")

	assert.Equal(
		expectedMessage,
		vErr.GetErrors()[0].GetMessage(),
		"Se passado um uuid inválido no CRequest, getEncodedParams deveria retornar a seguinte mensagem: " + VALIDATE_PARAMS_ERR)
}

func TestGetEncodedJsonBody_Success(t *testing.T) {
	assert := assert.New(t)

	var cr *CRequest = New(
		&custom_log.CLog{},
		"GET",
		urlBase,
		&BodySuccess{},
		getGetParam(categoryUid),
		nil)

	var validations = interfaces.SaveCategorySchema
	var sc = interfaces.SaveCategory{}
	var vErr = lerror.ValueError{}

	cr.getEncodedJsonBody(
		RequestValues{ JsonBody: &sc },
		RequestValidations{ JsonBody: validations },
		&vErr)

	assert.True(
		vErr.IsNil(),
		"getEncodedJsonBody não deveria adicionar nenhum erro a vErr se passado parâmetros válidos ao CRequest")
}

func TestGetEncodedJsonBody_Error(t *testing.T) {
	assert := assert.New(t)

	var cr *CRequest = New(
		&custom_log.CLog{},
		"GET",
		urlBase,
		&BodyError{},
		getGetParam(categoryUid),
		nil)

	var validations = interfaces.SaveCategorySchema
	var sc = interfaces.SaveCategory{}
	var vErr = lerror.ValueError{}

	cr.getEncodedJsonBody(
		RequestValues{ JsonBody: &sc },
		RequestValidations{ JsonBody: validations },
		&vErr)

	assert.Len(
		vErr.GetErrors(),
		1,
		"getEncodedJsonBody deveria retornar somente um registro de erro caso")

	assert.Equal(
		VALIDATE_BODY_ERR,
		vErr.GetErrors()[0].GetMessage(),
		"Se o corpo da requisição conter dados inválidos, getEncodedJsonBody deveria retornar a mensagem: " + VALIDATE_BODY_ERR)

	assert.ElementsMatch(
		vErr.GetErrors()[0].GetErrors(),
		[]error{
			errors.New("name: is required"),
			errors.New("useMarkdown: value is invalid"),
		})
}

/*
 Helper
*/

func getGetParam(param string) func (string) string {
	return func (name string) string {
		if (name == "categoryUid") {
			return param
		}
		return ""
	}
}

type BodySuccess struct {}

func (b *BodySuccess) Read(p []byte) (n int, err error) {
	copy(p, jsonSuccess)
	return len(p), nil
}

func (b *BodySuccess) Close() error {
	return nil
}

type BodyError struct {}

func (b *BodyError) Read(p []byte) (n int, err error) {
	copy(p, jsonError)
	return len(p), nil
}

func (b *BodyError) Close() error {
	return nil
}
