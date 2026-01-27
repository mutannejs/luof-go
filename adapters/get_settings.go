package adapters

import (
	"errors"

	"github.com/mutannejs/luof-go/adapters/sqlite"
	"github.com/mutannejs/luof-go/adapters/sqlite/get_repositories"
)

var (
	// DB var
	ENV_DB_NOT_FOUND = errors.New("the 'DB' env var not found")
	ENV_DB_INVALID = errors.New("the 'DB' env var is invalid, it should be: sqlite")
)

func GetRepositorySettings(env map[string]string) (values RepositorySettings, err error) {
	var db, exists = env["DB"]

	if !exists {
		err = ENV_DB_NOT_FOUND
		return
	}

	switch db {
		case "sqlite":
			values = RepositorySettings{
				GetConnection: sqlite.GetConnection,
				GetMigration: sqlite.GetMigration,
				GetRepositories: get_repositories.GetRepositories,
			}
		default: err = ENV_DB_INVALID
	}

	return
}
