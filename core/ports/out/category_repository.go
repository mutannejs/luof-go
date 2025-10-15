package out

import (
    "github.com/mutannejs/luof-go/core/domain"

    "github.com/google/uuid"
)

type CategoryRepository interface {
    Create (domain.Category) error
    Delete (uuid.UUID) error
    Update (uuid.UUID, domain.Category) error
    Exists (uuid.UUID) (bool, error)
}
