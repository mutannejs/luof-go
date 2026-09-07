package lerror

type MsgErrors struct {
	Message string `json:"message"`
	Errors []string `json:"errors"`
}

func (m *MsgErrors) GetMessage() string {
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

func (v *ValueError) AppendErr(msg string, errors ...string) {
	v.errors = append(
		v.errors,
		MsgErrors{msg, errors})
}
