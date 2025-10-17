package in

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/mutannejs/luof-go/core/out"

    "github.com/google/uuid"
)

type CategorySearchService struct {
    Repo out.CategoryRepository
}

func NewCategorySearch(repo out.CategoryRepository) *Category {
    return &Category{repo}
}

func (ls *Category) GetById(uid uuid.UUID) (link domain.Category, err error) {
    link, err = ls.Repo.GetById(uid)
    return 
}
