package lerror

import (
	"errors"
	"fmt"
)

var (
	BAD_REQUEST = errors.New("400")
	NOT_FOUND = errors.New("404")
	CONFLICT = errors.New("409")
	INTERNAL_SERVER_ERROR = errors.New("500")
)

// 400: Bad Request

func SetBadRequest(err *error) {
	setError(BAD_REQUEST, err)
}

func GetBadRequest(err error) error {
	setError(BAD_REQUEST, &err)
	return err
}

func GetBadRequestf(format string, errs ...error) (err error) {
	return getErrorf(BAD_REQUEST, format, errs)
}

// 404: Not Found

func SetNotFound(err *error) {
	setError(NOT_FOUND, err)
}

func GetNotFound(err error) error {
	setError(NOT_FOUND, &err)
	return err
}

func GetNotFoundf(format string, errs ...error) (err error) {
	return getErrorf(NOT_FOUND, format, errs)
}

// 409: Conflict

func SetConflict(err *error) {
	setError(CONFLICT, err)
}

func GetConflict(err error) error {
	setError(CONFLICT, &err)
	return err
}

func GetConflictf(format string, errs ...error) (err error) {
	return getErrorf(CONFLICT, format, errs)
}

// 500: Internal Server Error

func SetInternal(err *error) {
	setError(INTERNAL_SERVER_ERROR, err)
}

func GetInternal(err error) error {
	setError(INTERNAL_SERVER_ERROR, &err)
	return err
}

func GetInternalf(format string, errs ...error) (err error) {
	return getErrorf(INTERNAL_SERVER_ERROR, format, errs)
}

// funções auxiliares

func getErrorf(code error, format string, errs []error) (err error) {
	err = fmt.Errorf(format, errs)
	setError(code, &err)
	return
}

func setError(code error, err *error) {
	if *err != nil {
		*err = errors.Join(code, *err)
	}
}
