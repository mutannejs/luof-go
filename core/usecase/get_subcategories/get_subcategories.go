package get_subcategories

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"

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
) (subcategories []domain.Category, vErr lerror.ValueError) {
	var exists bool
	
	exists, vErr = gsUseCase.CategoryRepo.Exists(uid)

	if !vErr.IsNil() {
		return
	}

	if !exists {
		vErr = lerror.GetNotFound(domain.CATEGORY_NOT_EXISTS)
	} else {
		subcategories, vErr = gsUseCase.CategoryRepo.GetSubcategories(uid)
	}

	return
}
