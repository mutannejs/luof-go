package repository

import (
	"time"

	"github.com/mutannejs/luof-go/core/domain"

	"github.com/google/uuid"
)

type ReadBelongsTo interface {
    Exists(uuid.UUID, uuid.UUID) (bool, error)
    GetLinksByCategory(uuid.UUID) ([]domain.Link, error)
}

type WriteBelongsTo interface {
    Create(uuid.UUID, uuid.UUID, time.Time, bool) error
    Delete(uuid.UUID, uuid.UUID) error
    Update(uuid.UUID, uuid.UUID, bool) error
}

type BelongsTo interface {
    ReadBelongsTo
}
