package insert_subcategory

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	alternativeMockUidCategory = domain.AlternativeMockUidCategory
	isAncestor = domain.ANCESTOR_NOT_BECOME_A_SUBCATEGORY
	isSameCategory = domain.CANNOT_BE_A_SUBCATEGORY_OF_ITSELF
	isSubcategory = domain.IS_SUBCATEGORY
	mockUidCategory = domain.MockUidCategory
)

func TestInsertSubcategory_NotExists(t *testing.T) {
	var assert = assert.New(t)

	var cRepo = repository.NewCategoryMockRepository()
	var is = New(cRepo)

	cRepo.
		On(
			"IsSubcategory",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(false, nil).
		On(
			"IsAncestor",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(false, nil).
		On(
			"InsertSubcategory",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("time.Time"),
		).Return(nil)

	err := is.Execute(alternativeMockUidCategory, mockUidCategory)

	assert.NoError(
		err,
		"Tentar inserir uma subcategoria em outra categoria, ambas não relacionadas, não deveria retornar erro")
}

func TestInsertSubcategory_Exists(t *testing.T) {
	var assert = assert.New(t)

	var cRepo = repository.NewCategoryMockRepository()
	var is = New(cRepo)

	cRepo.
		On(
			"IsSubcategory",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(true, nil).
		On(
			"IsAncestor",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(true, nil).
		On(
			"InsertSubcategory",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("time.Time"),
		).Return(nil)

	err := is.Execute(alternativeMockUidCategory, mockUidCategory)

	assert.ErrorIs(
		err,
		isSubcategory,
		"Tentar inserir uma subcategoria em outra categoria, ambas já relacionadas, deveria retornar erro contendo " + isSubcategory.Error())
}

func TestInsertSubcategory_AncestorBecomeASubcategory(t *testing.T) {
	var assert = assert.New(t)

	var cRepo = repository.NewCategoryMockRepository()
	var is = New(cRepo)

	cRepo.
		On(
			"IsSubcategory",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(false, nil).
		On(
			"IsAncestor",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(true, nil).
		On(
			"InsertSubcategory",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("time.Time"),
		).Return(nil)

	err := is.Execute(alternativeMockUidCategory, mockUidCategory)

	assert.ErrorIs(
		err,
		isAncestor,
		"Tentar inserir uma subcategoria em outra categoria, a primeira sendo ancestral da outra, deveria retornar erro contendo " + isAncestor.Error())
}

func TestInsertSubcategory_SubcategoryOfItself(t *testing.T) {
	var assert = assert.New(t)

	var cRepo = repository.NewCategoryMockRepository()
	var is = New(cRepo)

	cRepo.
		On(
			"IsSubcategory",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(false, nil).
		On(
			"IsAncestor",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(false, nil).
		On(
			"InsertSubcategory",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("time.Time"),
		).Return(nil)

	err := is.Execute(mockUidCategory, mockUidCategory)

	assert.ErrorIs(
		err,
		isSameCategory,
		"Tentar inserir uma subcategoria nela mesma deveria retornar erro contendo " + isSameCategory.Error())
}
