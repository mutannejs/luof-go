package get_link_by_uid

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
)

type GetLinkByUid struct {
	Repo repository.Link
}

func New(repo repository.Link) GetLinkByUid {
	return GetLinkByUid{repo}
}

func (glbuUseCase *GetLinkByUid) Execute(
	uid uuid.UUID,
) (link domain.Link, err error) {
	var exists bool

	exists, err = glbuUseCase.Repo.Exists(uid)

	if err != nil {
		return
	}

	if !exists {
		err = lerror.GetNotFound(domain.LINK_NOT_EXISTS)
	} else {
		link, err = glbuUseCase.Repo.GetByUid(uid)
	}

	return
}
