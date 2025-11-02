package get_category_by_uid

import (
    "testing"

    "github.com/mutannejs/luof-go/core/usecase"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

var (
    categoryNotExists = usecase.CategoryNotExists
    mockCategory = usecase.MockCategory
    mockUidCategory = usecase.MockUidCategory
)

func TestGetCategoryByUid_NotExists(t *testing.T) {
    var assert = assert.New(t)

    var repo = usecase.NewCategoryMockRepository()
    var gcbu = New(repo)

    repo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, categoryNotExists)

    category, err := gcbu.Execute(mockUidCategory)

    assert.Zero(
                    category,
                    "Deveria ser retornado zero para um uid inválido")
    assert.ErrorIs(
                    err,
                    categoryNotExists,
                    "Buscar uma categoria que não existe deveria retornar erro contendo " + categoryNotExists.Error())
}

func TestGetCategoryByUid_Exists(t *testing.T) {
    var assert = assert.New(t)

    var repo = usecase.NewCategoryMockRepository()
    var gcbu = New(repo)

    repo.On("Exists", mockUidCategory).Return(true, nil)
    repo.On("GetByUid", mockUidCategory).Return(mockCategory, nil)

    category, err := gcbu.Execute(mockUidCategory)

    assert.Equal(
                    category,
                    mockCategory,
                    "A categoria retornada pela função deve ser a mesma retornada pelo repositório")
    assert.NoError(
                    err,
                    "Buscar uma categoria válida não deveria retornar erro")
}
