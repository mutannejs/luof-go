//go:build !luuid_error

package domain

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

var (
    linkMock = map[string]any{
        "url": "github.com/mutannejs/luof-go",
        "name": "luof",
        "description": "### luof-go repository",
        "useMarkdown": true,
    }
    categoryMock = map[string]any{
        "name": "development",
        "description": "links about development",
        "useMarkdown": false,
    }
)

// Testa a instanciação de uma nova Categoria para casos onde não ocorrem erros
func TestNewCategory(t *testing.T) {
    assert := assert.New(t)

    category, err := NewCategory(
        categoryMock["name"].(string),
        categoryMock["description"].(string),
        categoryMock["useMarkdown"].(bool),
    )

    assert.NoError(
                    err,
                    "Criação com dados válidos não deveria falhar")
    assert.NotZero(
                    category.GetUid(),
                    "A nova categoria deveria possuir como Uid um uuid diferente de zero")
    assert.Equal(
                    category.Name,
                    categoryMock["name"],
                    "O nome deveria ser igual ao argumento do construtor")
    assert.Equal(
                    category.Description.Content,
                    categoryMock["description"],
                    "A descrição deveria ser igual ao argumento do construtor")
    assert.Equal(
                    category.Description.UseMarkdown,
                    categoryMock["useMarkdown"],
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
        linkMock["url"].(string),
        linkMock["name"].(string),
        linkMock["description"].(string),
        linkMock["useMarkdown"].(bool),
    )

    assert.NoError(
                    err,
                    "criação com dados válidos não deveria falhar")
    assert.NotZero(
                    link.GetUid(),
                    "O novo link deveria possui como Uid um uuid diferente de zero")
    assert.Equal(
                    link.Url,
                    linkMock["url"],
                    "A url deveria ser igual ao argumento do construtor")
    assert.Equal(
                    link.Name,
                    linkMock["name"],
                    "O nome deveria ser igual ao argumento do construtor")
    assert.Equal(
                    link.Description.Content,
                    linkMock["description"],
                    "A descrição deveria ser igual ao argumento do construtor")
    assert.Equal(
                    link.Description.UseMarkdown,
                    linkMock["useMarkdown"],
                    "O valor UseMarkdown deveria ser igual ao argumento do construtor")
    assert.NotZero(
                    link.CreatedAt,
                    "O valor de CreatedAt deveria ser diferente de zero")
    assert.Zero(
                    link.UpdatedAt,
                    "O valor de UpdatedAt deveria ser zero")
}
