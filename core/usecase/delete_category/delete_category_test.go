package delete_category

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
    categoryNotExists = domain.CATEGORY_NOT_EXISTS
    mockUidCategory = domain.MockUidCategory
)

func TestDeleteCategory_NotExists(t *testing.T) {
    var assert = assert.New(t)

    var repo = repository.NewCategoryMockRepository()
    var dc = New(repo)

    repo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, nil)

    exists, err := dc.Execute(mockUidCategory)

    assert.False(
                    exists,
                    "Não deveria ser possível deletar uma categoria que não existe")
    assert.ErrorIs(
                    err,
                    categoryNotExists,
                    "Tentativa de deletar uma categoria que não existe deveria retornar erro contendo " + categoryNotExists.Error())
}

func TestDeleteCategory_Exists(t *testing.T) {
    var assert = assert.New(t)

    var repo = repository.NewCategoryMockRepository()
    var dc = New(repo)

    repo.On("Exists", mockUidCategory).Return(true, nil)
    repo.On("Delete", mockUidCategory).Return(nil)

    exists, err := dc.Execute(mockUidCategory)

    assert.True(
                    exists,
                    "Deveria ser possível deletar uma categoria válida")
    assert.NoError(
                    err,
                    "Deletar uma categoria válida não deveria retornar erro")
}
