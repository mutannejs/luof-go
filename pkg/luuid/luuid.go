package luuid

import (
    "errors"
    "github.com/google/uuid"
    "reflect"
)

var (
    UUID_ERROR_NEW = errors.New("error generating new uuid")
)

func Zero() uuid.UUID {
    var zero uuid.UUID
    return zero
}

func IsZero(uid uuid.UUID) bool {
    var zero uuid.UUID
    return reflect.DeepEqual(uid, zero)
}
