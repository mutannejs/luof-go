package usecase

import (
    "testing"

    "github.com/mutannejs/luof-go/core/usecase"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

var (
    categoryNotExists = usecase.CategoryNotExists
    mockLinks = usecase.MockLinks
    mockUidCategory = usecase.MockUidCategory
)

func TestGetLinksByCategory_NotExists(t *testing.T) {
    var assert = assert.New(t)

    var btRepo = usecase.NewBelongsToMockRepository()
    var cRepo = usecase.NewCategoryMockRepository()
    var glbc = NewGetLinksByCategory(btRepo, cRepo)

    cRepo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, categoryNotExists)

    links, err := glbc.Execute(mockUidCategory)

    assert.Zero(
                    links,
                    "Deveria ser retornado zero para um uid que não pertence a nenhuma categoria existente")
    assert.ErrorIs(
                    err,
                    categoryNotExists,
                    "Buscar uma categoria que não existe deveria retornar erro contendo " + categoryNotExists.Error())
}

func TestGetLinksByCategory_Exists(t *testing.T) {
    var assert = assert.New(t)

    var btRepo = usecase.NewBelongsToMockRepository()
    var cRepo = usecase.NewCategoryMockRepository()
    var glbc = NewGetLinksByCategory(btRepo, cRepo)

    cRepo.On("Exists", mockUidCategory).Return(true, nil)
    btRepo.On("GetLinksByCategory", mockUidCategory).Return(mockLinks, nil)

    links, err := glbc.Execute(mockUidCategory)

    assert.Contains(
                    links,
                    mockLinks[0],
                    "Todos os links armazenados no repositório devem ser retornados pela função")
    assert.Contains(
                    links,
                    mockLinks[1],
                    "Todos os links armazenados no repositório devem ser retornados pela função")
    assert.NoError(
                    err,
                    "Buscar por uma categoria válida não deveria retornar erro")
}
