package domain

import (
	"time"

	"github.com/mutannejs/luof-go/pkg/lerror"
	"github.com/mutannejs/luof-go/pkg/luuid"

	"github.com/google/uuid"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	LINK_ERROR_NEW = &i18n.Message{
		ID: "LINK_ERROR_NEW",
		Description: "Retornado quando uuid.New falha",
		Other: "error instantiate new link",
	}
	LINK_NOT_EXISTS = &i18n.Message{
		ID: "LINK_NOT_EXISTS",
		Description: "Retornado quando consultado um link que não existe",
		Other: "the searched link does not exist",
	}
)

type Link struct {
	uid uuid.UUID
	Url string
	Name string
	Description Description
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (l Link) GetUid() uuid.UUID {
	return l.uid
}

func (l *Link) SetUid(uid uuid.UUID) {
	l.uid = uid
}

func NewLink(
	url string,
	name string,
	contentDescription string,
	useMarkdown bool,
) (link Link, vError lerror.ValueError) {
	var uid uuid.UUID
	var err error

	uid, err = luuid.New()
	if err != nil {
		vError = lerror.GetInternals(LINK_ERROR_NEW, err)
		return
	}

	var createdAt time.Time = time.Now()
	var updatedAt time.Time
	var description = Description{contentDescription, useMarkdown}
	link = Link{uid, url, name, description, createdAt, updatedAt}

	return
}
