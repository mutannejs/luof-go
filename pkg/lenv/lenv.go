package lenv

import (
	"errors"

	"github.com/mutannejs/luof-go/pkg/lpath"

	"github.com/joho/godotenv"
)

var (
    LENV_ERROR_LOAD = errors.New("error load enviroment variables")
)

func Load(isTest bool) (env map[string]string, err error) {
    if lpath.ROOT_PATH == "" {
        err = LENV_ERROR_LOAD
    } else {
        env, err = godotenv.Read(lpath.ROOT_PATH + "/.env")
    }

    if isTest {
        env["ENV"] = "test"
    }

    return
}
