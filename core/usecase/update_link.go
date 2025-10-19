package usecase

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/mutannejs/luof-go/core/repository"
    "github.com/google/uuid"
)

type UpdateLink struct {
    Repo repository.Link
}

func NewUpdateLink(repo repository.Link) UpdateLink {
    return UpdateLink{repo}
}

func (ulUseCase *UpdateLink) Execute(
    uid uuid.UUID,
    url string,
    name string,
    description string,
    useMarkdown bool,
) (exists bool, err error) {
    exists, err = ulUseCase.Repo.Exists(uid)

    if !exists || err != nil {
        return
    }

    var link domain.Link
    link, err = domain.NewLink(url, name, description, useMarkdown)

    if err == nil {
        err = ulUseCase.Repo.Update(uid, link)
    }

    return
}
