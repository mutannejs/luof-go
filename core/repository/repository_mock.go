package repository

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/ltests"
)

func NewCategoryMockRepository() *ltests.MockCategoryRepository[domain.Category] {
	return &ltests.MockCategoryRepository[domain.Category]{}
}

func NewLinkMockRepository() *ltests.MockCrudRepository[domain.Link] {
	return &ltests.MockCrudRepository[domain.Link]{}
}

func NewBelongsToMockRepository() *ltests.BelongsToMockRepository[domain.Link] {
	return &ltests.BelongsToMockRepository[domain.Link]{}
}
