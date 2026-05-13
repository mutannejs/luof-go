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

func TestGetSubcategories_CategoryNotExists(t *testing.T) {
	var assert = assert.New(t)

	var cRepo = repository.NewCategoryMockRepository()
	var glbc = New(cRepo)

	cRepo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, nil)

	subcategories, err := glbc.Execute(mockUidCategory)

	assert.Zero(
		subcategories,
		"Deveria ser retornado zero para um uid que não pertence a nenhuma categoria existente")
	assert.ErrorIs(
		categoryNotExists,
		err,
		"Buscar uma categoria que não existe deveria retornar erro contendo " + categoryNotExists.Error())
}

func TestGetSubcategories_CategoryExists(t *testing.T) {
	var assert = assert.New(t)

	var cRepo = repository.NewCategoryMockRepository()
	var glbc = New(cRepo)

	cRepo.On("Exists", mockUidCategory).Return(true, nil)
	cRepo.On("GetSubcategories", mockUidCategory).Return(mockCategories, nil)

	subcategories, err := glbc.Execute(mockUidCategory)

	assert.Contains(
		subcategories,
		mockCategories[0],
		"Todas as subcategorias no repositório devem ser retornadas pela função")
	assert.Contains(
		subcategories,
		mockCategories[1],
		"Todas as subcategorias no repositório devem ser retornadas pela função")
	assert.NoError(
		err,
		"Buscar por uma categoria válida não deveria retornar erro")
}

func TestGetSubcategories_EmptyCategory(t *testing.T) {
	var assert = assert.New(t)

	var cRepo = repository.NewCategoryMockRepository()
	var glbc = New(cRepo)

	var emptySubcategories []domain.Category

	cRepo.On("Exists", mockUidCategory).Return(true, nil)
	cRepo.On("GetSubcategories", mockUidCategory).Return(emptySubcategories, nil)

	subcategories, err := glbc.Execute(mockUidCategory)

	assert.Len(
		subcategories,
		0,
		"Buscar as subcategorias de uma categoria vazia deveria retornar um array vazio")
	assert.NoError(
		err,
		"Buscar por uma categoria válida não deveria retornar erro")
}
