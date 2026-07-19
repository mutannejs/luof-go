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

	assert.Errorf(
		err,
		"%s: %s", CATEGORY_ERROR_NEW, luuid.UUID_ERROR_NEW)
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

	assert.Errorf(
		err,
		"%s: %s", LINK_ERROR_NEW, luuid.UUID_ERROR_NEW)
	assert.Zero(
		link,
		"O link retornado ao ocorrer erro deveria ser zero")
}
