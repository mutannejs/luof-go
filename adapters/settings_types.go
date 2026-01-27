package adapters

import (
	"database/sql"

	"github.com/mutannejs/luof-go/core/repository"

	"github.com/golang-migrate/migrate/v4"
)

type RepositorySettings struct {
	GetConnection func (map[string]string) (*sql.DB, error)
	GetMigration func (db *sql.DB) (m *migrate.Migrate, err error)
	GetRepositories func (db *sql.DB) repository.Repositories
}

type APISettings struct {
	StartServer func (map[string]string, repository.Repositories) error
}
