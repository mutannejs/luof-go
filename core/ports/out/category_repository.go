package out

import (
    "github.com/mutannejs/luof-go/core/domain"

    "github.com/google/uuid"
)

type CategoryRead interface {
    Exists(uuid.UUID) (bool, error)
    GetById(uuid.UUID) (domain.Category, error)
}

type CategoryWrite interface {
    Create(domain.Category) error
    Delete(uuid.UUID) error
    Update(uuid.UUID, domain.Category) error
}

type CategoryRepository interface {
    CategoryRead
    CategoryWrite
}
