package usecase

import (
    "errors"

    "github.com/mutannejs/luof-go/core/domain"
    "github.com/mutannejs/luof-go/core/repository"
    "github.com/mutannejs/luof-go/pkg/ltests"
)

var (
    // Erros mockados
    CategoryNotExists = errors.New(repository.CATEGORY_NOT_EXISTS)
    LinkNotExists = errors.New(repository.LINK_NOT_EXISTS)

    // Elementos do domínio
    AlternativeMockLink, _ = domain.NewLink(
        "github.com/mutannejs/luof-go",
        "luof-go",
        "back-end luof repository",
        false,
    )
    MockCategory, _ = domain.NewCategory(
        "development",
        "links about development",
        false,
    )
    MockLink, _ = domain.NewLink(
        "github.com/mutannejs/luof",
        "luof",
        "luof repository",
        false,
    )

    // Conjuntos de elementos do domínio
    MockLinks = []domain.Link{MockLink, AlternativeMockLink}

    // Identificadores dos elementos mockados
    MockUidCategory = MockCategory.GetUid()
    MockUidLink = MockLink.GetUid()
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
