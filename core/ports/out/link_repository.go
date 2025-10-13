package out

import (
    "github.com/mutannejs/luof-go/core/domain"

    "github.com/google/uuid"
)

type LinkRepository interface {
    Create (domain.Link) error
    Delete (uuid.UUID) error
    Update (uuid.UUID, domain.Link) error
    Exists (uuid.UUID) (bool, error)
}
