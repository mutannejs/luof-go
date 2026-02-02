package ltests

import (
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockCategoryRepository[T Identifiable] struct {
	mock.Mock
}

func (repo *MockCategoryRepository[T]) AreRelated(
	fatherUid uuid.UUID,
	childUid uuid.UUID,
) (bool, error) {
	args := repo.Called(fatherUid, childUid)
	return args.Bool(0), args.Error(1)
}

func (repo *MockCategoryRepository[T]) Create(item T) error {
	args := repo.Called(item)
	return args.Error(0)
}

func (repo *MockCategoryRepository[T]) Delete(uid uuid.UUID) error {
	args := repo.Called(uid)
	return args.Error(0)
}

func (repo *MockCategoryRepository[T]) DeleteSubcategory(
	childUid uuid.UUID,
) error {
	args := repo.Called(childUid)
	return args.Error(0)
}

func (repo *MockCategoryRepository[T]) Exists(uid uuid.UUID) (bool, error) {
	args := repo.Called(uid)
	return args.Bool(0), args.Error(1)
}

func (repo *MockCategoryRepository[T]) GetByUid(uid uuid.UUID) (T, error) {
	args := repo.Called(uid)
	return args.Get(0).(T), args.Error(1)
}

func (repo *MockCategoryRepository[T]) GetSubcategories(
	uid uuid.UUID,
) ([]T, error) {
	args := repo.Called(uid)
	return args.Get(0).([]T), args.Error(1)
}

func (repo *MockCategoryRepository[T]) InsertSubcategory(
	fatherUid uuid.UUID,
	childUid uuid.UUID,
	updatedAt time.Time,
) error {
	args := repo.Called(fatherUid, childUid, updatedAt)
	return args.Error(0)
}

func (repo *MockCategoryRepository[T]) IsSubcategory(
	fatherUid uuid.UUID,
	childUid uuid.UUID,
) (bool, error) {
	args := repo.Called(fatherUid, childUid)
	return args.Bool(0), args.Error(1)
}

func (repo *MockCategoryRepository[T]) Update(uid uuid.UUID, item T) error {
	args := repo.Called(uid, item)
	return args.Error(0)
}
