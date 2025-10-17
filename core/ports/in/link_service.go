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

func (ls *LinkService) Create(
    url string,
    name string,
    contentDescription string,
    useMarkdown bool,
    categoryUid uuid.UUID,
    isMain bool,
) (uid uuid.UUID, err error) {
    if uid, err = luuid.New(); err != nil {
        return
    }

    var description = domain.Description{contentDescription, useMarkdown}
    var createdAt time.Time = time.Now()
    var insertedAt time.Time = createdAt
    var updatedAt time.Time
    var link = domain.Link{uid, url, name, description, createdAt, updatedAt}

    if luuid.IsZero(categoryUid) {
        err = ls.Repo.CreateWithCategory(link, categoyUid, isMain, insertedAt)
    } else {
        err = ls.Repo.Create(link)
    }
    return
}

func (ls *LinkService) Delete(
    uid uuid.UUID,
) (ok bool, err error) {
    if err = ls.Repo.Delete(uid); err != nil {
        ok = false
    }
    return
}

func (ls *LinkService) Update(
    uid uuid.UUID,
    url string,
    name string,
    contentDescription string,
    useMarkdown bool,
) (ok bool, err error) {
    var description = domain.Description{contentDescription, useMarkdown}
    var updatedAt time.Time = time.Now()
    var link = domain.Link{uid, url, name, description, time.Time{}, updatedAt}
    if err = ls.Repo.Update(uid, link); err != nil {
        ok = false
    }
    return
}
