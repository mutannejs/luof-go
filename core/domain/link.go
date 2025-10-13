package domain

import (
    "time"
    "github.com/google/uuid"
)

type Link struct {
    Uid uuid.UUID
    Url string
    Name string
    Description string
    CreatedAt time.Time
    UpdatedAt time.Time
}
