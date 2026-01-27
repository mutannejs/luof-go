package ltests

import (
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type BelongsToMockRepository[T Identifiable] struct {
	mock.Mock
}

func (repo *BelongsToMockRepository[T]) Exists(
	uidLink uuid.UUID,
	uidCategory uuid.UUID,
) (bool, error) {
	args := repo.Called(uidLink, uidCategory)
	return args.Bool(0), args.Error(1)
}

func (repo *BelongsToMockRepository[T]) GetLinksByCategory(
	uid uuid.UUID,
) ([]T, error) {
	args := repo.Called(uid)
	return args.Get(0).([]T), args.Error(1)
}

func (repo *BelongsToMockRepository[T]) Create(
	uidLink uuid.UUID,
	uidCategory uuid.UUID,
	insertedAt time.Time,
	isMain bool,
) error {
	args := repo.Called(uidLink, uidCategory, insertedAt, isMain)
	return args.Error(0)
}

func (repo *BelongsToMockRepository[T]) Delete(
	uidLink uuid.UUID,
	uidCategory uuid.UUID,
) error {
	args := repo.Called(uidLink, uidCategory)
	return args.Error(0)
}

func (repo *BelongsToMockRepository[T]) Update(
	uidLink uuid.UUID,
	uidCategory uuid.UUID,
	isMain bool,
) error {
	args := repo.Called(uidLink, uidCategory, isMain)
	return args.Error(0)
}
