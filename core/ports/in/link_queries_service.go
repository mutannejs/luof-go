package in

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/mutannejs/luof-go/core/out"

    "github.com/google/uuid"
)

type LinkSearch struct {
    Repo out.LinkRepository
}

func NewLinkSearch(repo out.LinkRepository) *LinkSearch {
    return &LinkSearch{repo}
}

func (ls *LinkSearch) GetByCategory(categoryUid uuid.UUID) (links []domain.Link, err error) {
    links, err = ls.Repo.GetByCategory(categoryUid)
    return 
}

func (ls *LinkSearch) GetById(uid uuid.UUID) (link domain.Link, err error) {
    link, err = ls.Repo.GetById(uid)
    return 
}
