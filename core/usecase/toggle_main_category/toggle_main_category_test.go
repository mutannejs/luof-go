package toggle_main_category

import (
	"errors"
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	notBelongs = domain.NOT_BELONGS
	mockUidLink = domain.MockUidLink
	mockUidCategory = domain.MockUidCategory
)

func TestToggleMainCategory_NotExists(t *testing.T) {
	var assert = assert.New(t)

	var btRepo = repository.NewBelongsToMockRepository()
	var tmc = New(btRepo)

	btRepo.
		On(
			"Exists",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(false, nil).
		On(
			"Update",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
			true,
		).Return(errors.New(""))

	err := tmc.Execute(mockUidLink, mockUidCategory, true)

	assert.ErrorIs(
					err,
					notBelongs,
					"Tentar alterar a importância da categoria de um link, tal que a relação entre eles não existe, deveria retornar erro contendo " + notBelongs.Error())
}

func TestToggleMainCategory_Exists(t *testing.T) {
	var assert = assert.New(t)

	var btRepo = repository.NewBelongsToMockRepository()
	var tmc = New(btRepo)

	btRepo.
		On(
			"Exists",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(true, nil).
		On(
			"Update",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
			false,
		).Return(nil)

	err := tmc.Execute(mockUidLink, mockUidCategory, false)

	assert.NoError(
					err,
					"Tentar inserir um link em uma categoria, ambos já relacionados, não deveria retornar erro")
}
