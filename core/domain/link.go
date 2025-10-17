package domain

import (
    "github.com/google/uuid"
    "time"
)

type Link struct {
    Uid uuid.UUID
    Url string
    Name string
    Description Description
    CreatedAt time.Time
    UpdatedAt time.Time
}

type LinkDMService interface {
    Create(string, string, string, bool, uuid.UUID, bool) (uuid.UUID, error)
    Delete(uuid.UUID) (bool, error)
    Update(uuid.UUID, string, string, string, bool) (bool, error)
}

type LinkQueryService interface {
    GetByCategory(uuid.UUID) ([]domain.Link, error)
    GetById(uuid.UUID) (domain.Link, error)
}

type LinkService interface {
    LinkDMService
    LinkQueryService
}
