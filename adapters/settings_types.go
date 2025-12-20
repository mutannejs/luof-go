package adapters

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
)

type RepositorySettings struct {
    GetConnection func (map[string]string) (*sql.DB, error)
    GetMigration func (db *sql.DB) (m *migrate.Migrate, err error)
}

type APISettings struct {
    StartServer func (map[string]string) error
}
