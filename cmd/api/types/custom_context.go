package types

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/mutannejs/luof-go/core/repository"

    "github.com/Oudwins/zog"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type CustomContext struct {
    echo.Context
    Repositories repository.Repositories
    Rv RequestValues
}

const (
	JSON_BODY_ERR = "the request body could not be interpreted in JSON format"
	VALIDATE_ERR = "errors occurred during the validation of the request parameters"
)

func (cc *CustomContext) ExecRequetParamsOperations(
	paramsValue any,
	validation **zog.StructSchema,
) error {
	return cc.ExecRequetOperations(
        RequestValues{ Params: paramsValue },
        RequestValidations{ JsonBody: *validation },
    )
}

func (cc *CustomContext) ExecRequetJSONOperations(
	jsonValue any,
	validation **zog.StructSchema,
) error {
	return cc.ExecRequetOperations(
        RequestValues{ JsonBody: jsonValue },
        RequestValidations{ JsonBody: *validation },
    )
}

func (cc *CustomContext) ExecRequetOperations(
	values RequestValues,
	validations RequestValidations,
) error {
	var errorResp *ResponseError

	bytesBody, errorResp := cc.setJsonBody(values, validations)

	cc.logRequest(bytesBody, errorResp)

	if errorResp != nil {
        return echo.NewHTTPError(http.StatusBadRequest, errorResp)
	}

	cc.Rv = values
	return nil
}

func (cc *CustomContext) setJsonBody(
	values RequestValues, 
	validations RequestValidations,
) (
	bytesBody []byte,
	err *ResponseError,
) {
	var errorResp = ResponseError{}

    var jsonBody = make(map[string]any)
    var jsonBodyErr, bytesBodyErr error
    var parseErrs zog.ZogIssueList

    jsonBodyErr = json.NewDecoder(cc.Request().Body).Decode(&jsonBody)
	bytesBody, bytesBodyErr = json.Marshal(jsonBody)

    if jsonBodyErr != nil || bytesBodyErr != nil {
    	errorResp.Message = JSON_BODY_ERR
    	err = &errorResp
    } else if (validations.JsonBody != nil) {
        parseErrs = validations.JsonBody.Parse(jsonBody, values.JsonBody);
    }

    if parseErrs != nil {
    	errorResp.Message = VALIDATE_ERR
	    errorResp.Errors = cc.getErrors(parseErrs)
    	err = &errorResp
    }

    return
}

func (cc *CustomContext) getErrors(
	parseErrs zog.ZogIssueList,
) (errs ParamsErrors) {
    if parseErrs != nil {
		errs = make(map[string]string)

        for _, issue := range parseErrs {
        	errs[strings.Join(issue.Path, ".")] = issue.Message
        }
	}

	return
}

func (cc *CustomContext) logRequest(
	bytesBody []byte,
	errorResp *ResponseError,
) {
	var logReq *zerolog.Event

    if errorResp != nil {
    	logReq = log.Error()
    } else {
    	logReq = log.Info()
    }

    logReq = logReq.
        Str("path", cc.Request().URL.Path).
        Str("method", cc.Request().Method).
    	RawJSON("json_body", bytesBody)

    if errorResp != nil {
    	for key, value := range errorResp.Errors {
			logReq = logReq.Str(key, value)
		}

		logReq.Msg(errorResp.Message)
    } else {
		logReq.Send()
	}
}
