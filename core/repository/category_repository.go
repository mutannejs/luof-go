package repositories

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/google/uuid"
)

type Read interface {
    Exists(uuid.UUID) (bool, error)
    GetById(uuid.UUID) (domain.Category, error)
}

type Write interface {
    Create(domain.Category) error
    Delete(uuid.UUID) error
    Update(uuid.UUID, domain.Category) error
}

type Category interface {
    Read
    Write
}
