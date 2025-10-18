package use_case

import (
    "github.com/mutannejs/luof-go/domain"
    "github.com/mutannejs/luof-go/repository"
    "github.com/google/uuid"
)

type UpdateCategoryUseCase struct {
    Repo repository.Category
}

func UpdateCategory(repo repository.Category) UpdateCategoryUseCase {
    return UpdateCategoryUseCase{repo}
}

func (ucUseCase *UpdateCategoryUseCase) Execute(
    url string,
    name string,
    description string,
    useMarkdown bool,
) (exists bool, err error) {
    exists = ucUseCase.Repo.Exists(uid)

    if !exists {
        return
    }

    var category domain.Category
    category, err = domain.NewCategory(url, name, description, useMarkdown)

    if category != nil {
        err = ucUseCase.Repo.Update(uid, category)
    }

    return
}
