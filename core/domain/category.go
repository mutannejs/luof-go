package domain

import (
    "github.com/google/uuid"
    "time"
)

type Category struct {
    Uid uuid.UUID
    Name string
    Description Description
    CreatedAt time.Time
    UpdatedAt time.Time
}

type CategoryDMService interface {
    Create(string, string, bool) (uuid.UUID, error)
    Delete(uuid.UUID) (bool, error)
    Update(uuid.UUID, string, string, bool) (bool, error)
}

type CategoryQueryService interface {
    GetById(uuid.UUID) (domain.Category, error)
}

type CategoryService interface {
    CategoryDMService
    CategoryQueryService
}
