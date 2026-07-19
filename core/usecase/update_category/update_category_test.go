package update_category

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

func TestUpdateCategory_NotExists(t *testing.T) {
	var assert = assert.New(t)

	var repo = repository.NewCategoryMockRepository()
	var uc = New(repo)

	repo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, nil)

	exists, err := uc.Execute(
		mockUidCategory,
		mockCategory.Name,
		mockCategory.Description.Content,
		mockCategory.Description.UseMarkdown,
	)

	assert.False(
		exists,
		"Não deveria ser possível atualizar uma categoria que não existe")
	assert.ErrorIs(
		ltests.GetMsgError(err),
		categoryNotExists,
		"Tentar atualizar uma categoria que não existe deveria retornar erro contendo " + categoryNotExists.Error())
}

func TestUpdateCategory_Exists(t *testing.T) {
	var assert = assert.New(t)
	var category domain.Category

	var repo = repository.NewCategoryMockRepository()
	var uc = New(repo)

	repo.On("Exists", mockUidCategory).Return(true, nil)
	repo.On("Update", mockUidCategory, mock.MatchedBy(func(c domain.Category) bool {
		category = c
		return true
	})).Return(nil)

	exists, err := uc.Execute(
		mockUidCategory,
		mockCategory.Name,
		mockCategory.Description.Content,
		mockCategory.Description.UseMarkdown,
	)

	// Testa se a categoria enviada para Repository.Update está de acordo
	// com os argumentos passados à função
	// ! Pode estar errada, mesmo que NewCategory tenha sido corretamente
	// implementada
	assert.Equal(category.Name, mockCategory.Name)
	assert.Equal(category.Description.Content, mockCategory.Description.Content)
	assert.Equal(category.Description.UseMarkdown, mockCategory.Description.UseMarkdown)
	assert.NotZero(category.CreatedAt)

	// Testa o valor de UpdatedAt
	assert.Less(
		category.CreatedAt,
		category.UpdatedAt,
		"O valor de UpdatedAt deve ser maior que o valor de CreatedAt")

	// Testa o retorno da função
	assert.True(
		exists,
		"Deveria ser possível atualizar uma categoria válida")
	assert.NoError(
		err,
		"Atualizar uma categoria válida não deveria retornar erro")
}
