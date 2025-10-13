package in

import (
    "github.com/mutannejs/luof-go/core/domain"

    "fmt"
    "testing"
)

type RepoMock struct {}

func (*RepoMock) Create (l domain.Link) error {
    return nil
}

func TestCreateLinkService(t *testing.T) {
    linkService := NewLinkService(&RepoMock{})
    uid, err := linkService.Create("url", "name", "description")
    fmt.Println(uid, err)
    if err != nil {
        t.Fail()
    }
}
