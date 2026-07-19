package lerror

import (
	"errors"
)

var (
	BAD_REQUEST = errors.New("400")
	NOT_FOUND = errors.New("404")
	INTERNAL_SERVER_ERROR = errors.New("500")
)

func BadRequest(err *error) {
	setError(BAD_REQUEST, err)
}

func NotFound(err *error) {
	setError(NOT_FOUND, err)
}

func Internal(err *error) {
	setError(INTERNAL_SERVER_ERROR, err)
}

func setError(code error, err *error) {
	if *err != nil {
		*err = errors.Join(code, *err)
	}
}
