package custom

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"
	"github.com/mutannejs/luof-go/pkg/luuid"

	"github.com/Oudwins/zog"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ParamsErrors map[string]string

type ResponseError struct {
	Message string `json:"message"`
	Errors ParamsErrors `json:"errors"`
}

type RequestValues struct {
	JsonBody any
	Params any
}

type RequestValidations struct {
	JsonBody *zog.StructSchema
	Params *zog.StructSchema
}

type Context struct {
	echo.Context
	Repositories repository.Repositories
	Rv RequestValues
	errorResp *ResponseError
	logUid string
}

const (
	LOG_UID_ERR = "error generating new log_uid"
	JSON_BODY_ERR = "the request body could not be interpreted in JSON format"
	PARAMS_ERR = "the path params could not be converted in JSON format"
	VALIDATE_ERR = "errors occurred during the validation of the request parameters"
)

func (cc *Context) ErrLog() *zerolog.Event {
	return log.Error().Str("log_uid", cc.logUid)
}

func (cc *Context) InfoLog() *zerolog.Event {
	return log.Info().Str("log_uid", cc.logUid)
}

func (cc *Context) LogAndReturnErr(err error) error {
	var statusCodeAndMessages = err.(interface{ Unwrap() []error }).Unwrap()

	if statusCodeAndMessages == nil {
		cc.ErrLog().Err(err).Send()
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	var statusCode, errConv = strconv.Atoi(statusCodeAndMessages[0].Error())

	if errConv != nil {
		cc.ErrLog().Err(err).Send()
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	var errContent = statusCodeAndMessages[1]

	cc.ErrLog().Err(errContent).Send()
	return echo.NewHTTPError(statusCode, errContent)
}

func (cc *Context) ExecRequetParamsOperations(
	paramsValue any,
	validation **zog.StructSchema,
) error {
	return cc.ExecRequetOperations(
		RequestValues{ Params: paramsValue },
		RequestValidations{ Params: *validation },
	)
}

func (cc *Context) ExecRequetJSONOperations(
	jsonValue any,
	validation **zog.StructSchema,
) error {
	return cc.ExecRequetOperations(
		RequestValues{ JsonBody: jsonValue },
		RequestValidations{ JsonBody: *validation },
	)
}

func (cc *Context) ExecRequetOperations(
	values RequestValues,
	validations RequestValidations,
) error {
	bodyByteSlice := cc.setJsonBody(values, validations)
	paramsByteSlice := cc.setparamsByteSlice(values, validations)

	cc.logRequest(bodyByteSlice, paramsByteSlice)

	if cc.errorResp != nil {
		return echo.NewHTTPError(http.StatusBadRequest, cc.errorResp)
	}

	cc.Rv = values
	return nil
}

func (cc *Context) setJsonBody(
	values RequestValues, 
	validations RequestValidations,
) (
	bodyByteSlice []byte,
) {
	if values.JsonBody == nil {
		return
	}

	var jsonBody = make(map[string]any)

	jsonBodyErr := json.NewDecoder(cc.Request().Body).Decode(&jsonBody)
	bodyByteSlice, bodyByteSliceErr := json.Marshal(jsonBody)

	if jsonBodyErr != nil || bodyByteSliceErr != nil {
		cc.errorResp = &ResponseError{}
		cc.errorResp.Message = JSON_BODY_ERR
	} else if validations.JsonBody != nil {
		parseErrs := validations.JsonBody.Parse(jsonBody, values.JsonBody)
		cc.setValidateErr(parseErrs)
	}

	return
}

func (cc *Context) setparamsByteSlice(
	values RequestValues, 
	validations RequestValidations,
) (
	paramsByteSlice []byte,
) {
	if values.Params == nil {
		return
	}

	pathParamsMap := cc.getPathParams()

	paramsByteSlice, paramsByteSliceErr := json.Marshal(pathParamsMap)

	if paramsByteSliceErr != nil {
		cc.errorResp = &ResponseError{}
		cc.errorResp.Message = PARAMS_ERR
	} else if validations.Params != nil {
		parseErrs := validations.Params.Parse(pathParamsMap, values.Params)
		cc.setValidateErr(parseErrs)
	}

	return
}

func (cc *Context) getPathParams() map[string]any {
	var pathParams = make(map[string]any)

	handlerParams := strings.Split(cc.Path(), "/")

	for _, hParam := range handlerParams {
		if len(hParam) > 0 && string(hParam[0]) == ":" {
			paramName := string(hParam[1:])
			paramValue := cc.Param(paramName)
			pathParams[paramName] = paramValue
		}
	}

	return pathParams
}

func (cc *Context) setValidateErr(parseErrs zog.ZogIssueList) {
	if parseErrs == nil {
		return
	}

	cc.errorResp = &ResponseError{}
	cc.errorResp.Message = VALIDATE_ERR

	errs := make(map[string]string)

	for _, issue := range parseErrs {
		errs[strings.Join(issue.Path, ".")] = issue.Message
	}

	cc.errorResp.Errors = errs
}

func (cc *Context) logRequest(
	bodyByteSlice []byte,
	paramsByteSlice []byte,
) {
	var logReq *zerolog.Event

	if uid, err := luuid.New(); err != nil {
		err = lerror.GetInternalf("%s: %w", errors.New(LOG_UID_ERR), err)
		cc.LogAndReturnErr(err)
	} else {
		cc.logUid = uid.String()
	}

	if cc.errorResp != nil {
		logReq = cc.ErrLog()
	} else {
		logReq = cc.InfoLog()
	}

	logReq = logReq.
		Str("method", cc.Request().Method).
		Str("path", cc.Request().URL.Path)

	if len(bodyByteSlice) != 0 {
		logReq = logReq.RawJSON("json_body", bodyByteSlice)
	}

	if len(paramsByteSlice) != 0 {
		logReq = logReq.RawJSON("params_path", paramsByteSlice)
	}

	if cc.errorResp != nil {
		errorsByteSlice, _ := json.Marshal(cc.errorResp.Errors)
		logReq = logReq.RawJSON("errors", errorsByteSlice)
		logReq.Msg(cc.errorResp.Message)
	} else {
		logReq.Send()
	}
}
