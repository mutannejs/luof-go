//go:build !luuid_error

package luuid

import (
    "github.com/google/uuid"
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
