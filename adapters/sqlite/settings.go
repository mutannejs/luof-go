package sqlite

import (
	"database/sql"
	"errors"
	"os"

	"github.com/mutannejs/luof-go/pkg/lpath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/rs/zerolog/log"
	_ "modernc.org/sqlite"
)

var (
	SQLITE_DB_PATH_NOT_FOUND = errors.New("the 'SQLITE_DB_PATH' env var not found")
)

func GetConnection(env map[string]string) (db *sql.DB, err error) {
	path, exists := env["SQLITE_DB_PATH"]

	if !exists {
		err = SQLITE_DB_PATH_NOT_FOUND
		return
	}

	environmet, exists := env["ENV"]

	if exists {
		path += "." + environmet
	}

	path += ".sqlite"

	var absolutePath = lpath.GetAbsolutetPath(path)

	if _, err = os.OpenFile(absolutePath, os.O_CREATE, 0644); err != nil {
		return
	}

	log.Info().Str("sqlite_db_path", path).Msg("loading db")

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
