package create_category

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCategory(t *testing.T) {
	var assert = assert.New(t)
	var mockCategory = domain.MockCategory
	var category domain.Category

	var repo = repository.NewCategoryMockRepository()
	var cc = New(repo)

	repo.On("Create", mock.MatchedBy(func(c domain.Category) bool {
		category = c
		return true
	})).Return(nil)

	uid, err := cc.Execute(
		mockCategory.Name,
		mockCategory.Description.Content,
		mockCategory.Description.UseMarkdown,
	)

	// Testa se a categoria enviada para Repository.Create está de acordo
	// com os argumentos passados à função
	// ! Pode estar errada, mesmo que NewCategory tenha sido corretamente
	// implementada
	assert.Equal(mockCategory.Name, category.Name)
	assert.Equal(mockCategory.Description.Content, category.Description.Content)
	assert.Equal(mockCategory.Description.UseMarkdown, category.Description.UseMarkdown)
	assert.NotZero(category.CreatedAt)
	assert.Zero(category.UpdatedAt)

	// Testa o retorno da função
	assert.NotZero(
		uid,
		"Criação com dados válidos deveria retornar um uuid diferente de zero")
	assert.True(
		err.IsNil(),
		"Criação com dados válidos não deveria retornar erro")
}
