package repository

import (
    "github.com/mutannejs/luof-go/core/domain"

    "github.com/google/uuid"
)

type ReadCategory interface {
    Exists(uuid.UUID) (bool, error)
    GetByUid(uuid.UUID) (domain.Category, error)
}

type WriteCategory interface {
    Create(domain.Category) error
    Delete(uuid.UUID) error
    Update(uuid.UUID, domain.Category) error
}

type Category interface {
    ReadCategory
    WriteCategory
}

var (
    CATEGORY_NOT_EXISTS = "the searched category does not exist"
)
