package main

import (
	"database/sql"
	"fmt"

	"github.com/mutannejs/luof-go/adapters"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/lmigration"
)

func main() {
    var (
        env map[string]string
        err error
        db *sql.DB
        repoSettings adapters.RepositorySettings
        apiSettings adapters.APISettings
    )

    env, err = lenv.Load(true)

    fmt.Println(env, err)

    if err == nil {
        repoSettings, err = adapters.GetRepositorySettings(env)
    }

    fmt.Println(repoSettings, err)

    if err == nil {
        db, err = repoSettings.GetConnection(env)
        if err == nil {
            defer db.Close()
        }
    }

    fmt.Println(db, err)

    if err == nil {
        lmigration.Up(db, repoSettings.GetMigration)
    }

    if err == nil {
        apiSettings, err = adapters.GetAPISettings(env)
    }

    fmt.Println(apiSettings, err)

    if apiSettings.StartServer(env) != nil {
        fmt.Println("Ocorreu um erro");
    }
}
