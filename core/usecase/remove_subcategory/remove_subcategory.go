package remove_subcategory

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/google/uuid"
)

type RemoveSubcategory struct {
	Repo repository.Category
}

func New(repo repository.Category) RemoveSubcategory {
	return RemoveSubcategory{repo}
}

func (isUseCase *RemoveSubcategory) Execute(
	fatherUid uuid.UUID,
	childUid uuid.UUID,
) (err error) {
	var isSubcategory bool

	isSubcategory, err = isUseCase.Repo.IsSubcategory(fatherUid, childUid)

	if err != nil {
		return
	} else if !isSubcategory {
		return domain.NOT_IS_SUBCATEGORY
	}

	err = isUseCase.Repo.DeleteSubcategory(childUid)

	return
}
