package usecase

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/mutannejs/luof-go/core/repository"
    "github.com/google/uuid"
)

type GetLinksByCategory struct {
    BelongsToRepo repository.BelongsTo
    CategoryRepo repository.Category
}

func NewGetLinksByCategory(btRepo repository.BelongsTo, cRepo repository.Category) GetLinksByCategory {
    return GetLinksByCategory{btRepo, cRepo}
}

func (glbcUseCase *GetLinksByCategory) Execute(
    uid uuid.UUID,
) (links []domain.Link, err error) {
    var exists bool
    
    exists, err = glbcUseCase.CategoryRepo.Exists(uid)

    if exists {
        links, err = glbcUseCase.BelongsToRepo.GetLinksByCategory(uid)
    }

    return
}
