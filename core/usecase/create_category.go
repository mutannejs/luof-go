package usecase

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/mutannejs/luof-go/core/repository"

    "github.com/google/uuid"
)

type CreateCategory struct {
    Repo repository.Category
}

func NewCreateCategory(repo repository.Category) CreateCategory {
    return CreateCategory{repo}
}

func (ccUseCase *CreateCategory) Execute(
    name string,
    description string,
    useMarkdown bool,
) (uid uuid.UUID, err error) {
    var category domain.Category

    category, err = domain.NewCategory(name, description, useMarkdown)

    if err == nil {
        err = ccUseCase.Repo.Create(category)
        uid = category.GetUid()
    }

    return
}
