package lerror

import (
	"errors"
	"fmt"
)

var (
	BAD_REQUEST = errors.New("400")
	NOT_FOUND = errors.New("404")
	INTERNAL_SERVER_ERROR = errors.New("500")
)

func BadRequest(err *error) {
	setError(BAD_REQUEST, err)
}

func BadRequestf(format string, errs ...error) (err error) {
	return getErrorf(BAD_REQUEST, format, errs)
}

func NotFound(err *error) {
	setError(NOT_FOUND, err)
}

func NotFoundf(format string, errs ...error) (err error) {
	return getErrorf(NOT_FOUND, format, errs)
}

func Internal(err *error) {
	setError(INTERNAL_SERVER_ERROR, err)
}

func Internalf(format string, errs ...error) (err error) {
	return getErrorf(INTERNAL_SERVER_ERROR, format, errs)
}

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
