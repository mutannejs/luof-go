package luuid

import (
    "github.com/mutannejs/luof-go/pkg/llog"

    "errors"
    "fmt"
    "github.com/google/uuid"
)

func New () (id uuid.UUID, err error) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println(llog.UUID_PREFIX, llog.UUID_ERROR_MESSAGE)
            err = errors.New(llog.UUID_ERROR_MESSAGE)
        }
    }()
    id = uuid.New()
    return
}
