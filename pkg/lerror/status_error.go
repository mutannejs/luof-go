package lerror

// 400: Bad Request

const BAD_REQUEST = 400

// 404: Not Found

const NOT_FOUND = 404

func GetNotFound(errMsg string) ValueError {
	return getError(BAD_REQUEST, errMsg)
}

// 409: Conflict

const CONFLICT = 409

func GetConflict(errMsg string) ValueError {
	return getError(CONFLICT, errMsg)
}

// 500: Internal Server Error

const INTERNAL_SERVER_ERROR = 500

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
