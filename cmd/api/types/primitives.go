package types

import (
	z "github.com/Oudwins/zog"
)

var (
	UidValidate = z.String().UUID().Required()
)
