//go:build luuid_error

package luuid

import (
    "github.com/google/uuid"
)

func New() (uid uuid.UUID, err error) {
    uid = Zero()
    err = UUID_ERROR_NEW
    return
}
