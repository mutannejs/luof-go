package repository

import (
	"time"

	"github.com/mutannejs/luof-go/core/domain"

	"github.com/google/uuid"
)

type ReadSubcategory interface {
	AreRelatives(uuid.UUID, uuid.UUID) (bool, error)
	IsSubcategory(uuid.UUID, uuid.UUID) (bool, error)
	GetSubcategories(uuid.UUID) ([]domain.Category, error)
}

type WriteSubcategory interface {
	Create(uuid.UUID, uuid.UUID, time.Time) error
	Delete(uuid.UUID, uuid.UUID) error
}

type Subcategory interface {
	ReadSubcategory
	WriteSubcategory
}
