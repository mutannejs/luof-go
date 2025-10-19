package repository

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/google/uuid"
)

type ReadLink interface {
    Exists(uuid.UUID) (bool, error)
    GetByUid(uuid.UUID) (domain.Link, error)
}

type WriteLink interface {
    Create(domain.Link) error
    Delete(uuid.UUID) error
    Update(uuid.UUID, domain.Link) error
}

type Link interface {
    ReadLink
    WriteLink
}
