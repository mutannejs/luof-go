package usecase

import (
    "github.com/mutannejs/luof-go/core/repository"
    "github.com/google/uuid"
)

type DeleteCategory struct {
    Repo repository.Category
}

func NewDeleteCategory(repo repository.Category) DeleteCategory {
    return DeleteCategory{repo}
}

func (dcUseCase *DeleteCategory) Execute(
    uid uuid.UUID,
) (exists bool, err error) {
    exists, err = dcUseCase.Repo.Exists(uid)

    if exists {
        err = dcUseCase.Repo.Delete(uid)
    }

    return
}
