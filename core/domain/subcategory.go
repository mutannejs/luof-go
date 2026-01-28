package domain

import (
	"errors"
)

var (
	ARE_RELATIVES = errors.New("both categories are already related")
	NOT_IS_SUBCATEGORY = errors.New("the child is not subcategory of the father")
)
