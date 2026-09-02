package ltests

import (
	"github.com/google/uuid"
	"github.com/mutannejs/luof-go/pkg/lerror"
	"github.com/stretchr/testify/mock"
)

type LinkMockRepository[T Identifiable] struct {
	mock.Mock
}

func (repo *LinkMockRepository[T]) Exists(uid uuid.UUID) (bool, lerror.ValueError) {
	args := repo.Called(uid)
	return args.Bool(0), lerror.GetInternal(args.Error(1))
}

func (repo *LinkMockRepository[T]) GetByUid(uid uuid.UUID) (T, lerror.ValueError) {
	args := repo.Called(uid)
	return args.Get(0).(T), lerror.GetInternal(args.Error(1))
}

func (repo *LinkMockRepository[T]) Create(item T) lerror.ValueError {
	args := repo.Called(item)
	return lerror.GetInternal(args.Error(0))
}

func (repo *LinkMockRepository[T]) Delete(uid uuid.UUID) lerror.ValueError {
	args := repo.Called(uid)
	return lerror.GetInternal(args.Error(0))
}

func (repo *LinkMockRepository[T]) Update(uid uuid.UUID, item T) lerror.ValueError {
	args := repo.Called(uid, item)
	return lerror.GetInternal(args.Error(0))
}
