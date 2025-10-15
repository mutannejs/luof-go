package in

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/mutannejs/luof-go/core/ports/out"
    "github.com/mutannejs/luof-go/pkg/luuid"

    "github.com/google/uuid"
    "time"
)

type CategoryService struct {
    Repo out.CategoryRepository
}

func NewCategoryService(repo out.CategoryRepository) *CategoryService {
    return &CategoryService{repo}
}

func (cs *CategoryService) Create(
    name string,
    contentDescription string,
    useMarkdown bool,
) (uid uuid.UUID, err error) {
    if uid, err = luuid.New(); err != nil {
        return
    }
    var description = domain.Description{contentDescription, useMarkdown}
    var createdAt time.Time = time.Now()
    var updatedAt time.Time
    var category = domain.Category{uid, name, description, createdAt, updatedAt}
    err = cs.Repo.Create(category)
    return
}

func (cs *CategoryService) Delete(
    uid uuid.UUID,
) (ok bool, err error) {
    if ok, err = cs.Repo.Exists(uid); err != nil || !ok {
        return
    }
    if err = cs.Repo.Delete(uid); err != nil {
        ok = false
    }
    return
}

func (cs *CategoryService) Update(
    uid uuid.UUID,
    name string,
    contentDescription string,
    useMarkdown bool,
) (ok bool, err error) {
    if ok, err = cs.Repo.Exists(uid); err != nil || !ok {
        return
    }
    var description = domain.Description{contentDescription, useMarkdown}
    var updatedAt time.Time = time.Now()
    var category = domain.Category{uid, name, description, time.Time{}, updatedAt}
    if err = cs.Repo.Update(uid, category); err != nil {
        ok = false
    }
    return
}
