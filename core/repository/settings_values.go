package repository

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
)

type Repositories struct {
    Link Link
    Category Category
    BelongsTo BelongsTo
}

type SettingsValues struct {
    GetConnection func (map[string]string) (*sql.DB, error)
    GetMigration func (db *sql.DB) (m *migrate.Migrate, err error)
    GetRepositories func (db *sql.DB) Repositories
}
