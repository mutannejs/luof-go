package insert_subcategory

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	areRelated = domain.ARE_RELATED
	mockUidCategory = domain.MockUidCategory
	alternativeMockUidCategory = domain.AlternativeMockUidCategory
)

func TestInsertSubcategory_NotExists(t *testing.T) {
	var assert = assert.New(t)

	var cRepo = repository.NewCategoryMockRepository()
	var is = New(cRepo)

	cRepo.
		On(
			"AreRelated",
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
			"AreRelated",
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
		areRelated,
		"Tentar inserir uma subcategoria em outra categoria, ambas já relacionadas, deveria retornar erro contendo " + areRelated.Error())
}
