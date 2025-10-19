package use_case

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/mutannejs/luof-go/core/repository"
    "github.com/google/uuid"
)

type UpdateCategoryUseCase struct {
    Repo repository.Category
}

func UpdateCategory(repo repository.Category) UpdateCategoryUseCase {
    return UpdateCategoryUseCase{repo}
}

func (ucUseCase *UpdateCategoryUseCase) Execute(
    uid uuid.UUID,
    name string,
    description string,
    useMarkdown bool,
) (exists bool, err error) {
    exists, err = ucUseCase.Repo.Exists(uid)

    if !exists || err != nil {
        return
    }

    var category domain.Category
    category, err = domain.NewCategory(name, description, useMarkdown)

    if err == nil {
        err = ucUseCase.Repo.Update(uid, category)
    }

    return
}
