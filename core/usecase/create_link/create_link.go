package create_link

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
)

type CreateLink struct {
	Repo repository.Link
}

func New(repo repository.Link) CreateLink {
	return CreateLink{repo}
}

func (clUseCase *CreateLink) Execute(
	url string,
	name string,
	description string,
	useMarkdown bool,
) (uid uuid.UUID, vErr lerror.ValueError) {
	var link domain.Link

	link, vErr = domain.NewLink(url, name, description, useMarkdown)

	if vErr.IsNil() {
		vErr = clUseCase.Repo.Create(link)
		uid = link.GetUid()
	}

	return
}
