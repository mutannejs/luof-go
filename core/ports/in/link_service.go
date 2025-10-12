package in

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/mutannejs/luof-go/core/ports/out"
    "github.com/mutannejs/luof-go/pkg/llog"

    "errors"
    "fmt"
    "time"
    "github.com/google/uuid"
)

type LinkService struct {
    Repo out.LinkRepository
}

func NewLinkService() LinkService {
    var repo out.LinkRepository
    return LinkService{repo}
}

func (ls *LinkService) Create (
    url string,
    name string,
    description string,
) (id uuid.UUID, err error) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println(llog.UUID_PREFIX, llog.UUID_ERROR_MESSAGE)
            err = errors.New(llog.UUID_ERROR_MESSAGE)
        }
    }()
    id = uuid.New() // possible panic
    var createdAt time.Time = time.Now()
    var updatedAt time.Time
    link := domain.Link{id, url, name, description, createdAt, updatedAt}
    err = ls.Repo.Create(link)
    return
}
