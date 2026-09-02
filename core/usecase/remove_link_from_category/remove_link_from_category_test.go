package remove_link_from_category

import (
	"errors"
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/ltests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	notBelongs = domain.NOT_BELONGS
	mockUidLink = domain.MockUidLink
	mockUidCategory = domain.MockUidCategory
)

func TestRemoveLinkFromCategory_NotExists(t *testing.T) {
	var assert = assert.New(t)

	var btRepo = repository.NewBelongsToMockRepository()
	var cRepo = repository.NewCategoryMockRepository()
	var lRepo = repository.NewLinkMockRepository()
	var rlfc = New(btRepo, cRepo, lRepo)

	btRepo.
		On(
			"Exists",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(false, nil).
		On(
			"Delete",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(errors.New(""))

	cRepo.
		On(
			"Exists",
			mock.AnythingOfType("uuid.UUID"),
		).Return(true, nil)

	lRepo.
		On(
			"Exists",
			mock.AnythingOfType("uuid.UUID"),
		).Return(true, nil)

	err := rlfc.Execute(mockUidLink, mockUidCategory)

	assert.Equal(
		ltests.GetMsgError(err),
		notBelongs,
		"Tentar remover um link de uma categoria, ambos não relacionados, deveria retornar erro contendo " + notBelongs)
}

func TestRemoveLinkFromCategory_Exists(t *testing.T) {
	var assert = assert.New(t)
	
	var btRepo = repository.NewBelongsToMockRepository()
	var cRepo = repository.NewCategoryMockRepository()
	var lRepo = repository.NewLinkMockRepository()
	var rlfc = New(btRepo, cRepo, lRepo)

	btRepo.
		On(
			"Exists",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(true, nil)

	btRepo.
		On(
			"Delete",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(nil)

	cRepo.
		On(
			"Exists",
			mock.AnythingOfType("uuid.UUID"),
		).Return(true, nil)

	lRepo.
		On(
			"Exists",
			mock.AnythingOfType("uuid.UUID"),
		).Return(true, nil)

	err := rlfc.Execute(mockUidLink, mockUidCategory)

	assert.True(
		err.IsNil(),
		"Tentar remover um link de uma categoria, tal que essa relação existe, não deveria retornar erro")
}
