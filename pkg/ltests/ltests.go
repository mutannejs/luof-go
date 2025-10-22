package ltests

import (
    "github.com/google/uuid"
    "github.com/stretchr/testify/mock"
)

type Identifiable interface {
    GetUid() uuid.UUID
}

type MockCrudRepository[T Identifiable] struct {
    mock.Mock
}

func (repo *MockCrudRepository[T]) Exists(uid uuid.UUID) (bool, error) {
    args := repo.Called(uid)
    return args.Bool(0), args.Error(1)
}

func (repo *MockCrudRepository[T]) GetByUid(uid uuid.UUID) (T, error) {
    args := repo.Called(uid)
    return args.Get(0).(T), args.Error(1)
}

func (repo *MockCrudRepository[T]) Create(item T) error {
    args := repo.Called(item)
    return args.Error(0)
}

func (repo *MockCrudRepository[T]) Delete(uid uuid.UUID) error {
    args := repo.Called(uid)
    return args.Error(0)
}

func (repo *MockCrudRepository[T]) Update(uid uuid.UUID, item T) error {
    args := repo.Called(uid, item)
    return args.Error(0)
}
