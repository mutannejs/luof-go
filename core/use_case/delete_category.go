package use_case

import (
    "github.com/mutannejs/luof-go/core/repository"
    "github.com/google/uuid"
)

type DeleteCategoryUseCase struct {
    Repo repository.Category
}

func DeleteCategory(repo repository.Category) DeleteCategoryUseCase {
    return DeleteCategoryUseCase{repo}
}

func (dcUseCase *DeleteCategoryUseCase) Execute(
    uid uuid.UUID,
) (exists bool, err error) {
    exists, err = dcUseCase.Repo.Exists(uid)

    if exists {
        err = dcUseCase.Repo.Delete(uid)
    }

    return
}
