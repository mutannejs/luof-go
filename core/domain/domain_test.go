//go:build !luuid_error

package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Testa a instanciação de uma nova Categoria para casos onde não ocorrem erros
func TestNewCategory(t *testing.T) {
    assert := assert.New(t)

    category, err := NewCategory(
        CategoryMockMap["name"].(string),
        CategoryMockMap["description"].(string),
        CategoryMockMap["useMarkdown"].(bool),
    )

    assert.NoError(
                    err,
                    "Criação com dados válidos não deveria falhar")
    assert.NotZero(
                    category.GetUid(),
                    "A nova categoria deveria possuir como Uid um uuid diferente de zero")
    assert.Equal(
                    category.Name,
                    CategoryMockMap["name"],
                    "O nome deveria ser igual ao argumento do construtor")
    assert.Equal(
                    category.Description.Content,
                    CategoryMockMap["description"],
                    "A descrição deveria ser igual ao argumento do construtor")
    assert.Equal(
                    category.Description.UseMarkdown,
                    CategoryMockMap["useMarkdown"],
                    "O valor UseMarkdown deveria ser igual ao argumento do construtor")
    assert.NotZero(
                    category.CreatedAt,
                    "O valor de CreatedAt deveria ser diferente de zero")
    assert.Zero(
                    category.UpdatedAt,
                    "O valor de UpdatedAt deveria ser zero")
}

// Testa a instanciação de um novo Link para casos onde não ocorrem erros
func TestNewLink(t *testing.T) {
    assert := assert.New(t)

    link, err := NewLink(
        LinkMockMap["url"].(string),
        LinkMockMap["name"].(string),
        LinkMockMap["description"].(string),
        LinkMockMap["useMarkdown"].(bool),
    )

    assert.NoError(
                    err,
                    "criação com dados válidos não deveria falhar")
    assert.NotZero(
                    link.GetUid(),
                    "O novo link deveria possui como Uid um uuid diferente de zero")
    assert.Equal(
                    link.Url,
                    LinkMockMap["url"],
                    "A url deveria ser igual ao argumento do construtor")
    assert.Equal(
                    link.Name,
                    LinkMockMap["name"],
                    "O nome deveria ser igual ao argumento do construtor")
    assert.Equal(
                    link.Description.Content,
                    LinkMockMap["description"],
                    "A descrição deveria ser igual ao argumento do construtor")
    assert.Equal(
                    link.Description.UseMarkdown,
                    LinkMockMap["useMarkdown"],
                    "O valor UseMarkdown deveria ser igual ao argumento do construtor")
    assert.NotZero(
                    link.CreatedAt,
                    "O valor de CreatedAt deveria ser diferente de zero")
    assert.Zero(
                    link.UpdatedAt,
                    "O valor de UpdatedAt deveria ser zero")
}
