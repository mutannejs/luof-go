package ltests

import (
	"time"

	"github.com/google/uuid"
	"github.com/mutannejs/luof-go/pkg/lerror"
	"github.com/stretchr/testify/mock"
)

type BelongsToMockRepository[T Identifiable] struct {
	mock.Mock
}

func (repo *BelongsToMockRepository[T]) Exists(
	uidLink uuid.UUID,
	uidCategory uuid.UUID,
) (bool, lerror.ValueError) {
	args := repo.Called(uidLink, uidCategory)
	return args.Bool(0), lerror.GetInternal(args.Error(1))
}

func (repo *BelongsToMockRepository[T]) HasLinks(
	uidCategory uuid.UUID,
) (bool, lerror.ValueError) {
	args := repo.Called(uidCategory)
	return args.Bool(0), lerror.GetInternal(args.Error(1))
}

func (repo *BelongsToMockRepository[T]) GetLinksByCategory(
	uid uuid.UUID,
) ([]T, lerror.ValueError) {
	args := repo.Called(uid)
	return args.Get(0).([]T), lerror.GetInternal(args.Error(1))
}

func (repo *BelongsToMockRepository[T]) Create(
	uidLink uuid.UUID,
	uidCategory uuid.UUID,
	insertedAt time.Time,
	isMain bool,
) lerror.ValueError {
	args := repo.Called(uidLink, uidCategory, insertedAt, isMain)
	return lerror.GetInternal(args.Error(0))
}

func (repo *BelongsToMockRepository[T]) Delete(
	uidLink uuid.UUID,
	uidCategory uuid.UUID,
) lerror.ValueError {
	args := repo.Called(uidLink, uidCategory)
	return lerror.GetInternal(args.Error(0))
}

func (repo *BelongsToMockRepository[T]) Update(
	uidLink uuid.UUID,
	uidCategory uuid.UUID,
	isMain bool,
) lerror.ValueError {
	args := repo.Called(uidLink, uidCategory, isMain)
	return lerror.GetInternal(args.Error(0))
}

func (repo *BelongsToMockRepository[T]) SetHasNoMainCategory(
	uidLink uuid.UUID,
) lerror.ValueError {
	args := repo.Called(uidLink)
	return lerror.GetInternal(args.Error(0))
}
