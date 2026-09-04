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
	VALIDATE_ERR = "errors occurred during the validation of the request parameters"
)

func (cr *CRequest) setJsonBody(
	values RequestValues, 
	validations RequestValidations,
	vErr *lerror.ValueError,
) (
	bodyByteSlice []byte,
) {
	if values.JsonBody == nil {
		return
	}

	var jsonBody = make(map[string]any)

	jsonBodyErr := json.NewDecoder(cr.body).Decode(&jsonBody)
	bodyByteSlice, bodyByteSliceErr := json.Marshal(jsonBody)

	if jsonBodyErr != nil || bodyByteSliceErr != nil {
		vErr.AppendErr(JSON_BODY_ERR)
	} else if validations.JsonBody != nil {
		parseErrs := validations.JsonBody.Parse(jsonBody, values.JsonBody)
		cr.setValidateErr(parseErrs, vErr)
	}

	return
}

func (cr *CRequest) setParamsByteSlice(
	values RequestValues, 
	validations RequestValidations,
	vErr *lerror.ValueError,
) (
	paramsByteSlice []byte,
) {
	if values.Params == nil {
		return
	}

	pathParamsMap := cr.getPathParams()
	paramsByteSlice, paramsByteSliceErr := json.Marshal(pathParamsMap)

	if paramsByteSliceErr != nil {
		vErr.AppendErr(PARAMS_ERR)
	} else if validations.Params != nil {
		parseErrs := validations.Params.Parse(pathParamsMap, values.Params)
		cr.setValidateErr(parseErrs, vErr)
	}

	return
}

func (cr *CRequest) getPathParams() map[string]any {
	var pathParams = make(map[string]any)

	handlerParams := strings.Split(cr.path, "/")

	for _, hParam := range handlerParams {
		if len(hParam) > 0 && string(hParam[0]) == ":" {
			paramName := string(hParam[1:])
			paramValue := cr.getParam(paramName)
			pathParams[paramName] = paramValue
		}
	}

	return pathParams
}

func (cr * CRequest) setValidateErr(parseErrs zog.ZogIssueList, vErr *lerror.ValueError) {
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

	vErr.AppendErr(VALIDATE_ERR, errs...)
}
