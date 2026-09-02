package domain

import (
	"time"

	"github.com/mutannejs/luof-go/pkg/lerror"
	"github.com/mutannejs/luof-go/pkg/luuid"

	"github.com/google/uuid"
)

var (
	ANCESTOR_NOT_BECOME_A_SUBCATEGORY = "one ancestral category of another cannot become a subcategory of it"
	ARE_RELATED = "both categories are already related"
	CANNOT_BE_A_SUBCATEGORY_OF_ITSELF = "a category cannot be a subcategory of itself"
	CATEGORY_ERROR_NEW = "error instantiate new category"
	CATEGORY_NOT_EXISTS = "the searched category does not exist"
	CHILD_NOT_EXISTS = "the child category does not exist"
	FATHER_NOT_EXISTS = "the father category does not exist"
	IS_SUBCATEGORY = "the child is already subcategory of the father"
	HAS_SUBCATEGORIES = "the category cannot be deleted because it has subcategories"
	NOT_IS_SUBCATEGORY = "the child is not subcategory of the father"
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
