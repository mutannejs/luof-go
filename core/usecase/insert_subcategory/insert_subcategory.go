package insert_subcategory

import (
	"time"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"

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
) (vErr lerror.ValueError) {
	var isAncestor, isSubcategory bool

	if fatherUid == childUid {
		return lerror.GetConflict(domain.CANNOT_BE_A_SUBCATEGORY_OF_ITSELF)
	}

	isSubcategory, vErr = isUseCase.Repo.IsSubcategory(fatherUid, childUid)

	if !vErr.IsNil() {
		return
	} else if isSubcategory {
		return lerror.GetConflict(domain.IS_SUBCATEGORY)
	}

	isAncestor, vErr = isUseCase.Repo.IsAncestor(childUid, fatherUid)

	if !vErr.IsNil() {
		return
	} else if isAncestor {
		return lerror.GetConflict(domain.ANCESTOR_NOT_BECOME_A_SUBCATEGORY)
	}

	var exists bool

	exists, vErr = isUseCase.Repo.Exists(fatherUid)

	if !vErr.IsNil() {
		return
	} else if !exists {
		return lerror.GetNotFound(domain.FATHER_NOT_EXISTS)
	}

	exists, vErr = isUseCase.Repo.Exists(childUid)

	if !vErr.IsNil() {
		return
	} else if !exists {
		return lerror.GetNotFound(domain.CHILD_NOT_EXISTS)
	}

	vErr = isUseCase.Repo.InsertSubcategory(fatherUid, childUid, time.Now())

	return
}
