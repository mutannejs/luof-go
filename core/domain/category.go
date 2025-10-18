package domain

import (
    "github.com/google/uuid"
    "time"
)

type Category struct {
    Uid uuid.UUID
    Name string
    Description Description
    CreatedAt time.Time
    UpdatedAt time.Time
}

func NewCategory(
    name string,
    contentDescription string,
    useMarkdown bool,
) (Category, error) {
    if uid, err = uuid.New(); err != nil {
        return nil, err
    }

    var createdAt time.Time = time.Now()
    var updatedAt time.Time
    var description = domain.Description{contentDescription, useMarkdown}
    var category = domain.Category{uid, name, description, createdAt, updatedAt}

    return category, nil
}
