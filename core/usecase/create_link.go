package usecase

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/mutannejs/luof-go/core/repository"

    "github.com/google/uuid"
)

type CreateLink struct {
    Repo repository.Link
}

func NewCreateLink(repo repository.Link) CreateLink {
    return CreateLink{repo}
}

func (clUseCase *CreateLink) Execute(
    url string,
    name string,
    description string,
    useMarkdown bool,
) (uid uuid.UUID, err error) {
    var link domain.Link

    link, err = domain.NewLink(url, name, description, useMarkdown)

    if err == nil {
        err = clUseCase.Repo.Create(link)
        uid = link.GetUid()
    }

    return
}
