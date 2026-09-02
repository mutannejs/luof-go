package insert_subcategory

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/ltests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	alternativeMockUidCategory = domain.AlternativeMockUidCategory
	childNotExists = domain.CHILD_NOT_EXISTS
	fatherNotExists = domain.FATHER_NOT_EXISTS
	isAncestor = domain.ANCESTOR_NOT_BECOME_A_SUBCATEGORY
	isSameCategory = domain.CANNOT_BE_A_SUBCATEGORY_OF_ITSELF
	isSubcategory = domain.IS_SUBCATEGORY
	mockUidCategory = domain.MockUidCategory
)

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
			"Exists",
			mock.AnythingOfType("uuid.UUID"),
		).Return(true, nil).
		On(
			"InsertSubcategory",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("time.Time"),
		).Return(nil)

	err := is.Execute(alternativeMockUidCategory, mockUidCategory)

	assert.Equal(
		ltests.GetMsgError(err),
		isAncestor,
		"Tentar inserir uma subcategoria em outra categoria, a primeira sendo ancestral da outra, deveria retornar erro contendo " + isAncestor)
}

func TestInsertSubcategory_ChildNotExists(t *testing.T) {
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
			"Exists",
			alternativeMockUidCategory,
		).Return(true, nil).
		On(
			"Exists",
			mockUidCategory,
		).Return(false, nil).
		On(
			"InsertSubcategory",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("time.Time"),
		).Return(nil)

	err := is.Execute(alternativeMockUidCategory, mockUidCategory)

	assert.Equal(
		ltests.GetMsgError(err),
		childNotExists,
		"Tentar inserir uma subcategoria que não existe em uma categoria válida deveria retornar erro contendo " + childNotExists)
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
			"Exists",
			mock.AnythingOfType("uuid.UUID"),
		).Return(true, nil).
		On(
			"InsertSubcategory",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("time.Time"),
		).Return(nil)

	err := is.Execute(alternativeMockUidCategory, mockUidCategory)

	assert.Equal(
		ltests.GetMsgError(err),
		isSubcategory,
		"Tentar inserir uma subcategoria em outra categoria, ambas já relacionadas, deveria retornar erro contendo " + isSubcategory)
}

func TestInsertSubcategory_FatherNotExists(t *testing.T) {
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
			"Exists",
			alternativeMockUidCategory,
		).Return(false, nil).
		On(
			"Exists",
			mockUidCategory,
		).Return(true, nil).
		On(
			"InsertSubcategory",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("time.Time"),
		).Return(nil)

	err := is.Execute(alternativeMockUidCategory, mockUidCategory)

	assert.Equal(
		ltests.GetMsgError(err),
		fatherNotExists,
		"Tentar inserir uma subcategoria válida em uma categoria que não existe deveria retornar erro contendo " + fatherNotExists)
}

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
			"Exists",
			mock.AnythingOfType("uuid.UUID"),
		).Return(true, nil).
		On(
			"InsertSubcategory",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("time.Time"),
		).Return(nil)

	err := is.Execute(alternativeMockUidCategory, mockUidCategory)

	assert.True(
		err.IsNil(),
		"Tentar inserir uma subcategoria em outra categoria, ambas não relacionadas, não deveria retornar erro")
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
			"Exists",
			mock.AnythingOfType("uuid.UUID"),
		).Return(true, nil).
		On(
			"InsertSubcategory",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("time.Time"),
		).Return(nil)

	err := is.Execute(mockUidCategory, mockUidCategory)

	assert.Equal(
		ltests.GetMsgError(err),
		isSameCategory,
		"Tentar inserir uma subcategoria nela mesma deveria retornar erro contendo " + isSameCategory)
}
