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
) (exists bool, vErr lerror.ValueError) {
	exists, vErr = dlUseCase.Repo.Exists(uid)

	if !vErr.IsNil() {
		return
	}

	if !exists {
		vErr = lerror.GetNotFound(domain.LINK_NOT_EXISTS)
	} else {
		vErr = dlUseCase.Repo.Delete(uid)
	}

	return
}
