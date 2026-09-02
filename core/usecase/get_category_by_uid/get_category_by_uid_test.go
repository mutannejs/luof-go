package get_category_by_uid

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/ltests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	categoryNotExists = domain.CATEGORY_NOT_EXISTS
	mockCategory = domain.MockCategory
	mockUidCategory = domain.MockUidCategory
)

func TestGetCategoryByUid_NotExists(t *testing.T) {
	var assert = assert.New(t)

	var repo = repository.NewCategoryMockRepository()
	var gcbu = New(repo)

	repo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, nil)

	category, err := gcbu.Execute(mockUidCategory)

	assert.Zero(
		category,
		"Deveria ser retornado zero para um uid inválido")
	assert.Equal(
		ltests.GetMsgError(err),
		categoryNotExists,
		"Buscar uma categoria que não existe deveria retornar erro contendo " + categoryNotExists)
}

func TestGetCategoryByUid_Exists(t *testing.T) {
	var assert = assert.New(t)

	var repo = repository.NewCategoryMockRepository()
	var gcbu = New(repo)

	repo.On("Exists", mockUidCategory).Return(true, nil)
	repo.On("GetByUid", mockUidCategory).Return(mockCategory, nil)

	category, err := gcbu.Execute(mockUidCategory)

	assert.Equal(
		mockCategory,
		category,
		"A categoria retornada pela função deve ser a mesma retornada pelo repositório")
	assert.True(
		err.IsNil(),
		"Buscar uma categoria válida não deveria retornar erro")
}
