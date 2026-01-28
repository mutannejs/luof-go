package get_subcategories

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	categoryNotExists = domain.CATEGORY_NOT_EXISTS
	mockUidCategory = domain.MockUidCategory
	mockCategories = domain.MockCategories
)

func TestGetLinksByCategory_CategoryNotExists(t *testing.T) {
	var assert = assert.New(t)

	var sRepo = repository.NewSubcategoryMockRepository()
	var cRepo = repository.NewCategoryMockRepository()
	var glbc = New(cRepo, sRepo)

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

	var sRepo = repository.NewSubcategoryMockRepository()
	var cRepo = repository.NewCategoryMockRepository()
	var glbc = New(cRepo, sRepo)

	cRepo.On("Exists", mockUidCategory).Return(true, nil)
	sRepo.On("GetSubcategories", mockUidCategory).Return(mockCategories, nil)

	links, err := glbc.Execute(mockUidCategory)

	assert.Contains(
					links,
					mockCategories[0],
					"Todas as subcategorias no repositório devem ser retornadas pela função")
	assert.Contains(
					links,
					mockCategories[1],
					"Todas as subcategorias no repositório devem ser retornadas pela função")
	assert.NoError(
					err,
					"Buscar por uma categoria válida não deveria retornar erro")
}
