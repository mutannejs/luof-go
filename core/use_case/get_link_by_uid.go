package use_case

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/mutannejs/luof-go/core/repository"
    "github.com/google/uuid"
)

type GetLinkByUidUseCase struct {
    Repo repository.Link
}

func GetLinkByUid(repo repository.Link) GetLinkByUidUseCase {
    return GetLinkByUidUseCase{repo}
}

func (glbuUseCase *GetLinkByUidUseCase) Execute(
    uid uuid.UUID,
) (link domain.Link, err error) {
    var exists bool

    exists, err = glbuUseCase.Repo.Exists(uid)

    if exists {
        link, err = glbuUseCase.Repo.GetByUid(uid)
    }

    return
}
