package ltests

func GetMsgError(err error) error {
	return err.(interface{ Unwrap() []error }).Unwrap()[1]
}
