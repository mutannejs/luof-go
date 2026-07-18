package get_all_root_categories

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
)

type GetAllRootCategories struct {
	Repo repository.Category
}

func New(repo repository.Category) GetAllRootCategories {
	return GetAllRootCategories{repo}
}

func (garcUseCase *GetAllRootCategories) Execute() (
	categories []domain.Category,
	err error,
) {
	categories, err = garcUseCase.Repo.GetAllRootCategories()
	return
}
