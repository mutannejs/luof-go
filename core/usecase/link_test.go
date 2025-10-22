package usecase

import (
    "errors"
    "testing"

    "github.com/mutannejs/luof-go/core/domain"
    "github.com/mutannejs/luof-go/pkg/ltests"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

var (
    LINK_NOT_EXISTS = errors.New("not exists")
    mockLink, _ = domain.NewLink(
        "github.com/mutannejs/luof",
        "luof",
        "luof repository",
        false,
    )
    mockUidLink = mockLink.GetUid()
)

func NewLinkMockRepository() *ltests.MockCrudRepository[domain.Link] {
    return &ltests.MockCrudRepository[domain.Link]{}
}

func TestCreateLink(t *testing.T) {
    var assert = assert.New(t)
    var repo = NewLinkMockRepository()
    var cl = NewCreateLink(repo)

    repo.On("Create", mock.MatchedBy(func(l domain.Link) bool {
        return assert.True(
            assert.Equal(l.Url, mockLink.Url),
            assert.Equal(l.Name, mockLink.Name),
            assert.Equal(l.Description.Content, mockLink.Description.Content),
            assert.Equal(l.Description.UseMarkdown, mockLink.Description.UseMarkdown),
            assert.NotZero(l.CreatedAt),
            assert.Zero(l.UpdatedAt),
        )
    })).Return(nil)

    uid, err := cl.Execute(
        mockLink.Url,
        mockLink.Name,
        mockLink.Description.Content,
        mockLink.Description.UseMarkdown,
    )

    assert.NotZero(uid, "criação com dados válidos deveria retornar um uuid diferente de zero")
    assert.NoError(err, "criação com dados válidos não deveria retornar erro")
}

func TestDeleteLink_NotExists(t *testing.T) {
    var assert = assert.New(t)
    var repo = NewLinkMockRepository()
    var dl = NewDeleteLink(repo)

    repo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, LINK_NOT_EXISTS)

    exists, err := dl.Execute(mockUidLink)

    assert.False(exists, "não deveria ser possível deletar um link que não existe")
    assert.EqualError(err, LINK_NOT_EXISTS.Error(), "tentativa de deletar um link que não existe deveria retornar erro")
}

func TestDeleteLink_Exists(t *testing.T) {
    var assert = assert.New(t)
    var repo = NewLinkMockRepository()
    var dl = NewDeleteLink(repo)

    repo.On("Exists", mockUidLink).Return(true, nil)
    repo.On("Delete", mockUidLink).Return(nil)

    exists, err := dl.Execute(mockUidLink)

    assert.True(exists, "deveria ser possível deletar um link válida")
    assert.NoError(err, "deletar um link válida não deveria retornar erro")
}

func TestGetLinkByUid_NotExists(t *testing.T) {
    var assert = assert.New(t)
    var repo = NewLinkMockRepository()
    var glbu = NewGetLinkByUid(repo)

    repo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, LINK_NOT_EXISTS)

    link, err := glbu.Execute(mockUidLink)

    assert.Zero(link, "deveria ser retornado zero para um uid inválido")
    assert.EqualError(err, LINK_NOT_EXISTS.Error(), "buscar um link que não existe deveria retornar erro")
}

func TestGetLinkByUid_Exists(t *testing.T) {
    var assert = assert.New(t)
    var repo = NewLinkMockRepository()
    var glbu = NewGetLinkByUid(repo)

    repo.On("Exists", mockUidLink).Return(true, nil)
    repo.On("GetByUid", mockUidLink).Return(mockLink, nil)

    link, err := glbu.Execute(mockUidLink)

    assert.Equal(link.Url, mockLink.Url)
    assert.Equal(link.Name, mockLink.Name)
    assert.Equal(link.Description.Content, mockLink.Description.Content)
    assert.Equal(link.Description.UseMarkdown, mockLink.Description.UseMarkdown)
    assert.Equal(link.CreatedAt, mockLink.CreatedAt)
    assert.Equal(link.UpdatedAt, mockLink.UpdatedAt)
    assert.NoError(err, "buscar um link válida não deveria retornar erro")
}

func TestUpdateLink_NotExists(t *testing.T) {
    var assert = assert.New(t)
    var repo = NewLinkMockRepository()
    var ul = NewUpdateLink(repo)

    repo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, LINK_NOT_EXISTS)

    exists, err := ul.Execute(
        mockUidLink,
        mockLink.Url,
        mockLink.Name,
        mockLink.Description.Content,
        mockLink.Description.UseMarkdown,
    )

    assert.False(exists, "não deveria ser possível atualizar um link que não existe")
    assert.EqualError(err, LINK_NOT_EXISTS.Error(), "tentar atualizar um link que não existe deveria retornar erro")
}

func TestUpdateLink_Exists(t *testing.T) {
    var assert = assert.New(t)
    var repo = NewLinkMockRepository()
    var ul = NewUpdateLink(repo)

    repo.On("Exists", mockUidLink).Return(true, nil)
    repo.On("Update", mockUidLink, mock.MatchedBy(func(l domain.Link) bool {
        return assert.True(
            assert.Equal(l.Url, mockLink.Url),
            assert.Equal(l.Name, mockLink.Name),
            assert.Equal(l.Description.Content, mockLink.Description.Content),
            assert.Equal(l.Description.UseMarkdown, mockLink.Description.UseMarkdown),
            assert.NotZero(l.CreatedAt),
            assert.NotZero(l.UpdatedAt),
        )
    })).Return(nil)

    exists, err := ul.Execute(
        mockUidLink,
        mockLink.Url,
        mockLink.Name,
        mockLink.Description.Content,
        mockLink.Description.UseMarkdown,
    )

    assert.True(exists, "deveria ser possível atualizar um link válida")
    assert.NoError(err, "atualizar um link válida não deveria retornar erro")
}
