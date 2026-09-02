package get_category_by_uid

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
)

type GetCategoryByUid struct {
	Repo repository.Category
}

func New(repo repository.Category) GetCategoryByUid {
	return GetCategoryByUid{repo}
}

func (gcbuUseCase *GetCategoryByUid) Execute(
	uid uuid.UUID,
) (category domain.Category, vErr lerror.ValueError) {
	var exists bool

	exists, vErr = gcbuUseCase.Repo.Exists(uid)

	if !vErr.IsNil() {
		return
	}

	if !exists {
		vErr = lerror.GetNotFound(domain.CATEGORY_NOT_EXISTS)
	} else {
		category, vErr = gcbuUseCase.Repo.GetByUid(uid)
	}

	return
}
