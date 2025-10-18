package use_case

import (
    "github.com/mutannejs/luof-go/domain"
    "github.com/mutannejs/luof-go/repository"
    "github.com/google/uuid"
)

type CreateLinkUseCase struct {
    Repo repository.Link
}

func CreateLink(repo repository.Link) CreateLinkUseCase {
    return CreateLinkUseCase{repo}
}

func (clUseCase *CreateLinkUseCase) Execute(
    url string,
    name string,
    description string,
    useMarkdown bool,
) (uid uuid.UUID, err error) {
    var link domain.Link

    link, err = domain.NewLink(url, name, description, useMarkdown)

    if link != nil {
        err = clUseCase.Repo.Create(link)
        uid = link.Uid
    }

    return
}
