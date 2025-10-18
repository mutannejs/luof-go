package use_case

import (
    "github.com/mutannejs/luof-go/domain"
    "github.com/mutannejs/luof-go/repository"
    "github.com/google/uuid"
)

type UpdateLinkUseCase struct {
    Repo repository.Link
}

func UpdateLink(repo repository.Link) UpdateLinkUseCase {
    return UpdateLinkUseCase{repo}
}

func (ulUseCase *UpdateLinkUseCase) Execute(
    url string,
    name string,
    description string,
    useMarkdown bool,
) (exists bool, err error) {
    exists = ulUseCase.Repo.Exists(uid)

    if !exists {
        return
    }

    var link domain.Link
    link, err = domain.NewLink(url, name, description, useMarkdown)

    if link != nil {
        err = ulUseCase.Repo.Update(uid, link)
    }

    return
}
