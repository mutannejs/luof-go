package domain

import (
	"time"

	"github.com/mutannejs/luof-go/pkg/lerror"
	"github.com/mutannejs/luof-go/pkg/luuid"

	"github.com/google/uuid"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	ANCESTOR_NOT_BECOME_A_SUBCATEGORY = &i18n.Message{
		ID: "ANCESTOR_NOT_BECOME_A_SUBCATEGORY",
		Description: "Retornado quando requisitado inserir de uma categoria em outra categoria que é sua descedente (direta ou indireta)",
		Other: "one ancestral category of another cannot become a subcategory of it",
	}
	CANNOT_BE_A_SUBCATEGORY_OF_ITSELF = &i18n.Message{
		ID: "CANNOT_BE_A_SUBCATEGORY_OF_ITSELF",
		Description: "Retornado quando requisitado inserir uma categoria nela mesma",
		Other: "a category cannot be a subcategory of itself",
	}
	CATEGORY_ERROR_NEW = &i18n.Message{
		ID: "CATEGORY_ERROR_NEW",
		Description: "Retornado quando uuid.New falha",
		Other: "error instantiate new category",
	}
	CATEGORY_NOT_EXISTS = &i18n.Message{
		ID: "CATEGORY_NOT_EXISTS",
		Description: "Retornado quando consultada uma categoria que não existe",
		Other: "the searched category does not exist",
	}
	CHILD_NOT_EXISTS = &i18n.Message{
		ID: "CHILD_NOT_EXISTS",
		Description: "Retornado quando requisitado inserir uma categoria que não existe em outra",
		Other: "the child category does not exist",
	}
	FATHER_NOT_EXISTS = &i18n.Message{
		ID: "FATHER_NOT_EXISTS",
		Description: "Retornado quando requisitado inserir uma categoria em uma categoria que não existe",
		Other: "the father category does not exist",
	}
	IS_SUBCATEGORY = &i18n.Message{
		ID: "IS_SUBCATEGORY",
		Description: "Retornado quando requisitado inserir uma categoria em uma categoria que não existe",
		Other: "the child is already subcategory of the father",
	}
	HAS_SUBCATEGORIES = &i18n.Message{
		ID: "HAS_SUBCATEGORIES",
		Description: "Retornado quando requisitado deletar uma categoria que possui subcategorias",
		Other: "the category cannot be deleted because it has subcategories",
	}
	NOT_IS_SUBCATEGORY = &i18n.Message{
		ID: "NOT_IS_SUBCATEGORY",
		Description: "Retornado quando requisitado remover uma categoria de outra que não é sua categoria pai",
		Other: "the child is not subcategory of the father",
	}
)

type Category struct {
	uid uuid.UUID
	Name string
	Description Description
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c Category) GetUid() uuid.UUID {
	return c.uid
}

func (l *Category) SetUid(uid uuid.UUID) {
	l.uid = uid
}

func NewCategory(
	name string,
	contentDescription string,
	useMarkdown bool,
) (category Category, vError lerror.ValueError) {
	var uid uuid.UUID
	var err error

	uid, err = luuid.New()
	if err != nil {
		vError = lerror.GetInternals(CATEGORY_ERROR_NEW, err)
		return
	}

	var createdAt time.Time = time.Now()
	var updatedAt time.Time
	var description = Description{contentDescription, useMarkdown}
	category = Category{uid, name, description, createdAt, updatedAt}

	return
}
