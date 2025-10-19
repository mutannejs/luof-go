package usecase

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/mutannejs/luof-go/pkg/ltests"
    "github.com/mutannejs/luof-go/pkg/luuid"
    "github.com/google/uuid"
    "strings"
    "testing"
)

var (
    mockLink, _ = domain.NewLink(
        "github.com/mutannejs/luof-go",
        "luof",
        "luof repository",
        false,
    )
    mockUidLink = mockLink.GetUid()
    invalidUidLink = uuid.New()
    mockUpdatedLink, _ = domain.NewLink(
        "https://go.dev/",
        "golang",
        "*Go documentation*",
        true,
    )
)

type LinkRepository struct {
    Links map[uuid.UUID]domain.Link
}

func NewMockLinkRepository (t *testing.T, populate bool, testName string) *ltests.MockCRUDRepository[domain.Link] {
    var repo = ltests.NewMockCRUDRepository[domain.Link](populate, mockLink)
    repo.TestRepository(t, testName)
    return repo
}

func TestCreateLink(t *testing.T) {
    var repo = NewMockLinkRepository(t, false, "CreateLink")
    var cl = NewCreateLink(repo)

    uid, err := cl.Execute(
        mockLink.Url,
        mockLink.Name,
        mockLink.Description.Content,
        mockLink.Description.UseMarkdown,
    )

    if luuid.IsZero(uid) ||
            err != nil ||
            repo.Length() != 1 ||
            strings.Compare(mockLink.Url, repo.GetItem(uid).Url) != 0 ||
            strings.Compare(mockLink.Name, repo.GetItem(uid).Name) != 0 ||
            strings.Compare(mockLink.Description.Content, repo.GetItem(uid).Description.Content) != 0 ||
            mockLink.Description.UseMarkdown != repo.GetItem(uid).Description.UseMarkdown {
        ltests.PrintAndFail(t, "Insucesso na execução de CreateLink", err)
    }
}

func TestDeleteLink(t *testing.T) {
    var repo = NewMockLinkRepository(t, true, "DeleteLink")
    var dl = NewDeleteLink(repo)

    exists, err := dl.Execute(mockUidLink)

    if !exists ||
            err != nil ||
            repo.Length() != 0 {
        ltests.PrintAndFail(t, "Insucesso na execução de DeleteLink para um uid válido", err)
    }

    exists, err = dl.Execute(invalidUidLink)

    if exists || err != nil {
        ltests.PrintAndFail(t, "Insucesso na execução de DeleteLink para um uid inválido", err)
    }
}

func TestGetLinkByUid(t *testing.T) {
    var repo = NewMockLinkRepository(t, true, "GetLinkByUid")
    var glbu = NewGetLinkByUid(repo)

    link, err := glbu.Execute(mockUidLink)

    if err != nil ||
            strings.Compare(mockLink.Url, link.Url) != 0 ||
            strings.Compare(mockLink.Name, link.Name) != 0 ||
            strings.Compare(mockLink.Description.Content, link.Description.Content) != 0 ||
            mockLink.Description.UseMarkdown != link.Description.UseMarkdown ||
            mockLink.CreatedAt.Compare(link.CreatedAt) != 0 ||
            mockLink.UpdatedAt.Compare(link.UpdatedAt) != 0 {
        ltests.PrintAndFail(t, "Insucesso na execução de GetLinkByUid para um uid válido", err)
    }

    link, err = glbu.Execute(invalidUidLink)

    if (link != domain.Link{}) || err != nil {
        ltests.PrintAndFail(t, "Insucesso na execução de GetLinkByUid para um uid inválido", err)
    }
}

func TestUpdateLink(t *testing.T) {
    var repo = NewMockLinkRepository(t, true, "UpdateLink")
    var ul = NewUpdateLink(repo)
    var glbu = NewGetLinkByUid(repo)

    exists, err := ul.Execute(
        mockUidLink,
        mockUpdatedLink.Url,
        mockUpdatedLink.Name,
        mockUpdatedLink.Description.Content,
        mockUpdatedLink.Description.UseMarkdown,
    )

    updatedLink, _ := glbu.Execute(mockUidLink)

    if err != nil ||
            /*updatedLink.CreatedAt.Compare(link.CreatedAt) == 0 ||
            updatedLink.UpdatedAt.Compare(link.UpdatedAt) <= 0 */
            !exists ||
            strings.Compare(updatedLink.Name, mockUpdatedLink.Name) != 0 ||
            strings.Compare(updatedLink.Description.Content, mockUpdatedLink.Description.Content) != 0 ||
            updatedLink.Description.UseMarkdown != mockUpdatedLink.Description.UseMarkdown {
        ltests.PrintAndFail(t, "Insucesso na execução de UpdateLink para um uid válido", err)
    }
}
