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
