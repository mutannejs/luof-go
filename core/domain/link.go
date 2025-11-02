package domain

import (
    "errors"
    "time"

    "github.com/mutannejs/luof-go/pkg/luuid"

    "github.com/google/uuid"
)

var (
    LINK_ERROR_NEW = errors.New("error instantiate new link")
)

type Link struct {
    uid uuid.UUID
    Url string
    Name string
    Description Description
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (l Link) GetUid() uuid.UUID {
    return l.uid
}

func NewLink(
    url string,
    name string,
    contentDescription string,
    useMarkdown bool,
) (Link, error) {
    var uid uuid.UUID
    var err error

    uid, err = luuid.New()
    if err != nil {
        return Link{}, errors.Join(LINK_ERROR_NEW, err)
    }

    var createdAt time.Time = time.Now()
    var updatedAt time.Time
    var description = Description{contentDescription, useMarkdown}
    var link = Link{uid, url, name, description, createdAt, updatedAt}

    return link, nil
}
