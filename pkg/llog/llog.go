package llog

import (
    "fmt"
)

const (
    SQL_PREFIX  = "SQL >>"
)

func PrintSQLError(err error) {
    if err != nil {
        fmt.Println(SQL_PREFIX, err)
    }
}
