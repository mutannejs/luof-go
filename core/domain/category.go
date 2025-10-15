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
