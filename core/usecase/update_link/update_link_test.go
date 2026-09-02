package update_link

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/ltests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	linkNotExists = domain.LINK_NOT_EXISTS
	mockLink = domain.MockLink
	mockUidLink = domain.MockUidLink
)

func TestUpdateLink_NotExists(t *testing.T) {
	var assert = assert.New(t)

	var repo = repository.NewLinkMockRepository()
	var ul = New(repo)

	repo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, nil)

	exists, err := ul.Execute(
		mockUidLink,
		mockLink.Url,
		mockLink.Name,
		mockLink.Description.Content,
		mockLink.Description.UseMarkdown,
	)

	assert.False(
		exists,
		"Não deveria ser possível atualizar um link que não existe")
	assert.Equal(
		ltests.GetMsgError(err),
		linkNotExists,
		"Tentar atualizar um link que não existe deveria retornar erro contendo " + linkNotExists)
}

func TestUpdateLink_Exists(t *testing.T) {
	var assert = assert.New(t)
	var link domain.Link

	var repo = repository.NewLinkMockRepository()
	var ul = New(repo)

	repo.On("Exists", mockUidLink).Return(true, nil)
	repo.On("Update", mockUidLink, mock.MatchedBy(func(l domain.Link) bool {
		link = l
		return true
	})).Return(nil)

	exists, err := ul.Execute(
		mockUidLink,
		mockLink.Url,
		mockLink.Name,
		mockLink.Description.Content,
		mockLink.Description.UseMarkdown,
	)

	// Testa se o link enviado para Repository.Update está de acordo com os
	// argumentos passados à função
	// ! Pode estar errada, mesmo que NewLink tenha sido corretamente
	// implementada
	assert.Equal(link.Url, mockLink.Url)
	assert.Equal(link.Name, mockLink.Name)
	assert.Equal(link.Description.Content, mockLink.Description.Content)
	assert.Equal(link.Description.UseMarkdown, mockLink.Description.UseMarkdown)
	assert.NotZero(link.CreatedAt, mockLink.CreatedAt)

	// Testa o valor de UpdatedAt
	assert.Less(
		link.CreatedAt,
		link.UpdatedAt,
		"O valor de UpdatedAt deve ser maior que o valor de CreatedAt")

	// Testa o retorno da função
	assert.True(
		exists,
		"Deveria ser possível atualizar um link válido")
	assert.True(
		err.IsNil(),
		"Atualizar um link válido não deveria retornar erro")
}
