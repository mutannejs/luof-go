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
	var areRelated bool

	areRelated, err = isUseCase.Repo.AreRelated(fatherUid, childUid)

	if err != nil {
		return
	} else if areRelated {
		return domain.ARE_RELATED
	}

	err = isUseCase.Repo.InsertSubcategory(fatherUid, childUid, time.Now())

	return
}
