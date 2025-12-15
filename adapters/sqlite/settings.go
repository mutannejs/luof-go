package sqlite

import (
	"database/sql"
	"os"

	"github.com/mutannejs/luof-go/pkg/lpath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "modernc.org/sqlite"
)

func GetConnection(env map[string]string) (db *sql.DB, err error) {
    var path = env["SQLITE_DB_PATH"]

    if env["ENV"] == "test" {
        path += ".test"
    }

    path += ".sqlite"

    var absolutePath = lpath.GetAbsolutetPath(path)

    if _, err = os.OpenFile(absolutePath, os.O_CREATE, 0644); err != nil {
        return
    }

    return sql.Open("sqlite", absolutePath + "?_pragma=foreign_keys=true")
}

func GetMigration(db *sql.DB) (m *migrate.Migrate, err error) {
    var driver database.Driver

    driver, err = sqlite.WithInstance(db, &sqlite.Config{})

    if err == nil {
        m, err = migrate.NewWithDatabaseInstance(
            "file://" + lpath.ROOT_PATH + "/migrations/sqlite",
            "sqlite", driver)
    }

    return
}
