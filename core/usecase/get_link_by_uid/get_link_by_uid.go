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
) (link domain.Link, vErr lerror.ValueError) {
	var exists bool

	exists, vErr = glbuUseCase.Repo.Exists(uid)

	if !vErr.IsNil() {
		return
	}

	if !exists {
		vErr = lerror.GetNotFound(domain.LINK_NOT_EXISTS)
	} else {
		link, vErr = glbuUseCase.Repo.GetByUid(uid)
	}

	return
}
