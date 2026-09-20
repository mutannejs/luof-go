package lerror

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type MsgErrors struct {
	Message *i18n.Message `json:"message"`
	Errors []string `json:"errors"`
}

func (m *MsgErrors) GetMessage() *i18n.Message {
	return m.Message
}

func (m *MsgErrors) GetErrors() []string {
	return m.Errors
}

type ValueError struct {
	code int
	errors []MsgErrors
}

func (v *ValueError) IsNil() bool {
	return len(v.errors) == 0
}

func (v *ValueError) GetCode() int {
	return v.code
}

func (v *ValueError) GetErrors() []MsgErrors {
	return v.errors
}

func (v *ValueError) AppendErr(msg *i18n.Message, errors ...string) {
	var errs []string

	if errors != nil {
		errs = errors
	} else {
		errs = make([]string, 0)
	}

	var mE = MsgErrors{msg, errs}

	if v.errors == nil {
		v.errors = []MsgErrors{mE}
	} else {
		v.errors = append(v.errors, mE)
	}
}
