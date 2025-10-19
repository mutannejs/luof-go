package usecase

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/mutannejs/luof-go/core/repository"
    "github.com/google/uuid"
)

type GetLinkByUid struct {
    Repo repository.Link
}

func NewGetLinkByUid(repo repository.Link) GetLinkByUid {
    return GetLinkByUid{repo}
}

func (glbuUseCase *GetLinkByUid) Execute(
    uid uuid.UUID,
) (link domain.Link, err error) {
    var exists bool

    exists, err = glbuUseCase.Repo.Exists(uid)

    if exists {
        link, err = glbuUseCase.Repo.GetByUid(uid)
    }

    return
}
