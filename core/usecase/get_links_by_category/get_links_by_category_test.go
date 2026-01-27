package get_links_by_category

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	categoryNotExists = domain.CATEGORY_NOT_EXISTS
	mockLinks = domain.MockLinks
	mockUidCategory = domain.MockUidCategory
)

func TestGetLinksByCategory_CategoryNotExists(t *testing.T) {
	var assert = assert.New(t)

	var btRepo = repository.NewBelongsToMockRepository()
	var cRepo = repository.NewCategoryMockRepository()
	var glbc = New(btRepo, cRepo)

	cRepo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, nil)

	links, err := glbc.Execute(mockUidCategory)

	assert.Zero(
					links,
					"Deveria ser retornado zero para um uid que não pertence a nenhuma categoria existente")
	assert.ErrorIs(
					err,
					categoryNotExists,
					"Buscar uma categoria que não existe deveria retornar erro contendo " + categoryNotExists.Error())
}

func TestGetLinksByCategory_CategoryExists(t *testing.T) {
	var assert = assert.New(t)

	var btRepo = repository.NewBelongsToMockRepository()
	var cRepo = repository.NewCategoryMockRepository()
	var glbc = New(btRepo, cRepo)

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
