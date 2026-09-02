package repository

import (
	"time"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
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
	DeleteSubcategory(uuid.UUID) lerror.ValueError
	InsertSubcategory(uuid.UUID, uuid.UUID, time.Time) lerror.ValueError
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
