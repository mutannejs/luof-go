package in

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/mutannejs/luof-go/core/ports/out"
    "github.com/mutannejs/luof-go/pkg/luuid"

    "github.com/google/uuid"
    "time"
)

type LinkService struct {
    Repo out.LinkRepository
}

func NewLinkService(repo out.LinkRepository) *LinkService {
    return &LinkService{repo}
}

func (ls *LinkService) Create (
    url string,
    name string,
    description string,
) (uid uuid.UUID, err error) {
    uid, err = luuid.New()
    if err != nil {
        return
    }

    var createdAt time.Time = time.Now()
    var updatedAt time.Time
    link := domain.Link{uid, url, name, description, createdAt, updatedAt}
    err = ls.Repo.Create(link)
    return
}
