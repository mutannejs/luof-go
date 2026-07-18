package insert_link_in_category

import (
	"time"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/google/uuid"
)

type InsertLinkInCategory struct {
	BelongsToRepo repository.BelongsTo
	CategoryRepo repository.Category
	LinkRepo repository.Link
}

func New(btRepo repository.BelongsTo, cRepo repository.Category, lRepo repository.Link) InsertLinkInCategory {
	return InsertLinkInCategory{btRepo, cRepo, lRepo}
}

func (ilicUseCase *InsertLinkInCategory) Execute(
	linkUid uuid.UUID,
	categoryUid uuid.UUID,
	isMain bool,
) (err error) {
	var exists bool

	exists, err = ilicUseCase.BelongsToRepo.Exists(linkUid, categoryUid)

	if err != nil {
		return
	} else if exists {
		return domain.ALREADY_BELONGS
	}

	exists, err = ilicUseCase.CategoryRepo.Exists(categoryUid)

	if err != nil {
		return
	} else if !exists {
		return domain.CATEGORY_NOT_EXISTS
	}

	exists, err = ilicUseCase.LinkRepo.Exists(linkUid)

	if err != nil {
		return
	} else if !exists {
		return domain.LINK_NOT_EXISTS
	}

	err = ilicUseCase.BelongsToRepo.SetHasNoMainCategory(linkUid)
	err = ilicUseCase.BelongsToRepo.Create(linkUid, categoryUid, time.Now(), isMain)

	return
}
