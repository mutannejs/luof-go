package repository

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/ltests"
)

func NewCategoryMockRepository() *ltests.CategoryMockRepository[domain.Category] {
	return &ltests.CategoryMockRepository[domain.Category]{}
}

func NewLinkMockRepository() *ltests.LinkMockRepository[domain.Link] {
	return &ltests.LinkMockRepository[domain.Link]{}
}

func NewBelongsToMockRepository() *ltests.BelongsToMockRepository[domain.Link] {
	return &ltests.BelongsToMockRepository[domain.Link]{}
}
