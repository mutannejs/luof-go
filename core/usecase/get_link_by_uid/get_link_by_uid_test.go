package get_link_by_uid

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

func TestGetLinkByUid_NotExists(t *testing.T) {
	var assert = assert.New(t)

	var repo = repository.NewLinkMockRepository()
	var glbu = New(repo)

	repo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, nil)

	link, err := glbu.Execute(mockUidLink)

	assert.Zero(
		link,
		"Deveria ser retornado zero para um uid inválido")
	assert.Equal(
		ltests.GetMsgError(err),
		linkNotExists,
		"Buscar um link que não existe deveria retornar erro contendo " + linkNotExists)
}

func TestGetLinkByUid_Exists(t *testing.T) {
	var assert = assert.New(t)

	var repo = repository.NewLinkMockRepository()
	var glbu = New(repo)

	repo.On("Exists", mockUidLink).Return(true, nil)
	repo.On("GetByUid", mockUidLink).Return(mockLink, nil)

	link, err := glbu.Execute(mockUidLink)

	assert.Equal(
		mockLink,
		link,
		"O link retornado pela função deve ser o mesmo retornado pelo repositório")
	assert.True(
		err.IsNil(),
		"Buscar um link válido não deveria retornar erro")
}
