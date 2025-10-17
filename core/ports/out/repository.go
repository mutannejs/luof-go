package out

import (
    "github.com/mutannejs/luof-go/core/domain"

    "github.com/google/uuid"
)

type LinkRead interface {
    Exists(uuid.UUID) (bool, error)
    GetByCategory(uuid.UUID) ([]domain.Link, error)
    GetById(uuid.UUID) (domain.Link, error)
}

type LinkWrite interface {
    Create(domain.Link) error
    CreateWithCategory(link, uuid.UUID, bool, time.Time) error
    Delete(uuid.UUID) error
    Update(uuid.UUID, domain.Link) error
}

type LinkRepository interface {
    LinkRead
    LinkWrite
}
