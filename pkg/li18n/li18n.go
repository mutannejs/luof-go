package li18n

import (
	"github.com/mutannejs/luof-go/pkg/lpath"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	languages = []string{
		"pt",
	}
)

func GetBundle() (bundle *i18n.Bundle) {
	localesDir := lpath.GetAbsolutetPath("locales")

	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	for _, l := range languages {
		bundle.MustLoadMessageFile(lpath.Join(localesDir, "active." + l + ".toml"))
	}

	return
}
