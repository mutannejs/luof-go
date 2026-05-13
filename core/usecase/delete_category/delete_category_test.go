package delete_category

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	categoryNotExists = domain.CATEGORY_NOT_EXISTS
	hasLinks = domain.HAS_LINKS
	hasSubcategories = domain.HAS_SUBCATEGORIES
	mockUidCategory = domain.MockUidCategory
)

func TestDeleteCategory_NotExists(t *testing.T) {
	var assert = assert.New(t)

	var cRepo = repository.NewCategoryMockRepository()
	var btRepo = repository.NewBelongsToMockRepository()
	var dc = New(btRepo, cRepo)

	cRepo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, nil)
	btRepo.On("HasLinks", mock.AnythingOfType("uuid.UUID")).Return(false, nil)
	cRepo.On("HasSubcategories", mock.AnythingOfType("uuid.UUID")).Return(false, nil)
	cRepo.On("Delete", mockUidCategory).Return(nil)

	err := dc.Execute(mockUidCategory)

	assert.ErrorIs(
		categoryNotExists,
		err,
		"Tentativa de deletar uma categoria que não existe deveria retornar erro contendo " + categoryNotExists.Error())
}

func TestDeleteCategory_Exists(t *testing.T) {
	var assert = assert.New(t)

	var cRepo = repository.NewCategoryMockRepository()
	var btRepo = repository.NewBelongsToMockRepository()
	var dc = New(btRepo, cRepo)

	cRepo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(true, nil)
	btRepo.On("HasLinks", mock.AnythingOfType("uuid.UUID")).Return(false, nil)
	cRepo.On("HasSubcategories", mock.AnythingOfType("uuid.UUID")).Return(false, nil)
	cRepo.On("Delete", mockUidCategory).Return(nil)

	err := dc.Execute(mockUidCategory)

	assert.NoError(
		err,
		"Deletar uma categoria válida não deveria retornar erro")
}

func TestDeleteCategory_HasLinks(t *testing.T) {
	var assert = assert.New(t)

	var cRepo = repository.NewCategoryMockRepository()
	var btRepo = repository.NewBelongsToMockRepository()
	var dc = New(btRepo, cRepo)

	cRepo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(true, nil)
	btRepo.On("HasLinks", mock.AnythingOfType("uuid.UUID")).Return(true, nil)
	cRepo.On("HasSubcategories", mock.AnythingOfType("uuid.UUID")).Return(false, nil)
	cRepo.On("Delete", mockUidCategory).Return(nil)

	err := dc.Execute(mockUidCategory)

	assert.ErrorIs(
		hasLinks,
		err,
		"Tentativa de deletar uma categoria que possui links deveria retornar erro contendo " + hasLinks.Error())
}

func TestDeleteCategory_HasSubcategories(t *testing.T) {
	var assert = assert.New(t)

	var cRepo = repository.NewCategoryMockRepository()
	var btRepo = repository.NewBelongsToMockRepository()
	var dc = New(btRepo, cRepo)

	cRepo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(true, nil)
	btRepo.On("HasLinks", mock.AnythingOfType("uuid.UUID")).Return(false, nil)
	cRepo.On("HasSubcategories", mock.AnythingOfType("uuid.UUID")).Return(true, nil)
	cRepo.On("Delete", mockUidCategory).Return(nil)

	err := dc.Execute(mockUidCategory)

	assert.ErrorIs(
		hasSubcategories,
		err,
		"Tentativa de deletar uma categoria que possui subcategorias deveria retornar erro contendo " + hasSubcategories.Error())
}
