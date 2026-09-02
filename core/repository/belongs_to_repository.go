package repository

import (
	"time"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
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
