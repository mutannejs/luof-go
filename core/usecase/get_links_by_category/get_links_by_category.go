package get_links_by_category

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
)

type GetLinksByCategory struct {
	BelongsToRepo repository.BelongsTo
	CategoryRepo repository.Category
}

func New(btRepo repository.BelongsTo, cRepo repository.Category) GetLinksByCategory {
	return GetLinksByCategory{btRepo, cRepo}
}

func (glbcUseCase *GetLinksByCategory) Execute(
	uid uuid.UUID,
) (links []domain.Link, vErr lerror.ValueError) {
	var exists bool
	
	exists, vErr = glbcUseCase.CategoryRepo.Exists(uid)

	if !vErr.IsNil() {
		return
	}

	if !exists {
		vErr = lerror.GetNotFound(domain.CATEGORY_NOT_EXISTS)
	} else {
		links, vErr = glbcUseCase.BelongsToRepo.GetLinksByCategory(uid)
	}

	return
}
