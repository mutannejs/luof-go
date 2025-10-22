package usecase

import (
    "errors"
    "testing"

    "github.com/mutannejs/luof-go/core/domain"
    "github.com/mutannejs/luof-go/pkg/ltests"

    "github.com/google/uuid"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

var (
    NOT_EXISTS = errors.New("not exists")
    mockCategory, _ = domain.NewCategory(
        "development",
        "links about development",
        false,
    )
    mockUidCategory = mockCategory.GetUid()
)

type CategoryRepository struct {
    Categories map[uuid.UUID]domain.Category
}

func NewMockRepository() *ltests.MockTestifyRepository[domain.Category] {
    return &ltests.MockTestifyRepository[domain.Category]{}
}

func TestCreateCategory(t *testing.T) {
    var assert = assert.New(t)
    var repo = NewMockRepository()
    var cc = NewCreateCategory(repo)

    repo.On("Create", mock.MatchedBy(func(c domain.Category) bool {
        return assert.True(
            assert.Equal(c.Name, mockCategory.Name),
            assert.Equal(c.Description.Content, mockCategory.Description.Content),
            assert.Equal(c.Description.UseMarkdown, mockCategory.Description.UseMarkdown),
            assert.NotZero(c.CreatedAt),
            assert.Zero(c.UpdatedAt),
        )
    })).Return(nil)

    uid, err := cc.Execute(
        mockCategory.Name,
        mockCategory.Description.Content,
        mockCategory.Description.UseMarkdown,
    )

    assert.NotZero(uid, "criação com dados válidos deveria retornar um uuid diferente de zero")
    assert.NoError(err, "criação com dados válidos não deveria retornar erro")
}

func TestDeleteCategory_NotExists(t *testing.T) {
    var assert = assert.New(t)
    var repo = NewMockRepository()
    var dc = NewDeleteCategory(repo)

    repo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, NOT_EXISTS)

    exists, err := dc.Execute(mockUidCategory)

    assert.False(exists, "não deveria ser possível deletar uma categoria que não existe")
    assert.EqualError(err, NOT_EXISTS.Error(), "tentativa de deletar uma categoria que não existe deveria retornar erro")
}

func TestDeleteCategory_Exists(t *testing.T) {
    var assert = assert.New(t)
    var repo = NewMockRepository()
    var dc = NewDeleteCategory(repo)

    repo.On("Exists", mockUidCategory).Return(true, nil)
    repo.On("Delete", mockUidCategory).Return(nil)

    exists, err := dc.Execute(mockUidCategory)

    assert.True(exists, "deveria ser possível deletar uma categoria válida")
    assert.NoError(err, "deletar uma categoria válida não deveria retornar erro")
}

func TestGetCategoryByUid_NotExists(t *testing.T) {
    var assert = assert.New(t)
    var repo = NewMockRepository()
    var gcbu = NewGetCategoryByUid(repo)

    repo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, NOT_EXISTS)

    category, err := gcbu.Execute(mockUidCategory)

    assert.Zero(category, "deveria ser retornado zero para um uid inválido")
    assert.EqualError(err, NOT_EXISTS.Error(), "buscar uma categoria que não existe deveria retornar erro")
}

func TestGetCategoryByUid_Exists(t *testing.T) {
    var assert = assert.New(t)
    var repo = NewMockRepository()
    var gcbu = NewGetCategoryByUid(repo)

    repo.On("Exists", mockUidCategory).Return(true, nil)
    repo.On("GetByUid", mockUidCategory).Return(mockCategory, nil)

    category, err := gcbu.Execute(mockUidCategory)

    assert.Equal(category.Name, mockCategory.Name)
    assert.Equal(category.Description.Content, mockCategory.Description.Content)
    assert.Equal(category.Description.UseMarkdown, mockCategory.Description.UseMarkdown)
    assert.Equal(category.CreatedAt, mockCategory.CreatedAt)
    assert.Equal(category.UpdatedAt, mockCategory.UpdatedAt)
    assert.NoError(err, "buscar uma categoria válida não deveria retornar erro")
}

func TestUpdateCategory_NotExists(t *testing.T) {
    var assert = assert.New(t)
    var repo = NewMockRepository()
    var uc = NewUpdateCategory(repo)

    repo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, NOT_EXISTS)

    exists, err := uc.Execute(
        mockUidCategory,
        mockCategory.Name,
        mockCategory.Description.Content,
        mockCategory.Description.UseMarkdown,
    )

    assert.False(exists, "não deveria ser possível atualizar uma categoria que não existe")
    assert.EqualError(err, NOT_EXISTS.Error(), "tentar atualizar uma categoria que não existe deveria retornar erro")
}

func TestUpdateCategory_Exists(t *testing.T) {
    var assert = assert.New(t)
    var repo = NewMockRepository()
    var uc = NewUpdateCategory(repo)

    repo.On("Exists", mockUidCategory).Return(true, nil)
    repo.On("Update", mockUidCategory, mock.MatchedBy(func(c domain.Category) bool {
        return assert.True(
            assert.Equal(c.Name, mockCategory.Name),
            assert.Equal(c.Description.Content, mockCategory.Description.Content),
            assert.Equal(c.Description.UseMarkdown, mockCategory.Description.UseMarkdown),
            assert.NotZero(c.CreatedAt),
            assert.NotZero(c.UpdatedAt),
        )
    })).Return(nil)

    exists, err := uc.Execute(
        mockUidCategory,
        mockCategory.Name,
        mockCategory.Description.Content,
        mockCategory.Description.UseMarkdown,
    )

    assert.True(exists, "deveria ser possível atualizar uma categoria válida")
    assert.NoError(err, "atualizar uma categoria válida não deveria retornar erro")
}
