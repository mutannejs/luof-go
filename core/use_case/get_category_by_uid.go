package use_case

import (
    "github.com/mutannejs/luof-go/domain"
    "github.com/mutannejs/luof-go/repository"
    "github.com/google/uuid"
)

type GetCategoryByUidUseCase struct {
    Repo repository.Category
}

func GetCategoryByUid(repo repository.Category) GetCategoryByUidUseCase {
    return GetCategoryByUidUseCase{repo}
}

func (gcbuUseCase *GetCategoryByUidUseCase) Execute(
    uid uuid.UUID
) (category domain.Category, err error) {
    exists = gcbuUseCase.Repo.Exists(uid)

    if exists {
        category, err = gcbuUseCase.Repo.GetByUid(uid)
    }

    return
}
