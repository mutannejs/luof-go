package insert_subcategory

import (
	"time"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/google/uuid"
)

type InsertSubcategory struct {
	Repo repository.Subcategory
}

func New(repo repository.Subcategory) InsertSubcategory {
	return InsertSubcategory{repo}
}

func (isUseCase *InsertSubcategory) Execute(
	fatherUid uuid.UUID,
	childUid uuid.UUID,
) (err error) {
	var areRelatives bool

	areRelatives, err = isUseCase.Repo.AreRelatives(fatherUid, childUid)

	if err != nil {
		return
	} else if areRelatives {
		return domain.ARE_RELATIVES
	}

	err = isUseCase.Repo.Create(fatherUid, childUid, time.Now())

	return
}
