package ltests

import (
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type SubcategoryMockRepository[T Identifiable] struct {
	mock.Mock
}

func (repo *SubcategoryMockRepository[T]) AreRelatives(
	fatherUid uuid.UUID,
	childUid uuid.UUID,
) (bool, error) {
	args := repo.Called(fatherUid, childUid)
	return args.Bool(0), args.Error(1)
}

func (repo *SubcategoryMockRepository[T]) IsSubcategory(
	fatherUid uuid.UUID,
	childUid uuid.UUID,
) (bool, error) {
	args := repo.Called(fatherUid, childUid)
	return args.Bool(0), args.Error(1)
}

func (repo *SubcategoryMockRepository[T]) GetSubcategories(
	uid uuid.UUID,
) ([]T, error) {
	args := repo.Called(uid)
	return args.Get(0).([]T), args.Error(1)
}

func (repo *SubcategoryMockRepository[T]) Create(
	fatherUid uuid.UUID,
	childUid uuid.UUID,
	insertedAt time.Time,
) error {
	args := repo.Called(fatherUid, childUid, insertedAt)
	return args.Error(0)
}

func (repo *SubcategoryMockRepository[T]) Delete(
	fatherUid uuid.UUID,
	childUid uuid.UUID,
) error {
	args := repo.Called(fatherUid, childUid)
	return args.Error(0)
}
