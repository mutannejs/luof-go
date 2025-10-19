package domain

import (
    "github.com/mutannejs/luof-go/pkg/luuid"
    "github.com/google/uuid"
    "time"
)

type Category struct {
    uid uuid.UUID
    Name string
    Description Description
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (c Category) GetUid() uuid.UUID {
    return c.uid
}

func NewCategory(
    name string,
    contentDescription string,
    useMarkdown bool,
) (Category, error) {
    var uid uuid.UUID
    var err error

    uid, err = luuid.New()
    if err != nil {
        return Category{}, err
    }

    var createdAt time.Time = time.Now()
    var updatedAt time.Time
    var description = Description{contentDescription, useMarkdown}
    var category = Category{uid, name, description, createdAt, updatedAt}

    return category, nil
}
