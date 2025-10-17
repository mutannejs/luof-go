package luuid

import (
    "github.com/mutannejs/luof-go/pkg/llog"

    "errors"
    "fmt"
    "github.com/google/uuid"
    "reflect"
)

func New() (uid uuid.UUID, err error) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println(llog.UUID_PREFIX, llog.UUID_ERROR_MESSAGE)
            err = errors.New(llog.UUID_ERROR_MESSAGE)
        }
    }()
    uid = uuid.New()
    return
}

func IsZero(uid uuid.UUID) bool {
    var zero uuid.UUID
    return reflect.DeepEqual(uid, zero)
}
