package custom_log

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetConfig(env map[string]string) {
	if env["ENV"] != "prod" {
		cW := zerolog.ConsoleWriter{
			Out: os.Stderr,
			TimeFormat: time.TimeOnly,
		};

		cW.FormatLevel = func(i any) string {
			return strings.ToUpper(fmt.Sprintf("%-5s", i))
		}

		log.Logger = zerolog.New(cW).With().Timestamp().Logger()
	} else {
		// DateTime with nanoseconds
		zerolog.TimeFieldFormat = time.DateTime + ".000000000"
	}
}
