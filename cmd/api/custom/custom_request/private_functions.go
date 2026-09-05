package custom_request

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/Oudwins/zog"
)

const (
	JSON_BODY_ERR = "the request body could not be interpreted in JSON format"
	PARAMS_ERR = "the path params could not be converted in JSON format"
	VALIDATE_PARAMS_ERR = "errors occurred during the validation of the request parameters"
	VALIDATE_BODY_ERR = "errors occurred during the validation of the request body"
)

func (cr *CRequest) getEncodedJsonBody(
	values RequestValues, 
	validations RequestValidations,
	vErr *lerror.ValueError,
) (
	encodedBody []byte,
) {
	if values.JsonBody == nil {
		return
	}

	var bodyMap = make(map[string]any)

	bodyErr := json.NewDecoder(cr.body).Decode(&bodyMap)
	encodedBody, encodedBodyErr := json.Marshal(bodyMap)

	if bodyErr != nil || encodedBodyErr != nil {
		vErr.AppendErr(JSON_BODY_ERR)
	} else if validations.JsonBody != nil {
		parseErrs := validations.JsonBody.Parse(bodyMap, values.JsonBody)
		cr.setValidateErr(parseErrs, vErr, VALIDATE_BODY_ERR)
	}

	return
}

func (cr *CRequest) getEncodedParams(
	values RequestValues, 
	validations RequestValidations,
	vErr *lerror.ValueError,
) (
	encodedParams []byte,
) {
	if values.Params == nil {
		return
	}

	paramsMap := cr.getPathParams()
	encodedParams, encodedParamsErr := json.Marshal(paramsMap)

	if encodedParamsErr != nil {
		vErr.AppendErr(PARAMS_ERR)
	} else if validations.Params != nil {
		parseErrs := validations.Params.Parse(paramsMap, values.Params)
		cr.setValidateErr(parseErrs, vErr, VALIDATE_PARAMS_ERR)
	}

	return
}

// Retorna um map dos parâmetros passados na requisição
func (cr *CRequest) getPathParams() map[string]any {
	var pathParams = make(map[string]any)

	handlerParams := strings.SplitSeq(cr.path, "/")

	for hParam := range handlerParams {
		if len(hParam) > 0 && string(hParam[0]) == ":" {
			paramName := string(hParam[1:])
			paramValue := cr.getParam(paramName)
			pathParams[paramName] = paramValue
		}
	}

	return pathParams
}

// Adiciona todos os erros de validação presentes em "parseErrs"
// em um novo MsgErrors em vErr, setando sua mensagem com "msgError"
func (cr * CRequest) setValidateErr(
	parseErrs zog.ZogIssueList,
	vErr *lerror.ValueError,
	msgError string,
) {
	if parseErrs == nil {
		return
	}

	var errs = make([]error, len(parseErrs))

	for index, issue := range parseErrs {
		errs[index] = fmt.Errorf(
			"%s: %s",
			strings.Join(issue.Path, "."),
			issue.Message)
	}

	vErr.AppendErr(msgError, errs...)
}
