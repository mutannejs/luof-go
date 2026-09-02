package delete_link

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
	mockUidLink = domain.MockUidLink
)

func TestDeleteLink_NotExists(t *testing.T) {
	var assert = assert.New(t)

	var repo = repository.NewLinkMockRepository()
	var dl = New(repo)

	repo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, nil)

	exists, err := dl.Execute(mockUidLink)

	assert.False(
		exists,
		"Não deveria ser possível deletar um link que não existe")
	assert.Equal(
		ltests.GetMsgError(err),
		linkNotExists,
		"Tentativa de deletar um link que não existe deveria retornar erro contendo " + linkNotExists)
}

func TestDeleteLink_Exists(t *testing.T) {
	var assert = assert.New(t)

	var repo = repository.NewLinkMockRepository()
	var dl = New(repo)

	repo.On("Exists", mockUidLink).Return(true, nil)
	repo.On("Delete", mockUidLink).Return(nil)

	exists, err := dl.Execute(mockUidLink)

	assert.True(
		exists,
		"Deveria ser possível deletar um link válido")
	assert.True(
		err.IsNil(),
		"Deletar um link válido não deveria retornar erro")
}
