package delete_link

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
)

type DeleteLink struct {
	Repo repository.Link
}

func New(repo repository.Link) DeleteLink {
	return DeleteLink{repo}
}

func (dlUseCase *DeleteLink) Execute(
	uid uuid.UUID,
) (exists bool, err error) {
	exists, err = dlUseCase.Repo.Exists(uid)

	if err != nil {
		return
	}

	if !exists {
		err = lerror.GetNotFound(domain.LINK_NOT_EXISTS)
	} else {
		err = dlUseCase.Repo.Delete(uid)
	}

	return
}
