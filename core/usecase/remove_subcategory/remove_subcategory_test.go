package insert_subcategory

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	notIsSubcategory = domain.NOT_IS_SUBCATEGORY
	mockUidCategory = domain.MockUidCategory
	alternativeMockUidCategory = domain.AlternativeMockUidCategory
)

func TestRemoveSubcategory_Exists(t *testing.T) {
	var assert = assert.New(t)

	var sRepo = repository.NewSubcategoryMockRepository()
	var is = New(sRepo)

	sRepo.
		On(
			"IsSubcategory",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(true, nil).
		On(
			"Delete",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(nil)

	err := is.Execute(alternativeMockUidCategory, mockUidCategory)

	assert.NoError(
		err,
		"Tentar remover uma subcategoria válida não deveria retornar erro")
}

func TestRemoveSubcategory_NotExists(t *testing.T) {
	var assert = assert.New(t)

	var sRepo = repository.NewSubcategoryMockRepository()
	var is = New(sRepo)

	sRepo.
		On(
			"IsSubcategory",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(false, nil).
		On(
			"Delete",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(nil)

	err := is.Execute(alternativeMockUidCategory, mockUidCategory)

	assert.ErrorIs(
		err,
		notIsSubcategory,
		"Tentar remover uma categoria de outra, sem que ela seja uma subcategoria direta desta, deveria retornar erro contendo " + notIsSubcategory.Error())
}
