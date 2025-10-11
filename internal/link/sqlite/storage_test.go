package link_storage

import (
    "testing"
    "github.com/mutannejs/luof-go/pkg/ltests"
)

/**
 * Testa a criação da tabela Link
 */
func TestInitTable(t *testing.T) {
    tests, db := ltests.NewDataBaseTest(t)
    defer tests.CloseDataBaseTest()

    _, err := NewLinkStorage(db)
    tests.FailIfError(err)

    if exists, _ := tests.ExistsTable("link"); !exists {
        tests.PrintAndFail("The Link table was not created")
    }

    _, err = NewLinkStorage(db)
    tests.FailIfError(err)
}
