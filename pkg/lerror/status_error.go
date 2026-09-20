package lerror

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// 400: Bad Request

const BAD_REQUEST = 400

func GetBadRequests(errMsg *i18n.Message, errors ...string) ValueError {
	return getErrors(BAD_REQUEST, errMsg, errors...)
}

// 404: Not Found

const NOT_FOUND = 404

func GetNotFound(errMsg *i18n.Message) ValueError {
	return getError(NOT_FOUND, errMsg)
}

// 409: Conflict

const CONFLICT = 409

func GetConflict(errMsg *i18n.Message) ValueError {
	return getError(CONFLICT, errMsg)
}

// 500: Internal Server Error

const INTERNAL_SERVER_ERROR = 500

func GetInternals(errMsg *i18n.Message, errors ...error) ValueError {
	var errs []string = make([]string, len(errors))
	for i, err := range errors {
		errs[i] = err.Error()
	}
	return getErrors(INTERNAL_SERVER_ERROR, errMsg, errs...)
}

// funções auxiliares

func getError(code int, errMsg *i18n.Message) ValueError {
	return getErrors(code, errMsg)
}

func getErrors(code int, errMsg *i18n.Message, errors ...string) ValueError {
	var errs = make([]string, 0)

	if errors != nil {
		errs = errors
	}

	return ValueError{
		code,
		[]MsgErrors{
			{errMsg, errs},
		}}
}
