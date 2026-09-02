package ltests

import (
	"time"

	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type CategoryMockRepository[T Identifiable] struct {
	mock.Mock
}

func (repo *CategoryMockRepository[T]) AreRelated(
	fatherUid uuid.UUID,
	childUid uuid.UUID,
) (bool, lerror.ValueError) {
	args := repo.Called(fatherUid, childUid)
	return args.Bool(0), lerror.GetInternal(args.Error(1))
}

func (repo *CategoryMockRepository[T]) Create(item T) lerror.ValueError {
	args := repo.Called(item)
	return lerror.GetInternal(args.Error(0))
}

func (repo *CategoryMockRepository[T]) Delete(uid uuid.UUID) lerror.ValueError {
	args := repo.Called(uid)
	return lerror.GetInternal(args.Error(0))
}

func (repo *CategoryMockRepository[T]) DeleteSubcategory(
	childUid uuid.UUID,
) lerror.ValueError {
	args := repo.Called(childUid)
	return lerror.GetInternal(args.Error(0))
}

func (repo *CategoryMockRepository[T]) Exists(uid uuid.UUID) (bool, lerror.ValueError) {
	args := repo.Called(uid)
	return args.Bool(0), lerror.GetInternal(args.Error(1))
}

func (repo *CategoryMockRepository[T]) GetAllRootCategories() ([]T, lerror.ValueError) {
	args := repo.Called()
	return args.Get(0).([]T), lerror.GetInternal(args.Error(1))
}

func (repo *CategoryMockRepository[T]) GetByUid(uid uuid.UUID) (T, lerror.ValueError) {
	args := repo.Called(uid)
	return args.Get(0).(T), lerror.GetInternal(args.Error(1))
}

func (repo *CategoryMockRepository[T]) GetSubcategories(
	uid uuid.UUID,
) ([]T, lerror.ValueError) {
	args := repo.Called(uid)
	return args.Get(0).([]T), lerror.GetInternal(args.Error(1))
}

func (repo *CategoryMockRepository[T]) HasSubcategories(uid uuid.UUID) (bool, lerror.ValueError) {
	args := repo.Called(uid)
	return args.Bool(0), lerror.GetInternal(args.Error(1))
}

func (repo *CategoryMockRepository[T]) InsertSubcategory(
	fatherUid uuid.UUID,
	childUid uuid.UUID,
	updatedAt time.Time,
) lerror.ValueError {
	args := repo.Called(fatherUid, childUid, updatedAt)
	return lerror.GetInternal(args.Error(0))
}

func (repo *CategoryMockRepository[T]) IsAncestor(
	ancestorUid uuid.UUID,
	categoryUid uuid.UUID,
) (bool, lerror.ValueError) {
	args := repo.Called(ancestorUid, categoryUid)
	return args.Bool(0), lerror.GetInternal(args.Error(1))
}

func (repo *CategoryMockRepository[T]) IsSubcategory(
	fatherUid uuid.UUID,
	childUid uuid.UUID,
) (bool, lerror.ValueError) {
	args := repo.Called(fatherUid, childUid)
	return args.Bool(0), lerror.GetInternal(args.Error(1))
}

func (repo *CategoryMockRepository[T]) Update(uid uuid.UUID, item T) lerror.ValueError {
	args := repo.Called(uid, item)
	return lerror.GetInternal(args.Error(0))
}
