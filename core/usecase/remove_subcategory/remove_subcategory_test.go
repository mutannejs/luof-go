package remove_subcategory

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/ltests"

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

	var cRepo = repository.NewCategoryMockRepository()
	var is = New(cRepo)

	cRepo.
		On(
			"Exists",
			mock.AnythingOfType("uuid.UUID"),
		).Return(true, nil).
		On(
			"IsSubcategory",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(true, nil).
		On(
			"DeleteSubcategory",
			mock.AnythingOfType("uuid.UUID"),
		).Return(nil)

	err := is.Execute(alternativeMockUidCategory, mockUidCategory)

	assert.True(
		err.IsNil(),
		"Tentar remover uma subcategoria válida não deveria retornar erro")
}

func TestRemoveSubcategory_NotExists(t *testing.T) {
	var assert = assert.New(t)

	var cRepo = repository.NewCategoryMockRepository()
	var is = New(cRepo)

	cRepo.
		On(
			"Exists",
			mock.AnythingOfType("uuid.UUID"),
		).Return(true, nil).
		On(
			"IsSubcategory",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(false, nil).
		On(
			"DeleteSubcategory",
			mock.AnythingOfType("uuid.UUID"),
		).Return(nil)

	err := is.Execute(alternativeMockUidCategory, mockUidCategory)

	assert.Equal(
		ltests.GetMsgError(err),
		notIsSubcategory,
		"Tentar remover uma categoria de outra, sem que ela seja uma subcategoria direta desta, deveria retornar erro contendo " + notIsSubcategory)
}
