package out

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/google/uuid"
)

type Read interface {
    GetLinksByCategory(uuid.UUID) ([]domain.Link, error)
}

type Link interface {
    Read
}
