package domain

import (
	"errors"
)

var (
	ALREADY_BELONGS = errors.New("the link already belongs to the category")
	NOT_BELONGS = errors.New("the link does not belong to the category")
)
