//go:build luuid_error

package domain

import (
	"testing"

	"github.com/mutannejs/luof-go/pkg/luuid"

	"github.com/stretchr/testify/assert"
)

// Testa a instanciação de uma nova Categoria para casos onde ocorrem erros de UUID
func TestNewCategory(t *testing.T) {
    assert := assert.New(t)

    category, err := NewCategory(
        CategoryMockMap["name"].(string),
        CategoryMockMap["description"].(string),
        CategoryMockMap["useMarkdown"].(bool),
    )

    assert.ErrorIs(
                    err,
                    CATEGORY_ERROR_NEW,
                    "Erro na criação de categoria deveria conter o erro " + CATEGORY_ERROR_NEW.Error())
    assert.ErrorIs(
                    err,
                    luuid.UUID_ERROR_NEW,
                    "Erro na criação de categoria deveria conter o erro " + luuid.UUID_ERROR_NEW.Error())
    assert.Zero(
                    category,
                    "A categoria retornada ao ocorrer erro deveria ser zero")
}

// Testa a instanciação de um novo Link para casos onde ocorrem erros de UUID
func TestNewLink(t *testing.T) {
    assert := assert.New(t)

    link, err := NewLink(
        LinkMockMap["url"].(string),
        LinkMockMap["name"].(string),
        LinkMockMap["description"].(string),
        LinkMockMap["useMarkdown"].(bool),
    )

    assert.ErrorIs(
                    err,
                    LINK_ERROR_NEW,
                    "Erro na criação de link deveria conter o erro " + LINK_ERROR_NEW.Error())
    assert.ErrorIs(
                    err,
                    luuid.UUID_ERROR_NEW,
                    "Erro na criação de link deveria conter o erro " + luuid.UUID_ERROR_NEW.Error())
    assert.Zero(
                    link,
                    "O link retornado ao ocorrer erro deveria ser zero")
}
