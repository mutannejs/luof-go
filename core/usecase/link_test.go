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
    linkNotExists = errors.New(repository.LINK_NOT_EXISTS)
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
    var link domain.Link

    var repo = NewLinkMockRepository()
    var cl = NewCreateLink(repo)

    repo.On("Create", mock.MatchedBy(func(l domain.Link) bool {
        link = l
        return true
    })).Return(nil)

    uid, err := cl.Execute(
        mockLink.Url,
        mockLink.Name,
        mockLink.Description.Content,
        mockLink.Description.UseMarkdown,
    )

    // Testa se o link enviado para Repository.Create está de acordo com os
    // argumentos passados à função
    // ! Pode estar errada, mesmo que NewLink tenha sido corretamente
    // implementada
    assert.Equal(link.Url, mockLink.Url)
    assert.Equal(link.Name, mockLink.Name)
    assert.Equal(link.Description.Content, mockLink.Description.Content)
    assert.Equal(link.Description.UseMarkdown, mockLink.Description.UseMarkdown)
    assert.NotZero(link.CreatedAt)
    assert.Zero(link.UpdatedAt)

    // Testa o retorno da função
    assert.NotZero(
                    uid,
                    "Criação com dados válidos deveria retornar um uuid diferente de zero")
    assert.NoError( 
                    err,
                    "Criação com dados válidos não deveria retornar erro")
}

func TestDeleteLink_NotExists(t *testing.T) {
    var assert = assert.New(t)

    var repo = NewLinkMockRepository()
    var dl = NewDeleteLink(repo)

    repo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, linkNotExists)

    exists, err := dl.Execute(mockUidLink)

    assert.False(
                    exists,
                    "Não deveria ser possível deletar um link que não existe")
    assert.ErrorIs(
                    err,
                    linkNotExists,
                    "Tentativa de deletar um link que não existe deveria retornar erro")
}

func TestDeleteLink_Exists(t *testing.T) {
    var assert = assert.New(t)

    var repo = NewLinkMockRepository()
    var dl = NewDeleteLink(repo)

    repo.On("Exists", mockUidLink).Return(true, nil)
    repo.On("Delete", mockUidLink).Return(nil)

    exists, err := dl.Execute(mockUidLink)

    assert.True(
                    exists,
                    "Deveria ser possível deletar um link válido")
    assert.NoError(
                    err,
                    "Deletar um link válido não deveria retornar erro")
}

func TestGetLinkByUid_NotExists(t *testing.T) {
    var assert = assert.New(t)

    var repo = NewLinkMockRepository()
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

    var repo = NewLinkMockRepository()
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

func TestUpdateLink_NotExists(t *testing.T) {
    var assert = assert.New(t)

    var repo = NewLinkMockRepository()
    var ul = NewUpdateLink(repo)

    repo.On("Exists", mock.AnythingOfType("uuid.UUID")).Return(false, linkNotExists)

    exists, err := ul.Execute(
        mockUidLink,
        mockLink.Url,
        mockLink.Name,
        mockLink.Description.Content,
        mockLink.Description.UseMarkdown,
    )

    assert.False(
                    exists,
                    "Não deveria ser possível atualizar um link que não existe")
    assert.EqualError(
                    err,
                    linkNotExists.Error(),
                    "Tentar atualizar um link que não existe deveria retornar erro")
}

func TestUpdateLink_Exists(t *testing.T) {
    var assert = assert.New(t)
    var link domain.Link

    var repo = NewLinkMockRepository()
    var ul = NewUpdateLink(repo)

    repo.On("Exists", mockUidLink).Return(true, nil)
    repo.On("Update", mockUidLink, mock.MatchedBy(func(l domain.Link) bool {
        link = l
        return true
    })).Return(nil)

    exists, err := ul.Execute(
        mockUidLink,
        mockLink.Url,
        mockLink.Name,
        mockLink.Description.Content,
        mockLink.Description.UseMarkdown,
    )

    // Testa se o link enviado para Repository.Update está de acordo com os
    // argumentos passados à função
    // ! Pode estar errada, mesmo que NewLink tenha sido corretamente
    // implementada
    assert.Equal(link.Url, mockLink.Url)
    assert.Equal(link.Name, mockLink.Name)
    assert.Equal(link.Description.Content, mockLink.Description.Content)
    assert.Equal(link.Description.UseMarkdown, mockLink.Description.UseMarkdown)
    assert.NotZero(link.CreatedAt, mockLink.CreatedAt)

    // Testa o valor de UpdatedAt
    assert.Less(
                    link.CreatedAt,
                    link.UpdatedAt,
                    "O valor de UpdatedAt deve ser maior que o valor de CreatedAt")

    // Testa o retorno da função
    assert.True(
                    exists,
                    "Deveria ser possível atualizar um link válido")
    assert.NoError(
                    err,
                    "Atualizar um link válido não deveria retornar erro")
}
