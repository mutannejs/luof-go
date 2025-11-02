package ltests

import (
    "github.com/google/uuid"
)

type Identifiable interface {
    GetUid() uuid.UUID
}
