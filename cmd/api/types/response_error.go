package types

type ParamsErrors map[string]string

type ResponseError struct {
	Message string
	Errors ParamsErrors
}
