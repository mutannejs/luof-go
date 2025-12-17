package get_links_by_category

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

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
) (links []domain.Link, err error) {
    var exists bool
    
    exists, err = glbcUseCase.CategoryRepo.Exists(uid)

    if err != nil {
        return
    }

    if !exists {
        err = domain.CATEGORY_NOT_EXISTS
    } else {
        links, err = glbcUseCase.BelongsToRepo.GetLinksByCategory(uid)
    }

    return
}
