package delete_category

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/google/uuid"
)

type DeleteCategory struct {
	Repo repository.Category
}

func New(repo repository.Category) DeleteCategory {
	return DeleteCategory{repo}
}

func (dcUseCase *DeleteCategory) Execute(
	uid uuid.UUID,
) (exists bool, err error) {
	exists, err = dcUseCase.Repo.Exists(uid)

	if err != nil {
		return
	}

	if !exists {
		err = domain.CATEGORY_NOT_EXISTS
	} else {
		err = dcUseCase.Repo.Delete(uid)
	}

	return
}
