package main

import (
	"database/sql"
	"errors"
	"time"

	"github.com/mutannejs/luof-go/adapters"
	"github.com/mutannejs/luof-go/cmd/api/middleware"
	"github.com/mutannejs/luof-go/cmd/api/route"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/lmigration"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	DEFAULT_PORT = "8123"
)

func main() {
	var (
		env map[string]string
		err error
		db *sql.DB
		repoSettings adapters.RepositorySettings
	)

	zerolog.TimeFieldFormat = time.DateTime
	env, err = lenv.Load()

	if err != nil {
		log.Error().Err(err).Msg("error loading env")
		return
	} else {
		log.Info().Msg("env loaded successfully in '" + env["ENV"] + "'")
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
		defer db.Close()
	}

	log.Info().Msg("running migrations...")
	err = lmigration.Up(db, repoSettings.GetMigration)

	if err != nil && !errors.Is(err, lmigration.ErrNoChange) {
		log.Error().Err(err).Msg("error running migrations")
		return
	} else {
		log.Info().Msg("migrations completed successfully")
	}

	if startServer(env, repoSettings.GetRepositories(db)) != nil {
		log.Error().Err(err).Msg("error start server api")
	}
}

func startServer(env map[string]string, repositories repository.Repositories) error {
	var address string

	if envPort, exists := env["SERVER_PORT"]; exists {
		address = ":" + envPort
	} else {
		address = ":" + DEFAULT_PORT
	}

	var e *echo.Echo = echo.New()

	e.HideBanner = true

	middleware.SetMiddleware(e, repositories)
	route.SetRootRoutes(e)

	return e.Start(address)
}
