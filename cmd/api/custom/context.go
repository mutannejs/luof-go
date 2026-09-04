package custom

import (
	"github.com/google/uuid"
	"github.com/mutannejs/luof-go/cmd/api/custom/custom_log"
	"github.com/mutannejs/luof-go/cmd/api/custom/custom_request"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"
	"github.com/mutannejs/luof-go/pkg/luuid"

	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
	Repositories repository.Repositories
	Log custom_log.CLog
}

const (
	LOG_UID_ERR = "error generating new log_uid"
)

func (cc *Context) Init() *custom_request.CRequest {
	cc.Log = custom_log.CLog{}

	var uid uuid.UUID
	var err error

	if uid, err = luuid.New(); err != nil {
		vErr := lerror.GetInternals(LOG_UID_ERR, err)
		err = cc.Log.ReturnErr(vErr)
	} else {
		cc.Log.SetUid(uid.String())
	}

	return custom_request.New(
		&cc.Log,
		cc.Request().Method,
		cc.Path(),
		cc.Request().Body,
		cc.Param,
		err)
}
