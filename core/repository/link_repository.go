package repository

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/google/uuid"
)

var (
	READ_LINK_EXISTS_ERROR = &i18n.Message{
		ID: "READ_LINK_EXISTS_ERROR",
		Description: "Retornado quando link_repository.Exists falha",
		Other: "error verifying in database the existence of a link",
	}
	READ_LINK_GET_BY_UID_ERROR = &i18n.Message{
		ID: "READ_LINK_GET_BY_UID_ERROR",
		Description: "Retornado quando link_repository.GetByUid falha",
		Other: "error querying link in database",
	}
	WRITE_LINK_CREATE_ERROR = &i18n.Message{
		ID: "WRITE_LINK_CREATE_ERROR",
		Description: "Retornado quando link_repository.Create falha",
		Other: "error creating link in database",
	}
	WRITE_LINK_DELETE_ERROR = &i18n.Message{
		ID: "WRITE_LINK_DELETE_ERROR",
		Description: "Retornado quando link_repository.Delete falha",
		Other: "error deleting link in database",
	}
	WRITE_LINK_UPDATE_ERROR = &i18n.Message{
		ID: "WRITE_LINK_UPDATE_ERROR",
		Description: "Retornado quando link_repository.Update falha",
		Other: "error updating link in database",
	}
)

type ReadLink interface {
	Exists(uuid.UUID) (bool, lerror.ValueError)
	GetByUid(uuid.UUID) (domain.Link, lerror.ValueError)
}

type WriteLink interface {
	Create(domain.Link) lerror.ValueError
	Delete(uuid.UUID) lerror.ValueError
	Update(uuid.UUID, domain.Link) lerror.ValueError
}

type Link interface {
	ReadLink
	WriteLink
}
