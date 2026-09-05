package lerror


type MsgErrors struct {
	message string
	errors []error
}

func (m *MsgErrors) GetMessage() string {
	return m.message
}

func (m *MsgErrors) GetErrors() []error {
	return m.errors
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

func (v *ValueError) AppendErr(msg string, errors ...error) {
	v.errors = append(
		v.errors,
		MsgErrors{msg, errors})
}
