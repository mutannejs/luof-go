package domain

import (
    "github.com/google/uuid"
    "time"
)

type Link struct {
    Uid uuid.UUID
    Url string
    Name string
    Description Description
    CreatedAt time.Time
    UpdatedAt time.Time
}

func NewLink(
    url string,
    name string,
    contentDescription string,
    useMarkdown bool,
) (Link, error) {
    if uid, err = uuid.New(); err != nil {
        return nil, err
    }

    var createdAt time.Time = time.Now()
    var insertedAt time.Time = createdAt
    var updatedAt time.Time
    var description = domain.Description{contentDescription, useMarkdown}
    var link = domain.Link{uid, url, name, description, createdAt, updatedAt}

    return link, nil
}
