package usecase

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
    categoryNotExists = errors.New(repository.CATEGORY_NOT_EXISTS)
    mockCategory, _ = domain.NewCategory(
        "development",
        "links about development",
        false,
    )
    mockUidCategory = mockCategory.GetUid()
)

func NewCategoryMockRepository() *ltests.MockCrudRepository[domain.Category] {
    return &ltests.MockCrudRepository[domain.Category]{}
}

func TestCreateCategory(t *testing.T) {
    var assert = assert.New(t)
    var category domain.Category

    var repo = NewCategoryMockRepository()
    var cc = NewCreateCategory(repo)

    repo.On("Create", mock.MatchedBy(func(c domain.Category) bool {
        category = c
        return true
    })).Return(nil)

    uid, err := cc.Execute(
        mockCategory.Name,
        mockCategory.Description.Content,
        mockCategory.Description.UseMarkdown,
    )

    // Testa se a categoria enviada para Repository.Create está de acordo
    // com os argumentos passados à função
    // ! Pode estar errada, mesmo que NewCategory tenha sido corretamente
    // implementada
    assert.Equal(category.Name, mockCategory.Name)
    assert.Equal(category.Description.Content, mockCategory.Description.Content)
    assert.Equal(category.Description.UseMarkdown, mockCategory.Description.UseMarkdown)
    assert.NotZero(category.CreatedAt)
    assert.Zero(category.UpdatedAt)

    // Testa o retorno da função
    assert.NotZero(
                    uid,
                    "Criação com dados válidos deveria retornar um uuid diferente de zero")
    assert.NoError(
                    err,
                    "Criação com dados válidos não deveria retornar erro")
}

func TestDeleteCategory_NotExists(t *testing.T) {
    var assert = assert.New(t)

    var repo = NewCategoryMockRepository()
    var dc = NewDeleteCategory(repo)

    repo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, categoryNotExists)

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

    var repo = NewCategoryMockRepository()
    var dc = NewDeleteCategory(repo)

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

func TestGetCategoryByUid_NotExists(t *testing.T) {
    var assert = assert.New(t)

    var repo = NewCategoryMockRepository()
    var gcbu = NewGetCategoryByUid(repo)

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

    var repo = NewCategoryMockRepository()
    var gcbu = NewGetCategoryByUid(repo)

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

func TestUpdateCategory_NotExists(t *testing.T) {
    var assert = assert.New(t)

    var repo = NewCategoryMockRepository()
    var uc = NewUpdateCategory(repo)

    repo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, categoryNotExists)

    exists, err := uc.Execute(
        mockUidCategory,
        mockCategory.Name,
        mockCategory.Description.Content,
        mockCategory.Description.UseMarkdown,
    )

    assert.False(
                    exists,
                    "Não deveria ser possível atualizar uma categoria que não existe")
    assert.ErrorIs(
                    err,
                    categoryNotExists,
                    "Tentar atualizar uma categoria que não existe deveria retornar erro contendo " + categoryNotExists.Error())
}

func TestUpdateCategory_Exists(t *testing.T) {
    var assert = assert.New(t)
    var category domain.Category

    var repo = NewCategoryMockRepository()
    var uc = NewUpdateCategory(repo)

    repo.On("Exists", mockUidCategory).Return(true, nil)
    repo.On("Update", mockUidCategory, mock.MatchedBy(func(c domain.Category) bool {
        category = c
        return true
    })).Return(nil)

    exists, err := uc.Execute(
        mockUidCategory,
        mockCategory.Name,
        mockCategory.Description.Content,
        mockCategory.Description.UseMarkdown,
    )

    // Testa se a categoria enviada para Repository.Update está de acordo
    // com os argumentos passados à função
    // ! Pode estar errada, mesmo que NewCategory tenha sido corretamente
    // implementada
    assert.Equal(category.Name, mockCategory.Name)
    assert.Equal(category.Description.Content, mockCategory.Description.Content)
    assert.Equal(category.Description.UseMarkdown, mockCategory.Description.UseMarkdown)
    assert.NotZero(category.CreatedAt)

    // Testa o valor de UpdatedAt
    assert.Less(
                    category.CreatedAt,
                    category.UpdatedAt,
                    "O valor de UpdatedAt deve ser maior que o valor de CreatedAt")

    // Testa o retorno da função
    assert.True(
                    exists,
                    "Deveria ser possível atualizar uma categoria válida")
    assert.NoError(
                    err,
                    "Atualizar uma categoria válida não deveria retornar erro")
}
