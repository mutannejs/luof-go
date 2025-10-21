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

func TestNewCategory(t *testing.T) {
    category, err := NewCategory(
        categoryMock["name"].(string),
        categoryMock["description"].(string),
        categoryMock["useMarkdown"].(bool),
    )

    assert.NoError(t, err, "criação com dados válidos não deveria falhar")
    assert.NotZero(t, category.GetUid(), "uid deveria ser um uuid diferente de zero")
    assert.Equal(t, category.Name, categoryMock["name"], "o nome deveria ser igual ao argumento do construtor")
    assert.Equal(t, category.Description.Content, categoryMock["description"], "a descrição deveria ser igual ao argumento do construtor")
    assert.Equal(t, category.Description.UseMarkdown, categoryMock["useMarkdown"], "o valor useMarkdown deveria ser igual ao argumento do construtor")
    assert.NotZero(t, category.CreatedAt, "createdAt deveria ser diferente de zero")
    assert.Zero(t, category.UpdatedAt, "updatedAt deveria ser diferente de zero")
}

func TestNewLink(t *testing.T) {
    link, err := NewLink(
        linkMock["url"].(string),
        linkMock["name"].(string),
        linkMock["description"].(string),
        linkMock["useMarkdown"].(bool),
    )

    assert.NoError(t, err, "criação com dados válidos não deveria falhar")
    assert.NotZero(t, link.GetUid(), "uid deveria ser um uuid diferente de zero")
    assert.Equal(t, link.Url, linkMock["url"], "a url deveria ser igual ao argumento do construtor")
    assert.Equal(t, link.Name, linkMock["name"], "o nome deveria ser igual ao argumento do construtor")
    assert.Equal(t, link.Description.Content, linkMock["description"], "a descrição deveria ser igual ao argumento do construtor")
    assert.Equal(t, link.Description.UseMarkdown, linkMock["useMarkdown"], "o valor useMarkdown deveria ser igual ao argumento do construtor")
    assert.NotZero(t, link.CreatedAt, "createdAt deveria ser diferente de zero")
    assert.Zero(t, link.UpdatedAt, "updatedAt deveria ser diferente de zero")
}
