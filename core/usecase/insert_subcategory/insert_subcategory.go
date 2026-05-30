package insert_subcategory

import (
	"time"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/google/uuid"
)

type InsertSubcategory struct {
	Repo repository.Category
}

func New(repo repository.Category) InsertSubcategory {
	return InsertSubcategory{repo}
}

func (isUseCase *InsertSubcategory) Execute(
	fatherUid uuid.UUID,
	childUid uuid.UUID,
) (err error) {
	var isAncestor, isSubcategory bool

	if fatherUid == childUid {
		return domain.CANNOT_BE_A_SUBCATEGORY_OF_ITSELF
	}

	isSubcategory, err = isUseCase.Repo.IsSubcategory(fatherUid, childUid)

	if err != nil {
		return
	} else if isSubcategory {
		return domain.IS_SUBCATEGORY
	}

	isAncestor, err = isUseCase.Repo.IsAncestor(childUid, fatherUid)

	if err != nil {
		return
	} else if isAncestor {
		return domain.ANCESTOR_NOT_BECOME_A_SUBCATEGORY
	}

	var exists bool

	exists, err = isUseCase.Repo.Exists(fatherUid)

	if err != nil {
		return
	} else if !exists {
		return domain.FATHER_NOT_EXISTS
	}

	exists, err = isUseCase.Repo.Exists(childUid)

	if err != nil {
		return
	} else if !exists {
		return domain.CHILD_NOT_EXISTS
	}

	err = isUseCase.Repo.InsertSubcategory(fatherUid, childUid, time.Now())

	return
}
