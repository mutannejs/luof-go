package toggle_main_category

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/google/uuid"
)

type ToggleMainCategory struct {
    Repo repository.BelongsTo
}

func New(repo repository.BelongsTo) ToggleMainCategory {
    return ToggleMainCategory{repo}
}

func (tmcUseCase *ToggleMainCategory) Execute(
    linkUid uuid.UUID,
    categoryUid uuid.UUID,
    isMain bool,
) (err error) {
    var exists bool

    exists, err = tmcUseCase.Repo.Exists(linkUid, categoryUid)

    if err != nil {
        return
    } else if !exists {
        return domain.NOT_BELONGS
    }

    err = tmcUseCase.Repo.Update(linkUid, categoryUid, isMain)

    return
}
