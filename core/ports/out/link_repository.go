package out

import (
    "github.com/mutannejs/luof-go/core/domain"
)

type LinkRepository interface {
    Create (domain.Link) error
}
