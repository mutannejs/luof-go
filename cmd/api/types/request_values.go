package types

import (
    z "github.com/Oudwins/zog"
)

type RequestValues struct {
	JsonBody any
	Params any
}

type RequestValidations struct {
	JsonBody *z.StructSchema
	Params *z.StructSchema
}
