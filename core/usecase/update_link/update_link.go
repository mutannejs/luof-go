package update_link

import (
	"time"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
)

type UpdateLink struct {
	Repo repository.Link
}

func New(repo repository.Link) UpdateLink {
	return UpdateLink{repo}
}

func (ulUseCase *UpdateLink) Execute(
	uid uuid.UUID,
	url string,
	name string,
	description string,
	useMarkdown bool,
) (exists bool, vErr lerror.ValueError) {
	exists, vErr = ulUseCase.Repo.Exists(uid)

	if !vErr.IsNil() {
		return
	} else if !exists {
		return exists, lerror.GetNotFound(domain.LINK_NOT_EXISTS)
	}

	var link domain.Link
	link, vErr = domain.NewLink(url, name, description, useMarkdown)

	if vErr.IsNil() {
		link.UpdatedAt = time.Now()
		vErr = ulUseCase.Repo.Update(uid, link)
	}

	return
}
