package custom_request

import (
	"io"

	"github.com/Oudwins/zog"
	"github.com/mutannejs/luof-go/cmd/api/custom/custom_log"
)

/**
 * Estrutura que grupa valores e métodos para validação
 * e recuperação dos valores enviados na requisição
 */
type CRequest struct {
	log *custom_log.CLog
	method string
	path string
	body io.ReadCloser
	getParam func(name string) string
	sendJson func(code int, i any) error
	err error
}

/**
 * Recebe o endereço das estruturas que receberão os dados
 * enviados na requisição
 */
type RequestValues struct {
	JsonBody any
	Params any
}

/**
 * Define os schemas usados para validar os dados enviados
 * na requisição
 */
type RequestValidations struct {
	JsonBody *zog.StructSchema
	Params *zog.StructSchema
}

/**
 * Instância um objeto do tipo *CRequest
 */
func New(
	log *custom_log.CLog,
	method string,
	path string,
	body io.ReadCloser,
	getParam func(name string) string,
	sendJson func(code int, i any) error,
	err error,
) *CRequest {
	return &CRequest{log, method, path, body, getParam, sendJson, err}
}
