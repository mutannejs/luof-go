package insert_link_in_category

import (
	"time"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/google/uuid"
)

type InsertLinkInCategory struct {
	Repo repository.BelongsTo
}

func New(repo repository.BelongsTo) InsertLinkInCategory {
	return InsertLinkInCategory{repo}
}

func (ilicUseCase *InsertLinkInCategory) Execute(
	linkUid uuid.UUID,
	categoryUid uuid.UUID,
	isMain bool,
) (err error) {
	var exists bool

	exists, err = ilicUseCase.Repo.Exists(linkUid, categoryUid)

	if err != nil {
		return
	} else if exists {
		return domain.ALREADY_BELONGS
	}

	err = ilicUseCase.Repo.SetHasNoMainCategory(linkUid)

	err = ilicUseCase.Repo.Create(linkUid, categoryUid, time.Now(), isMain)

	return
}
