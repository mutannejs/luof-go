package repository

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
)

type SettingsValues struct {
    GetConnection func (map[string]string) (*sql.DB, error)
    GetMigration func (db *sql.DB) (m *migrate.Migrate, err error)
}
