package adapters

import (
	"errors"

	"github.com/mutannejs/luof-go/adapters/echo"
	"github.com/mutannejs/luof-go/adapters/sqlite"
)

var (
    // DB var
    ENV_DB_NOT_FOUND = errors.New("the 'DB' env var not found")
    ENV_DB_INVALID = errors.New("the 'DB' env var is invalid, it should be: sqlite")
    // API var
    ENV_API_NOT_FOUND = errors.New("the 'API' env var not found")
    ENV_API_INVALID = errors.New("the 'API' env var is invalid, it should be: echo")
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
            }
        default: err = ENV_DB_INVALID
    }

    return
}

func GetAPISettings(env map[string]string) (values APISettings, err error) {
    var api, exists = env["API"]

    if !exists {
        err = ENV_API_NOT_FOUND
        return
    }

    switch api {
        case "echo":
            values = APISettings{
                StartServer: echo.StartServer,
            }
        default: err = ENV_API_INVALID
    }

    return
}
