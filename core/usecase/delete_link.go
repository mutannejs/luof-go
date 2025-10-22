package usecase

import (
    "github.com/mutannejs/luof-go/core/repository"

    "github.com/google/uuid"
)

type DeleteLink struct {
    Repo repository.Link
}

func NewDeleteLink(repo repository.Link) DeleteLink {
    return DeleteLink{repo}
}

func (dlUseCase *DeleteLink) Execute(
    uid uuid.UUID,
) (exists bool, err error) {
    exists, err = dlUseCase.Repo.Exists(uid)

    if exists {
        err = dlUseCase.Repo.Delete(uid)
    }

    return
}
