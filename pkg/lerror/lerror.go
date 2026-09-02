package lerror

var (
	BAD_REQUEST = 400
	NOT_FOUND = 404
	CONFLICT = 409
	INTERNAL_SERVER_ERROR = 500
)

type MsgErrors struct {
	message string
	errors []error
}

func (m *MsgErrors) GetMessage() string {
	return m.message
}

type ValueError struct {
	code int
	errors []MsgErrors
}

func (v *ValueError) IsNil() bool {
	return len(v.errors) == 0
}

func (v *ValueError) GetErrors() []MsgErrors {
	return v.errors
}

// 400: Bad Request

// 404: Not Found

func GetNotFound(errMsg string) ValueError {
	return getError(BAD_REQUEST, errMsg)
}

// 409: Conflict

func GetConflict(errMsg string) ValueError {
	return getError(CONFLICT, errMsg)
}

// 500: Internal Server Error

func GetInternal(err error) ValueError {
	if err == nil {
		return ValueError{}
	}
	return getError(INTERNAL_SERVER_ERROR, err.Error())
}

func GetInternals(errMsg string, errors ...error) ValueError {
	return getErrors(INTERNAL_SERVER_ERROR, errMsg, errors...)
}

// funções auxiliares

func getError(code int, errMsg string) ValueError {
	return getErrors(code, errMsg)
}

func getErrors(code int, errMsg string, errors ...error) ValueError {
	return ValueError{
		code,
		[]MsgErrors{
			{errMsg, errors},
		}}
}
