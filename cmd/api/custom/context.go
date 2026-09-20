package custom

import (
	"github.com/mutannejs/luof-go/cmd/api/custom/custom_log"
	"github.com/mutannejs/luof-go/cmd/api/custom/custom_request"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"
	"github.com/mutannejs/luof-go/pkg/luuid"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	LOG_UID_ERR = &i18n.Message{
		ID: "LOG_UID_ERR",
		Description: "Retornado quando uuid.New falha",
		Other: "error generating new log_uid",
	}
)

type Context struct {
	echo.Context
	Repositories repository.Repositories
	Log custom_log.CLog
	Localizer *i18n.Localizer
}

func (cc *Context) Init() *custom_request.CRequest {
	var err = cc.initLog()

	return custom_request.New(
		&cc.Log,
		cc.Request().Method,
		cc.Path(),
		cc.Request().URL.Path,
		cc.Request().Body,
		cc.Param,
		cc.JSON,
		err)
}

func (cc *Context) LogRoute() error {
	if err := cc.initLog(); err != nil {
		return err
	}

	cc.Log.LogRoute(
		cc.Request().URL.Path,
		cc.Request().Method)

	return nil
}

func (cc *Context) initLog() error {
	cc.Log = custom_log.CLog{}

	var uid uuid.UUID
	var err error

	if uid, err = luuid.New(); err != nil {
		vErr := lerror.GetInternals(LOG_UID_ERR, err)
		err = cc.Log.ReturnErr(vErr)
	} else {
		cc.Log.SetUid(uid.String())
	}

	return err
}
