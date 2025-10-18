package use_case

import (
    "github.com/mutannejs/luof-go/domain"
    "github.com/mutannejs/luof-go/repository"
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

    if category != nil {
        err = ccUseCase.Repo.Create(category)
        uid = category.Uid
    }

    return
}
