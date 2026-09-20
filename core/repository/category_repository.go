package repository

import (
	"time"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/google/uuid"
)

var (
	READ_CATEGORY_EXISTS_ERROR = &i18n.Message{
		ID: "READ_CATEGORY_EXISTS_ERROR",
		Description: "Retornado quando category_repository.Exists falha",
		Other: "error verifying in database the existence of a category",
	}
	READ_CATEGORY_GET_BY_UID_ERROR = &i18n.Message{
		ID: "READ_CATEGORY_GET_BY_UID_ERROR",
		Description: "Retornado quando category_repository.GetByUid falha",
		Other: "error querying category in database",
	}
	READ_CATEGORY_ARE_RELATED_ERROR = &i18n.Message{
		ID: "READ_CATEGORY_ARE_RELATED_ERROR",
		Description: "Retornado quando category_repository.AreRelated falha",
		Other: "error verifying in the database if two categories are related",
	}
	READ_CATEGORY_GET_ALL_ROOT_CATEGORIES_ERROR = &i18n.Message{
		ID: "READ_CATEGORY_GET_ALL_ROOT_CATEGORIES_ERROR",
		Description: "Retornado quando category_repository.GetAllRootCategories falha",
		Other: "error querying all root categories in the database",
	}
	READ_CATEGORY_GET_SUBCATEGORIES_ERROR = &i18n.Message{
		ID: "READ_CATEGORY_GET_SUBCATEGORIES_ERROR",
		Description: "Retornado quando category_repository.GetSubcategories falha",
		Other: "error querying subcategories in the database ",
	}
	READ_CATEGORY_HAS_SUBCATEGORIES_ERROR = &i18n.Message{
		ID: "READ_CATEGORY_HAS_SUBCATEGORIES_ERROR",
		Description: "Retornado quando category_repository.HasSubcategories falha",
		Other: "error verifying in the database if category has subcategories",
	}
	READ_CATEGORY_IS_ANCESTOR_ERROR = &i18n.Message{
		ID: "READ_CATEGORY_IS_ANCESTOR_ERROR",
		Description: "Retornado quando category_repository.IsAncestor falha",
		Other: "error verifying in the database if category is an ancestor of another category",
	}
	READ_CATEGORY_IS_SUBCATEGORY_ERROR = &i18n.Message{
		ID: "READ_CATEGORY_IS_SUBCATEGORY_ERROR",
		Description: "Retornado quando category_repository.IsSubcategory falha",
		Other: "error verifying in the database if category is an subcategory of another category",
	}
	WRITE_CATEGORY_CREATE_ERROR = &i18n.Message{
		ID: "WRITE_CATEGORY_CREATE_ERROR",
		Description: "Retornado quando category_repository.Create falha",
		Other: "error creating category in database",
	}
	WRITE_CATEGORY_DELETE_ERROR = &i18n.Message{
		ID: "WRITE_CATEGORY_DELETE_ERROR",
		Description: "Retornado quando category_repository.Delete falha",
		Other: "error deleting category in database",
	}
	WRITE_CATEGORY_UPDATE_ERROR = &i18n.Message{
		ID: "WRITE_CATEGORY_UPDATE_ERROR",
		Description: "Retornado quando category_repository.Update falha",
		Other: "error updating category in database",
	}
	WRITE_CATEGORY_DELETE_SUBCATEGORY_ERROR = &i18n.Message{
		ID: "WRITE_CATEGORY_DELETE_SUBCATEGORY_ERROR",
		Description: "Retornado quando category_repository.DeleteSubcategory falha",
		Other: "error deleting subcategory in database",
	}
	WRITE_CATEGORY_INSERT_SUBCATEGORY_ERROR = &i18n.Message{
		ID: "WRITE_CATEGORY_INSERT_SUBCATEGORY_ERROR",
		Description: "Retornado quando category_repository.InsertSubcategory falha",
		Other: "error inserting subcategory in database",
	}
)

type ReadCategory interface {
	Exists(uuid.UUID) (bool, lerror.ValueError)
	GetByUid(uuid.UUID) (domain.Category, lerror.ValueError)
}

type ReadSubcategory interface {
	AreRelated(uuid.UUID, uuid.UUID) (bool, lerror.ValueError)
	GetAllRootCategories() ([]domain.Category, lerror.ValueError)
	GetSubcategories(uuid.UUID) ([]domain.Category, lerror.ValueError)
	HasSubcategories(uuid.UUID) (bool, lerror.ValueError)
	IsAncestor(uuid.UUID, uuid.UUID) (bool, lerror.ValueError)
	IsSubcategory(uuid.UUID, uuid.UUID) (bool, lerror.ValueError)
}

type WriteCategory interface {
	Create(domain.Category) lerror.ValueError
	Delete(uuid.UUID) lerror.ValueError
	Update(uuid.UUID, domain.Category) lerror.ValueError
}

type WriteSubcategory interface {
	DeleteSubcategory(uuid.UUID) lerror.ValueError
	InsertSubcategory(uuid.UUID, uuid.UUID, time.Time) lerror.ValueError
}

type Category interface {
	ReadCategory
	ReadSubcategory
	WriteCategory
	WriteSubcategory
}
