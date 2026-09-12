package custom_log

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mutannejs/luof-go/pkg/lerror"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	LOG_KEY_ERRORS = "errors"
	LOG_KEY_JSON_BODY = "json_body"
	LOG_KEY_METHOD = "method"
	LOG_KEY_PARAMS_PATH = "params_path"
	LOG_KEY_PATH = "path"
	LOG_KEY_STATUS_CODE = "status_code"
	LOG_KEY_UID = "log_uid"
	UNKNOWN_ERROR = "unknown error"
)

/**
 * Armazena um identificador usado para relacionar logs
 * feitos na mesma requisição
 */
type CLog struct {
	logUid string
}

/**
 * Seta logUid
 */
func (l *CLog) SetUid(logUid string) {
	l.logUid = logUid
}

/**
 * Inicia um log de erro identificado pelo identificador
 * armazenado na estrutura
 */
func (l *CLog) ErrLog() *zerolog.Event {
	return log.Error().Str(LOG_KEY_UID, l.logUid)
}

/**
 * Inicia um log de informação identificado pelo identificador
 * armazenado na estrutura
 */
func (l *CLog) InfoLog() *zerolog.Event {
	return log.Info().Str(LOG_KEY_UID, l.logUid)
}

/**
 * Loga o erro passado como argumento e retorna [*echo.HTTPError] com
 * código 500 e o erro informado
 */
func (l *CLog) ReturnInternalErr(err error) error {
	if err == nil {
		err = errors.New(UNKNOWN_ERROR)
	}

	l.ErrLog().
		Int(LOG_KEY_STATUS_CODE, http.StatusInternalServerError).
		Err(err).
		Send()
	return echo.NewHTTPError(http.StatusInternalServerError, err)
}

/**
 * Loga o erro passado como argumento e retorna [*echo.HTTPError] com
 * o código e erro informados
 */
func (l *CLog) ReturnErr(vErr lerror.ValueError) error {
	if vErr.IsNil() {
		return l.ReturnInternalErr(nil)
	}

	errorsByteSlice, err := json.Marshal(vErr.GetErrors())
	if err != nil {
		return l.ReturnInternalErr(err)
	}

	l.ErrLog().
		Int(LOG_KEY_STATUS_CODE, vErr.GetCode()).
		RawJSON(LOG_KEY_ERRORS, errorsByteSlice).
		Send()
	return echo.NewHTTPError(vErr.GetCode(), vErr.GetErrors())
}

/**
 * Loga método e caminho da requisição
 */
func (l *CLog) LogRoute(method, resolvedPath string) {
	l.InfoLog().
		Str(LOG_KEY_METHOD, method).
		Str(LOG_KEY_PATH, resolvedPath).
		Send()
}

/**
 * Loga método, caminho, corpo e parâmetros recebidos na requisição,
 * e qualquer erro que tenha ocorrido
 */
func (l *CLog) LogRequest(
	paramsByteSlice []byte,
	bodyByteSlice []byte,
	method string,
	resolvedPath string,
	vErr lerror.ValueError,
) {
	var logReq *zerolog.Event

	if !vErr.IsNil() {
		logReq = l.ErrLog()
	} else {
		logReq = l.InfoLog()
	}

	logReq = logReq.
		Str(LOG_KEY_METHOD, method).
		Str(LOG_KEY_PATH, resolvedPath)

	if len(bodyByteSlice) != 0 {
		logReq = logReq.RawJSON(LOG_KEY_JSON_BODY, bodyByteSlice)
	}

	if len(paramsByteSlice) != 0 {
		logReq = logReq.RawJSON(LOG_KEY_PARAMS_PATH, paramsByteSlice)
	}

	if !vErr.IsNil() {
		errorsByteSlice, err := json.Marshal(vErr.GetErrors())
		if err != nil {
			logReq.Err(err).Send()
			return
		}

		logReq.
			RawJSON(LOG_KEY_ERRORS, errorsByteSlice).
			Int(LOG_KEY_STATUS_CODE, http.StatusBadRequest).
			Send()
	} else {
		logReq.Send()
	}
}
