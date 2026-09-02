package insert_link_in_category

import (
	"time"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"

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
) (vErr lerror.ValueError) {
	var exists bool

	exists, vErr = ilicUseCase.BelongsToRepo.Exists(linkUid, categoryUid)

	if !vErr.IsNil() {
		return
	} else if exists {
		return lerror.GetConflict(domain.ALREADY_BELONGS)
	}

	exists, vErr = ilicUseCase.CategoryRepo.Exists(categoryUid)

	if !vErr.IsNil() {
		return
	} else if !exists {
		return lerror.GetNotFound(domain.CATEGORY_NOT_EXISTS)
	}

	exists, vErr = ilicUseCase.LinkRepo.Exists(linkUid)

	if !vErr.IsNil() {
		return
	} else if !exists {
		return lerror.GetNotFound(domain.LINK_NOT_EXISTS)
	}

	vErr = ilicUseCase.BelongsToRepo.SetHasNoMainCategory(linkUid)
	vErr = ilicUseCase.BelongsToRepo.Create(linkUid, categoryUid, time.Now(), isMain)

	return
}
