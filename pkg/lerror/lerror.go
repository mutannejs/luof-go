package lerror

import (
	"fmt"
)

func Internal(err *error) {
	if *err != nil {
		*err = fmt.Errorf("500%w", *err)
	}
}

func BadRequest(err *error) {
	if *err != nil {
		*err = fmt.Errorf("400%w", *err)
	}
}
