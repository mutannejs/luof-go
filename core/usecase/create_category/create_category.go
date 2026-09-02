package create_category

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
)

type CreateCategory struct {
	Repo repository.Category
}

func New(repo repository.Category) CreateCategory {
	return CreateCategory{repo}
}

func (ccUseCase *CreateCategory) Execute(
	name string,
	description string,
	useMarkdown bool,
) (uid uuid.UUID, vErr lerror.ValueError) {
	var category domain.Category

	category, vErr = domain.NewCategory(name, description, useMarkdown)

	if vErr.IsNil() {
		vErr = ccUseCase.Repo.Create(category)
		uid = category.GetUid()
	}

	return
}
