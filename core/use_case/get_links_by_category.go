package use_case

import (
    "github.com/mutannejs/luof-go/core/domain"
    "github.com/mutannejs/luof-go/core/repository"
    "github.com/google/uuid"
)

type GetLinksByCategoryUseCase struct {
    BelongsToRepo repository.BelongsTo
    CategoryRepo repository.Category
}

func GetLinksByCategory(btRepo repository.BelongsTo, cRepo repository.Category) GetLinksByCategoryUseCase {
    return GetLinksByCategoryUseCase{btRepo, cRepo}
}

func (glbcUseCase *GetLinksByCategoryUseCase) Execute(
    uid uuid.UUID,
) (links []domain.Link, err error) {
    var exists bool
    
    exists, err = glbcUseCase.CategoryRepo.Exists(uid)

    if exists {
        links, err = glbcUseCase.BelongsToRepo.GetLinksByCategory(uid)
    }

    return
}
