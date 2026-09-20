package repository

import (
	"time"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/google/uuid"
)

var (
	READ_BELONGS_TO_EXISTS_ERROR = &i18n.Message{
		ID: "READ_BELONGS_TO_EXISTS_ERROR",
		Description: "Retornado quando belongs_to_repository.Exists falha",
		Other: "error verifying in database if the link belongs to the category",
	}
	READ_BELONGS_TO_GET_LINKS_BY_CATEGORY_ERROR = &i18n.Message{
		ID: "READ_BELONGS_TO_GET_LINKS_BY_CATEGORY_ERROR",
		Description: "Retornado quando belongs_to_repository.GetLinksByCategory falha",
		Other: "error querying links belonging to category in database",
	}
	READ_BELONGS_TO_HAS_LINKS_ERROR = &i18n.Message{
		ID: "READ_BELONGS_TO_HAS_LINKS_ERROR",
		Description: "Retornado quando belongs_to_repository.HasLinks falha",
		Other: "error verifying in database if the category has links",
	}
	WRITE_BELONGS_TO_CREATE_ERROR = &i18n.Message{
		ID: "WRITE_BELONGS_TO_CREATE_ERROR",
		Description: "Retornado quando belongs_to_repository.Create falha",
		Other: "error inserting link in category in database",
	}
	WRITE_BELONGS_TO_DELETE_ERROR = &i18n.Message{
		ID: "WRITE_BELONGS_TO_DELETE_ERROR",
		Description: "Retornado quando belongs_to_repository.Delete falha",
		Other: "error removing link from category in database",
	}
	WRITE_BELONGS_TO_UPDATE_ERROR = &i18n.Message{
		ID: "WRITE_BELONGS_TO_UPDATE_ERROR",
		Description: "Retornado quando belongs_to_repository.Update falha",
		Other: "error updating the relationship between link and category in database",
	}
	WRITE_BELONGS_TO_SET_HAS_NO_MAIN_CATEGORY_ERROR = &i18n.Message{
		ID: "WRITE_BELONGS_TO_SET_HAS_NO_MAIN_CATEGORY_ERROR",
		Description: "Retornado quando belongs_to_repository.SetHasNoMainCategory falha",
		Other: "error seting all relationships between link and categories as non-main in database",
	}
)

type ReadBelongsTo interface {
	Exists(uuid.UUID, uuid.UUID) (bool, lerror.ValueError)
	GetLinksByCategory(uuid.UUID) ([]domain.Link, lerror.ValueError)
	HasLinks(uuid.UUID) (bool, lerror.ValueError)
}

type WriteBelongsTo interface {
	Create(uuid.UUID, uuid.UUID, time.Time, bool) lerror.ValueError
	Delete(uuid.UUID, uuid.UUID) lerror.ValueError
	Update(uuid.UUID, uuid.UUID, bool) lerror.ValueError
	SetHasNoMainCategory(uuid.UUID) lerror.ValueError
}

type BelongsTo interface {
	ReadBelongsTo
	WriteBelongsTo
}
