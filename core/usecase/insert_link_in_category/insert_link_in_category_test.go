package insert_link_in_category

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	alreadyBelongs = domain.ALREADY_BELONGS
	categoryNotExists = domain.CATEGORY_NOT_EXISTS
	linkNotExists = domain.LINK_NOT_EXISTS
	mockUidLink = domain.MockUidLink
	mockUidCategory = domain.MockUidCategory
)

func TestInsertLinkInCategory_CategoryNotExists(t *testing.T) {
	var assert = assert.New(t)

	var btRepo = repository.NewBelongsToMockRepository()
	var cRepo = repository.NewCategoryMockRepository()
	var lRepo = repository.NewLinkMockRepository()
	var ilic = New(btRepo, cRepo, lRepo)

	btRepo.
		On(
			"Exists",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(false, nil).
		On(
			"Create",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("time.Time"),
			true,
		).Return(nil).
		On(
			"SetHasNoMainCategory",
			mock.AnythingOfType("uuid.UUID"),
		).Return(nil)

	cRepo.
		On(
			"Exists",
			mock.AnythingOfType("uuid.UUID"),
		).Return(false, nil)

	lRepo.
		On(
			"Exists",
			mock.AnythingOfType("uuid.UUID"),
		).Return(true, nil)

	err := ilic.Execute(mockUidLink, mockUidCategory, true)

	assert.Error(
		categoryNotExists,
		err,
		"Tentar inserir um link em uma categoria que não existe deveria retornar erro contendo " + categoryNotExists.Error())
}

func TestInsertLinkInCategory_Exists(t *testing.T) {
	var assert = assert.New(t)

	var btRepo = repository.NewBelongsToMockRepository()
	var cRepo = repository.NewCategoryMockRepository()
	var lRepo = repository.NewLinkMockRepository()
	var ilic = New(btRepo, cRepo, lRepo)

	btRepo.
		On(
			"Exists",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(true, nil).
		On(
			"Create",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("time.Time"),
			false,
		).Return(nil).
		On(
			"SetHasNoMainCategory",
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

	err := ilic.Execute(mockUidLink, mockUidCategory, false)

	assert.ErrorIs(
		alreadyBelongs,
		err,
		"Tentar inserir um link em uma categoria, ambos já relacionados, deveria retornar erro contendo " + alreadyBelongs.Error())
}

func TestInsertLinkInCategory_LinkNotExists(t *testing.T) {
	var assert = assert.New(t)

	var btRepo = repository.NewBelongsToMockRepository()
	var cRepo = repository.NewCategoryMockRepository()
	var lRepo = repository.NewLinkMockRepository()
	var ilic = New(btRepo, cRepo, lRepo)

	btRepo.
		On(
			"Exists",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(false, nil).
		On(
			"Create",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("time.Time"),
			true,
		).Return(nil).
		On(
			"SetHasNoMainCategory",
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
		).Return(false, nil)

	err := ilic.Execute(mockUidLink, mockUidCategory, true)

	assert.Error(
		linkNotExists,
		err,
		"Tentar inserir um link que não existe em uma categoria válida deveria retornar erro contendo " + linkNotExists.Error())
}

func TestInsertLinkInCategory_NotExists(t *testing.T) {
	var assert = assert.New(t)

	var btRepo = repository.NewBelongsToMockRepository()
	var cRepo = repository.NewCategoryMockRepository()
	var lRepo = repository.NewLinkMockRepository()
	var ilic = New(btRepo, cRepo, lRepo)

	btRepo.
		On(
			"Exists",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
		).Return(false, nil).
		On(
			"Create",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("time.Time"),
			true,
		).Return(nil).
		On(
			"SetHasNoMainCategory",
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

	err := ilic.Execute(mockUidLink, mockUidCategory, true)

	assert.NoError(
		err,
		"Tentar inserir um link em uma categoria, ambos ainda não relacionados, não deveria retornar erro")
}
