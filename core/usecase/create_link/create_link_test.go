package create_link

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateLink(t *testing.T) {
	var assert = assert.New(t)
	var mockLink = domain.MockLink
	var link domain.Link

	var repo = repository.NewLinkMockRepository()
	var cl = New(repo)

	repo.On("Create", mock.MatchedBy(func(l domain.Link) bool {
		link = l
		return true
	})).Return(nil)

	uid, err := cl.Execute(
		mockLink.Url,
		mockLink.Name,
		mockLink.Description.Content,
		mockLink.Description.UseMarkdown,
	)

	// Testa se o link enviado para Repository.Create está de acordo com os
	// argumentos passados à função
	// ! Pode estar errada, mesmo que NewLink tenha sido corretamente
	// implementada
	assert.Equal(link.Url, mockLink.Url)
	assert.Equal(link.Name, mockLink.Name)
	assert.Equal(link.Description.Content, mockLink.Description.Content)
	assert.Equal(link.Description.UseMarkdown, mockLink.Description.UseMarkdown)
	assert.NotZero(link.CreatedAt)
	assert.Zero(link.UpdatedAt)

	// Testa o retorno da função
	assert.NotZero(
					uid,
					"Criação com dados válidos deveria retornar um uuid diferente de zero")
	assert.NoError( 
					err,
					"Criação com dados válidos não deveria retornar erro")
}
