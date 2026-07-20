package toggle_main_category

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
)

type ToggleMainCategory struct {
	BelongsToRepo repository.BelongsTo
	CategoryRepo repository.Category
	LinkRepo repository.Link
}

func New(btRepo repository.BelongsTo, cRepo repository.Category, lRepo repository.Link) ToggleMainCategory {
	return ToggleMainCategory{btRepo, cRepo, lRepo}
}

func (tmcUseCase *ToggleMainCategory) Execute(
	linkUid uuid.UUID,
	categoryUid uuid.UUID,
	isMain bool,
) (err error) {
	var exists bool

	exists, err = tmcUseCase.CategoryRepo.Exists(categoryUid)

	if err != nil {
		return
	} else if !exists {
		return lerror.GetNotFound(domain.CATEGORY_NOT_EXISTS)
	}

	exists, err = tmcUseCase.LinkRepo.Exists(linkUid)

	if err != nil {
		return
	} else if !exists {
		return lerror.GetNotFound(domain.LINK_NOT_EXISTS)
	}

	exists, err = tmcUseCase.BelongsToRepo.Exists(linkUid, categoryUid)

	if err != nil {
		return
	} else if !exists {
		return lerror.GetNotFound(domain.NOT_BELONGS)
	}

	err = tmcUseCase.BelongsToRepo.Update(linkUid, categoryUid, isMain)

	return
}
