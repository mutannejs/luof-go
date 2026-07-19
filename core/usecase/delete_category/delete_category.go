package delete_category

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
)

type DeleteCategory struct {
	BelongsToRepo repository.BelongsTo
	CategoryRepo repository.Category
}

func New(btRepo repository.BelongsTo, cRepo repository.Category) DeleteCategory {
	return DeleteCategory{btRepo, cRepo}
}

func (dcUseCase *DeleteCategory) Execute(
	uid uuid.UUID,
) (err error) {
	exists, err := dcUseCase.CategoryRepo.Exists(uid)

	if err != nil {
		return
	} else if !exists {
		return lerror.GetNotFound(domain.CATEGORY_NOT_EXISTS)
	}

	hasLinks, err := dcUseCase.BelongsToRepo.HasLinks(uid)

	if err != nil {
		return
	} else if hasLinks {
		return lerror.GetConflict(domain.HAS_LINKS)
	}

	hasSubcategories, err := dcUseCase.CategoryRepo.HasSubcategories(uid)

	if err != nil {
		return
	} else if hasSubcategories {
		return lerror.GetConflict(domain.HAS_SUBCATEGORIES)
	}

	return dcUseCase.CategoryRepo.Delete(uid)
}
