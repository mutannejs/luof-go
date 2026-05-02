package repository

import (
	"time"

	"github.com/mutannejs/luof-go/core/domain"

	"github.com/google/uuid"
)

type ReadCategory interface {
	Exists(uuid.UUID) (bool, error)
	GetByUid(uuid.UUID) (domain.Category, error)
}

type ReadSubcategory interface {
	AreRelated(uuid.UUID, uuid.UUID) (bool, error)
	GetSubcategories(uuid.UUID) ([]domain.Category, error)
	HasSubcategories(uuid.UUID) (bool, error)
	IsAncestor(uuid.UUID, uuid.UUID) (bool, error)
	IsSubcategory(uuid.UUID, uuid.UUID) (bool, error)
}

type WriteCategory interface {
	Create(domain.Category) error
	Delete(uuid.UUID) error
	DeleteSubcategory(uuid.UUID) error
	InsertSubcategory(uuid.UUID, uuid.UUID, time.Time) error
	Update(uuid.UUID, domain.Category) error
}

type WriteSubcategory interface {
	DeleteSubcategory(uuid.UUID) error
	InsertSubcategory(uuid.UUID, uuid.UUID, time.Time) error
}

type Category interface {
	ReadCategory
	ReadSubcategory
	WriteCategory
	WriteSubcategory
}
