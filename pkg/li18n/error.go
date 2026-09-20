package li18n

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type I18nError struct {
	msg *i18n.Message
}

func (i *I18nError) Error() string {
	return i.msg.Description
}

func Msg2Error(msg *i18n.Message) error {
	return &I18nError{msg}
}
