package luuid

import (
    "errors"
    "github.com/google/uuid"
    "reflect"
)

var (
    UUID_ERROR_NEW = errors.New("error generating new uuid")
)

func New() (uid uuid.UUID, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = UUID_ERROR_NEW
        }
    }()
    uid = uuid.New()
    return
}

func Zero() uuid.UUID {
    var zero uuid.UUID
    return zero
}

func IsZero(uid uuid.UUID) bool {
    var zero uuid.UUID
    return reflect.DeepEqual(uid, zero)
}
