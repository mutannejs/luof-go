package llog

import (
    "fmt"
)

const (
    SQL_PREFIX  = "SQL >>"
    UUID_PREFIX = "UUID >>"
    UUID_ERROR_MESSAGE = "error generating uuid"
)

func PrintSQLError (err error) {
    if err != nil {
        fmt.Println(SQL_PREFIX, err)
    }
}
