package lenv

import (
	"flag"

	"github.com/mutannejs/luof-go/pkg/li18n"
	"github.com/mutannejs/luof-go/pkg/lpath"

	"github.com/joho/godotenv"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	LENV_ERROR_LOAD = li18n.Msg2Error(&i18n.Message{
		ID: "LENV_ERROR_LOAD",
		Description: "Erro retornado quando uuid.New falha",
		Other: "error load enviroment variables",
	})
)

func Load() (env map[string]string, err error) {
	env, err = loadEnv()

	var environment string
	var isTest bool

	flag.StringVar(&environment, "env", "", "execution environment")
	flag.BoolVar(&isTest, "test", false, "is test environment")
	flag.Parse()

	if isTest || environment == "test" {
		env["ENV"] = "test"
	} else if environment != "" {
		env["ENV"] = environment
	}

	return
}

func LoadTest() (env map[string]string, err error) {
	env, err = loadEnv()

	if err == nil {
		env["ENV"] = "test"
	}

	return
}

func loadEnv() (env map[string]string, err error) {
	if lpath.ROOT_PATH == "" {
		err = LENV_ERROR_LOAD
	} else {
		env, err = godotenv.Read(lpath.ROOT_PATH + "/.env")
	}

	return
}
