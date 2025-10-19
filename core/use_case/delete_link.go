package use_case

import (
    "github.com/mutannejs/luof-go/core/repository"
    "github.com/google/uuid"
)

type DeleteLinkUseCase struct {
    Repo repository.Link
}

func DeleteLink(repo repository.Link) DeleteLinkUseCase {
    return DeleteLinkUseCase{repo}
}

func (dlUseCase *DeleteLinkUseCase) Execute(
    uid uuid.UUID,
) (exists bool, err error) {
    exists, err = dlUseCase.Repo.Exists(uid)

    if exists {
        err = dlUseCase.Repo.Delete(uid)
    }

    return
}
