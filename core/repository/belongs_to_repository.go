package repository

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/google/uuid"
)

type ReadBelongsTo interface {
    GetLinksByCategory(uuid.UUID) ([]domain.Link, error)
}

type BelongsTo interface {
    ReadBelongsTo
}
