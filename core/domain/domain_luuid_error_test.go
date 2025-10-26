//go:build luuid_error

package domain

import (
    "testing"

    "github.com/mutannejs/luof-go/pkg/luuid"

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
    assert := assert.New(t)
    category, err := NewCategory(
        categoryMock["name"].(string),
        categoryMock["description"].(string),
        categoryMock["useMarkdown"].(bool),
    )

    assert.ErrorIs(err, CATEGORY_ERROR_NEW, "erro na criação de categoria deveria conter o erro " + CATEGORY_ERROR_NEW.Error())
    assert.ErrorIs(err, luuid.UUID_ERROR_NEW, "erro na criação de categoria deveria conter o erro " + luuid.UUID_ERROR_NEW.Error())
    assert.Zero(category, "a categoria retornado deveria ser zero")
}

func TestNewLink(t *testing.T) {
    assert := assert.New(t)
    link, err := NewLink(
        linkMock["url"].(string),
        linkMock["name"].(string),
        linkMock["description"].(string),
        linkMock["useMarkdown"].(bool),
    )

    assert.ErrorIs(err, LINK_ERROR_NEW, "erro na criação de link deveria conter o erro " + LINK_ERROR_NEW.Error())
    assert.ErrorIs(err, luuid.UUID_ERROR_NEW, "erro na criação de link deveria conter o erro " + luuid.UUID_ERROR_NEW.Error())
    assert.Zero(link, "o link retornado deveria ser zero")
}
