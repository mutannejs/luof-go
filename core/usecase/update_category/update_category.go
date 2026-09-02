package update_category

import (
	"time"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"

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
) (exists bool, vErr lerror.ValueError) {
	exists, vErr = ucUseCase.Repo.Exists(uid)

	if !vErr.IsNil() {
		return
	} else if !exists {
		return exists, lerror.GetNotFound(domain.CATEGORY_NOT_EXISTS)
	}

	var category domain.Category
	category, vErr = domain.NewCategory(name, description, useMarkdown)

	if vErr.IsNil() {
		category.UpdatedAt = time.Now()
		vErr = ucUseCase.Repo.Update(uid, category)
	}

	return
}
