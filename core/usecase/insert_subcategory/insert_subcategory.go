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
	var isAncestor bool

	isAncestor, err = isUseCase.Repo.IsAncestor(fatherUid, childUid)

	if err != nil {
		return
	} else if isAncestor {
		return domain.ANCESTOR_NOT_BECOME_A_SUBCATEGORY
	}

	err = isUseCase.Repo.InsertSubcategory(fatherUid, childUid, time.Now())

	return
}
