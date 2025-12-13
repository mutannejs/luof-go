package repository

import (
	"errors"

	"github.com/mutannejs/luof-go/core/domain"

	"github.com/mutannejs/luof-go/pkg/ltests"
)

var (
    // Erros
    CategoryNotExists = errors.New(CATEGORY_NOT_EXISTS)
    LinkNotExists = errors.New(LINK_NOT_EXISTS)
)

// Repositórios
func NewCategoryMockRepository() *ltests.MockCrudRepository[domain.Category] {
    return &ltests.MockCrudRepository[domain.Category]{}
}

func NewLinkMockRepository() *ltests.MockCrudRepository[domain.Link] {
    return &ltests.MockCrudRepository[domain.Link]{}
}

func NewBelongsToMockRepository() *ltests.BelongsToMockRepository[domain.Link] {
    return &ltests.BelongsToMockRepository[domain.Link]{}
}
