package delete_link

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
    linkNotExists = repository.LinkNotExists
    mockUidLink = domain.MockUidLink
)

func TestDeleteLink_NotExists(t *testing.T) {
    var assert = assert.New(t)

    var repo = repository.NewLinkMockRepository()
    var dl = New(repo)

    repo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, linkNotExists)

    exists, err := dl.Execute(mockUidLink)

    assert.False(
                    exists,
                    "Não deveria ser possível deletar um link que não existe")
    assert.ErrorIs(
                    err,
                    linkNotExists,
                    "Tentativa de deletar um link que não existe deveria retornar erro")
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
    assert.NoError(
                    err,
                    "Deletar um link válido não deveria retornar erro")
}
