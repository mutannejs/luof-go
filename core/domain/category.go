package domain

import (
	"errors"
	"time"

	"github.com/mutannejs/luof-go/pkg/luuid"

	"github.com/google/uuid"
)

var (
    CATEGORY_ERROR_NEW = errors.New("error instantiate new category")
    CATEGORY_NOT_EXISTS = errors.New("the searched category does not exist")
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

func (l *Category) SetUid(uid uuid.UUID) {
    l.uid = uid
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
        return Category{}, errors.Join(CATEGORY_ERROR_NEW, err)
    }

    var createdAt time.Time = time.Now()
    var updatedAt time.Time
    var description = Description{contentDescription, useMarkdown}
    var category = Category{uid, name, description, createdAt, updatedAt}

    return category, nil
}
