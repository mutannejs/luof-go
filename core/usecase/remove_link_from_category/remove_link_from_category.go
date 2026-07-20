package remove_link_from_category

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
)

type RemoveLinkFromCategory struct {
	BelongsToRepo repository.BelongsTo
	CategoryRepo repository.Category
	LinkRepo repository.Link
}

func New(btRepo repository.BelongsTo, cRepo repository.Category, lRepo repository.Link) RemoveLinkFromCategory {
	return RemoveLinkFromCategory{btRepo, cRepo, lRepo}
}

func (rlfcUseCase *RemoveLinkFromCategory) Execute(
	linkUid uuid.UUID,
	categoryUid uuid.UUID,
) (err error) {
	var exists bool

	exists, err = rlfcUseCase.CategoryRepo.Exists(categoryUid)

	if err != nil {
		return
	} else if !exists {
		return lerror.GetNotFound(domain.CATEGORY_NOT_EXISTS)
	}

	exists, err = rlfcUseCase.LinkRepo.Exists(linkUid)

	if err != nil {
		return
	} else if !exists {
		return lerror.GetNotFound(domain.LINK_NOT_EXISTS)
	}

	exists, err = rlfcUseCase.BelongsToRepo.Exists(linkUid, categoryUid)

	if err != nil {
		return
	} else if !exists {
		return lerror.GetNotFound(domain.NOT_BELONGS)
	}

	err = rlfcUseCase.BelongsToRepo.Delete(linkUid, categoryUid)

	return
}
