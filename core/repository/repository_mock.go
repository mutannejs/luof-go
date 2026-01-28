package repository

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/ltests"
)

func NewCategoryMockRepository() *ltests.MockCrudRepository[domain.Category] {
	return &ltests.MockCrudRepository[domain.Category]{}
}

func NewLinkMockRepository() *ltests.MockCrudRepository[domain.Link] {
	return &ltests.MockCrudRepository[domain.Link]{}
}

func NewBelongsToMockRepository() *ltests.BelongsToMockRepository[domain.Link] {
	return &ltests.BelongsToMockRepository[domain.Link]{}
}

func NewSubcategoryMockRepository() *ltests.SubcategoryMockRepository[domain.Category] {
	return &ltests.SubcategoryMockRepository[domain.Category]{}
}
