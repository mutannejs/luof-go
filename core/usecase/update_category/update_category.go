package update_category

import (
	"time"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/google/uuid"
)

type UpdateCategory struct {
	Repo repository.Category
}

func New(repo repository.Category) UpdateCategory {
	return UpdateCategory{repo}
}

func (ucUseCase *UpdateCategory) Execute(
	uid uuid.UUID,
	name string,
	description string,
	useMarkdown bool,
) (exists bool, err error) {
	exists, err = ucUseCase.Repo.Exists(uid)

	if err != nil {
		return
	} else if !exists {
		return exists, domain.CATEGORY_NOT_EXISTS
	}

	var category domain.Category
	category, err = domain.NewCategory(name, description, useMarkdown)

	if err == nil {
		category.UpdatedAt = time.Now()
		err = ucUseCase.Repo.Update(uid, category)
	}

	return
}
