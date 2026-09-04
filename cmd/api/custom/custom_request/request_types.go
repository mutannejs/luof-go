package custom_request

import (
	"io"

	"github.com/Oudwins/zog"
	"github.com/mutannejs/luof-go/cmd/api/custom/custom_log"
)

type CRequest struct {
	log *custom_log.CLog
	method string
	path string
	body io.ReadCloser
	getParam func (name string) string
	err error
}

type RequestValues struct {
	JsonBody any
	Params any
}

type RequestValidations struct {
	JsonBody *zog.StructSchema
	Params *zog.StructSchema
}

func New(
	log *custom_log.CLog,
	method string,
	path string,
	body io.ReadCloser,
	getParam func (name string) string,
	err error,
) *CRequest {
	return &CRequest{log, method, path, body, getParam, err}
}
