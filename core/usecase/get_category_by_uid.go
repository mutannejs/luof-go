package usecase

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/mutannejs/luof-go/core/repository"
    "github.com/google/uuid"
)

type GetCategoryByUid struct {
    Repo repository.Category
}

func NewGetCategoryByUid(repo repository.Category) GetCategoryByUid {
    return GetCategoryByUid{repo}
}

func (gcbuUseCase *GetCategoryByUid) Execute(
    uid uuid.UUID,
) (category domain.Category, err error) {
    var exists bool

    exists, err = gcbuUseCase.Repo.Exists(uid)

    if exists {
        category, err = gcbuUseCase.Repo.GetByUid(uid)
    }

    return
}
