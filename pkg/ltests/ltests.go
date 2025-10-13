package ltests

import (
    "database/sql"
    "fmt"
    "os"
    "testing"
    _ "github.com/mattn/go-sqlite3"
)

const (
    DB_PATH = "./luof_test.db"
    DB_DRIVER = "sqlite3"
    PRINT_PREFIX = "LTests >>"
)

type LTests struct {
    T *testing.T
    DB *sql.DB
}

func NewTest(t *testing.T) (*LTests) {
    return &LTests{t, nil}
}

func NewDataBaseTest(t *testing.T) (*LTests, *sql.DB) {
    const PERMISSIONS = 0666
    var content []byte
    err := os.WriteFile(DB_PATH, content, PERMISSIONS)
    if err != nil {
        fmt.Println(err)
        t.Fail()
    }
    db, err := sql.Open(DB_DRIVER, DB_PATH)
    if err != nil {
        fmt.Println(err)
        t.Fail()
    }
    return &LTests{t, db}, db
}

func (tests *LTests) CloseDataBaseTest() {
    tests.DB.Close()
    os.Remove(DB_PATH)
}

func (tests *LTests) FailIfError(err error) {
    if err != nil {
        fmt.Println(PRINT_PREFIX, err)
        tests.T.Fail()
    }
}

func (tests *LTests) ExistsTable(tableName string) (bool, error) {
    const EXISTS = 1
    var result int
    existsTableQuery := `SELECT EXISTS(SELECT 1 FROM sqlite_master WHERE type = 'table' AND name = ?);`
    err := tests.DB.QueryRow(existsTableQuery, tableName).Scan(&result)
    return result == EXISTS, err
}

func (tests *LTests) PrintAndFail(message string) {
    fmt.Println(PRINT_PREFIX, message)
    tests.T.Fail()
}
