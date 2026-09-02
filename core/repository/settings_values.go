package repository

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/mutannejs/luof-go/pkg/lerror"
)

type Repositories struct {
	Link Link
	Category Category
	BelongsTo BelongsTo
}

type SettingsValues struct {
	GetConnection func (map[string]string) (*sql.DB, lerror.ValueError)
	GetMigration func (*sql.DB) (*migrate.Migrate, lerror.ValueError)
	GetRepositories func (*sql.DB) Repositories
}
