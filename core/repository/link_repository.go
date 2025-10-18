package out

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/google/uuid"
)

type Read interface {
    Exists(uuid.UUID) (bool, error)
    GetById(uuid.UUID) (domain.Link, error)
}

type Write interface {
    Create(domain.Link) error
    Delete(uuid.UUID) error
    Update(uuid.UUID, domain.Link) error
}

type Link interface {
    Read
    Write
}
