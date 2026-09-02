package get_all_root_categories

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"
)

type GetAllRootCategories struct {
	Repo repository.Category
}

func New(repo repository.Category) GetAllRootCategories {
	return GetAllRootCategories{repo}
}

func (garcUseCase *GetAllRootCategories) Execute() (
	categories []domain.Category,
	vErr lerror.ValueError,
) {
	categories, vErr = garcUseCase.Repo.GetAllRootCategories()
	return
}
