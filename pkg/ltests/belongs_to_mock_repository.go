package ltests

import (
	"time"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func NewBelongsToMockRepository() *BelongsToMockRepository[domain.Link] {
	return &BelongsToMockRepository[domain.Link]{}
}

type BelongsToMockRepository[T Identifiable] struct {
	mock.Mock
}

func (repo *BelongsToMockRepository[T]) Exists(
	uidLink uuid.UUID,
	uidCategory uuid.UUID,
) (bool, lerror.ValueError) {
	args := repo.Called(uidLink, uidCategory)
	if args.Error(1) == nil {
		return args.Bool(0), lerror.ValueError{}
	}
	return args.Bool(0), lerror.GetInternals(repository.READ_BELONGS_TO_EXISTS_ERROR, args.Error(1))
}

func (repo *BelongsToMockRepository[T]) HasLinks(
	uidCategory uuid.UUID,
) (bool, lerror.ValueError) {
	args := repo.Called(uidCategory)
	if args.Error(1) == nil {
		return args.Bool(0), lerror.ValueError{}
	}
	return args.Bool(0), lerror.GetInternals(repository.READ_BELONGS_TO_HAS_LINKS_ERROR, args.Error(1))
}

func (repo *BelongsToMockRepository[T]) GetLinksByCategory(
	uid uuid.UUID,
) ([]T, lerror.ValueError) {
	args := repo.Called(uid)
	if args.Error(1) == nil {
		return args.Get(0).([]T), lerror.ValueError{}
	}
	return args.Get(0).([]T), lerror.GetInternals(repository.READ_BELONGS_TO_GET_LINKS_BY_CATEGORY_ERROR, args.Error(1))
}

func (repo *BelongsToMockRepository[T]) Create(
	uidLink uuid.UUID,
	uidCategory uuid.UUID,
	insertedAt time.Time,
	isMain bool,
) lerror.ValueError {
	args := repo.Called(uidLink, uidCategory, insertedAt, isMain)
	if args.Error(0) == nil {
		return lerror.ValueError{}
	}
	return lerror.GetInternals(repository.WRITE_BELONGS_TO_CREATE_ERROR, args.Error(0))
}

func (repo *BelongsToMockRepository[T]) Delete(
	uidLink uuid.UUID,
	uidCategory uuid.UUID,
) lerror.ValueError {
	args := repo.Called(uidLink, uidCategory)
	if args.Error(0) == nil {
		return lerror.ValueError{}
	}
	return lerror.GetInternals(repository.WRITE_BELONGS_TO_DELETE_ERROR, args.Error(0))
}

func (repo *BelongsToMockRepository[T]) Update(
	uidLink uuid.UUID,
	uidCategory uuid.UUID,
	isMain bool,
) lerror.ValueError {
	args := repo.Called(uidLink, uidCategory, isMain)
	if args.Error(0) == nil {
		return lerror.ValueError{}
	}
	return lerror.GetInternals(repository.WRITE_BELONGS_TO_UPDATE_ERROR, args.Error(0))
}

func (repo *BelongsToMockRepository[T]) SetHasNoMainCategory(
	uidLink uuid.UUID,
) lerror.ValueError {
	args := repo.Called(uidLink)
	if args.Error(0) == nil {
		return lerror.ValueError{}
	}
	return lerror.GetInternals(repository.WRITE_BELONGS_TO_SET_HAS_NO_MAIN_CATEGORY_ERROR, args.Error(0))
}
