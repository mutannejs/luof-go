package use_case

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/mutannejs/luof-go/core/repository"
    "github.com/google/uuid"
)

type CreateCategoryUseCase struct {
    Repo repository.Category
}

func CreateCategory(repo repository.Category) CreateCategoryUseCase {
    return CreateCategoryUseCase{repo}
}

func (ccUseCase *CreateCategoryUseCase) Execute(
    name string,
    description string,
    useMarkdown bool,
) (uid uuid.UUID, err error) {
    var category domain.Category

    category, err = domain.NewCategory(name, description, useMarkdown)

    if err == nil {
        err = ccUseCase.Repo.Create(category)
        uid = category.Uid
    }

    return
}
