package get_all_root_categories

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/stretchr/testify/assert"
)

var (
	mockCategory = domain.MockCategory
)

func TestGetAllRootCategories_Empty(t *testing.T) {
	var assert = assert.New(t)

	var repo = repository.NewCategoryMockRepository()
	var alrc = New(repo)

	repo.On("GetAllRootCategories").Return(make([]domain.Category, 0), nil)

	categories, err := alrc.Execute()

	assert.True(
		err.IsNil(),
		"Buscar todas as categorias raízes não deveria retornar erro")
	assert.Len(
		categories,
		0)
}

func TestGetAllRootCategories_NoEmpty(t *testing.T) {
	var assert = assert.New(t)

	var repo = repository.NewCategoryMockRepository()
	var alrc = New(repo)

	var mockCategories = []domain.Category{mockCategory}
	repo.On("GetAllRootCategories").Return(mockCategories, nil)

	categories, err := alrc.Execute()

	assert.True(
		err.IsNil(),
		"Buscar todas as categorias raízes não deveria retornar erro")
	assert.Len(
		categories,
		1)
}

func TestGetAllRootCategories_NotEmpty(t *testing.T) {}
