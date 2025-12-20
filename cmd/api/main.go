package main

import (
	"database/sql"
	"errors"

	"github.com/mutannejs/luof-go/adapters"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/lmigration"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
    var (
        env map[string]string
        err error
        db *sql.DB
        repoSettings adapters.RepositorySettings
        apiSettings adapters.APISettings
    )

    zerolog.TimeFieldFormat = zerolog.TimestampFunc().Format("HH:mm:ss")
    env, err = lenv.Load(false)

    if err != nil {
        log.Error().Err(err).Msg("error loading env")
        return
    } else {
        log.Info().Msg("env loaded successfully")
    }

    repoSettings, err = adapters.GetRepositorySettings(env)

    if err != nil {
        log.Error().Err(err).Msg("error loading repository")
        return
    } else {
        log.Info().Msg("repository loaded successfully")
    }

    db, err = repoSettings.GetConnection(env)

    if err != nil {
        log.Error().Err(err).Msg("error loading db")
        return
    } else {
        log.Info().Msg("db loaded successfully")
    }

    defer db.Close()

    log.Info().Msg("running migrations...")
    err = lmigration.Up(db, repoSettings.GetMigration)

    if err != nil && !errors.Is(err, lmigration.ErrNoChange) {
        log.Error().Err(err).Msg("error running migrations")
        return
    } else {
        log.Info().Msg("migrations completed successfully")
    }

    apiSettings, err = adapters.GetAPISettings(env)

    if err != nil {
        log.Error().Err(err).Msg("error loading api")
        return
    } else {
        log.Info().Msg("api loaded successfully")
    }

    if apiSettings.StartServer(env) != nil {
        log.Error().Err(err).Msg("error start server api")
    }
}
