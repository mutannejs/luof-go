package remove_link_from_category

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/google/uuid"
)

type RemoveLinkFromCategory struct {
	Repo repository.BelongsTo
}

func New(repo repository.BelongsTo) RemoveLinkFromCategory {
	return RemoveLinkFromCategory{repo}
}

func (rlfcUseCase *RemoveLinkFromCategory) Execute(
	linkUid uuid.UUID,
	categoryUid uuid.UUID,
) (err error) {
	var exists bool

	exists, err = rlfcUseCase.Repo.Exists(linkUid, categoryUid)

	if err != nil {
		return
	} else if !exists {
		return domain.NOT_BELONGS
	}

	err = rlfcUseCase.Repo.Delete(linkUid, categoryUid)

	return
}
