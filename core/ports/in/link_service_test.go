package in

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/mutannejs/luof-go/pkg/ltests"

    "fmt"
    "github.com/google/uuid"
    "slices"
    "testing"
)

type RepoMock struct {
    Links []domain.Link
}

func (repo *RepoMock) Create(l domain.Link) error {
    repo.Links = append(repo.Links, l)
    return nil
}

func (repo *RepoMock) Delete(uid uuid.UUID) error {
    repo.Links = slices.DeleteFunc(repo.Links, func (l domain.Link) bool {
        return l.Uid == uid
    })
    return nil
}

func (repo *RepoMock) Update(uid uuid.UUID, link domain.Link) error {
    var index = slices.IndexFunc(repo.Links, func (l domain.Link) bool {
        return l.Uid == uid
    })
    repo.Links[index].Url = link.Url
    repo.Links[index].Name = link.Name
    return nil
}

func (repo *RepoMock) Exists(uid uuid.UUID) (bool, error) {
    return slices.ContainsFunc(repo.Links, func (l domain.Link) bool {
        return l.Uid == uid
    }), nil
}

func TestCreateLinkService(t *testing.T) {
    test := ltests.NewTest(t)
    repoMock := &RepoMock{make([]domain.Link, 0)}
    linkService := NewLinkService(repoMock)

    uidGo, err := linkService.Create("https://go.dev/", "Golang", "Go language", false)
    test.FailIfError(err)

    uidLuof, err := linkService.Create("https://github.com/mutanne/luofgo", "luofgo", "Luof Front-end", true)
    test.FailIfError(err)

    _, err = linkService.Update(uidLuof, "https://github.com/mutannejs/luof-go", "luof-go", "Luof Back-end", false)
    test.FailIfError(err)

    _, err = linkService.Delete(uidGo)
    test.FailIfError(err)

    repoRM := linkService.Repo.(*RepoMock)
    fmt.Println("Length: ", len(repoRM.Links))
    for _, link := range repoRM.Links {
        fmt.Printf(
            "url: %v\nname: %v\nuseMarkdown: %v\ncreatedAt: %v\nupdateAt: %v\n\n",
            link.Url, link.Name, link.Description.UseMarkdown, link.CreatedAt, link.UpdatedAt,
        )
    }
}
