package get_link_by_uid

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
    linkNotExists = repository.LinkNotExists
    mockLink = domain.MockLink
    mockUidLink = domain.MockUidLink
)

func TestGetLinkByUid_NotExists(t *testing.T) {
    var assert = assert.New(t)

    var repo = repository.NewLinkMockRepository()
    var glbu = NewGetLinkByUid(repo)

    repo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, linkNotExists)

    link, err := glbu.Execute(mockUidLink)

    assert.Zero(
                    link,
                    "Deveria ser retornado zero para um uid inválido")
    assert.EqualError(
                    err,
                    linkNotExists.Error(),
                    "Buscar um link que não existe deveria retornar erro")
}

func TestGetLinkByUid_Exists(t *testing.T) {
    var assert = assert.New(t)

    var repo = repository.NewLinkMockRepository()
    var glbu = NewGetLinkByUid(repo)

    repo.On("Exists", mockUidLink).Return(true, nil)
    repo.On("GetByUid", mockUidLink).Return(mockLink, nil)

    link, err := glbu.Execute(mockUidLink)

    assert.Equal(
                    link,
                    mockLink,
                    "O link retornado pela função deve ser o mesmo retornado pelo repositório")
    assert.NoError(
                    err,
                    "Buscar um link válido não deveria retornar erro")
}
