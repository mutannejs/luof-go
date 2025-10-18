package use_case

import (
    "github.com/mutannejs/luof-go/domain"
    "github.com/mutannejs/luof-go/repository"
    "github.com/google/uuid"
)

type GetLinksByCategoryUseCase struct {
    BelongsToRepo repository.BelongsTo
    CategoryRepo repository.Category
}

func GetLinkByUid(repo repository.BelongsTo, repo repository.Category) GetLinksByCategoryUseCase {
    return GetLinksByCategoryUseCase{repo}
}

func (glbcUseCase *GetLinksByCategoryUseCase) Execute(
    uid uuid.UUID
) (links []domain.Link, err error) {
    exists = glbcUseCase.CategoryRepo.Exists(uid)

    if exists {
        links, err = glbcUseCase.BelongsTo.GetLinksByCategory(uid)
    }

    return
}
