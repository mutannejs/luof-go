package get_subcategories

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"

	"github.com/google/uuid"
)

type GetSubcategories struct {
	CategoryRepo repository.Category
}

func New(cRepo repository.Category) GetSubcategories {
	return GetSubcategories{cRepo}
}

func (gsUseCase *GetSubcategories) Execute(
	uid uuid.UUID,
) (subcategories []domain.Category, err error) {
	var exists bool
	
	exists, err = gsUseCase.CategoryRepo.Exists(uid)

	if err != nil {
		return
	}

	if !exists {
		err = domain.CATEGORY_NOT_EXISTS
	} else {
		subcategories, err = gsUseCase.CategoryRepo.GetSubcategories(uid)
	}

	return
}
