package domain

import (
	"errors"
)

var (
	ALREADY_BELONGS = errors.New("the link already belongs to the category")
	HAS_LINKS = errors.New("the category cannot be deleted because it has links")
	NOT_BELONGS = errors.New("the link does not belong to the category")
)
