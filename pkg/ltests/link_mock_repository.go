package ltests

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func NewLinkMockRepository() *LinkMockRepository[domain.Link] {
	return &LinkMockRepository[domain.Link]{}
}

type LinkMockRepository[T Identifiable] struct {
	mock.Mock
}

func (repo *LinkMockRepository[T]) Exists(uid uuid.UUID) (bool, lerror.ValueError) {
	args := repo.Called(uid)
	if args.Error(1) == nil {
		return args.Bool(0), lerror.ValueError{}
	}
	return args.Bool(0), lerror.GetInternals(repository.READ_LINK_EXISTS_ERROR, args.Error(1))
}

func (repo *LinkMockRepository[T]) GetByUid(uid uuid.UUID) (T, lerror.ValueError) {
	args := repo.Called(uid)
	if args.Error(1) == nil {
		return args.Get(0).(T), lerror.ValueError{}
	}
	return args.Get(0).(T), lerror.GetInternals(repository.READ_CATEGORY_GET_BY_UID_ERROR, args.Error(1))
}

func (repo *LinkMockRepository[T]) Create(item T) lerror.ValueError {
	args := repo.Called(item)
	if args.Error(0) == nil {
		return lerror.ValueError{}
	}
	return lerror.GetInternals(repository.WRITE_LINK_CREATE_ERROR, args.Error(0))
}

func (repo *LinkMockRepository[T]) Delete(uid uuid.UUID) lerror.ValueError {
	args := repo.Called(uid)
	if args.Error(0) == nil {
		return lerror.ValueError{}
	}
	return lerror.GetInternals(repository.WRITE_LINK_DELETE_ERROR, args.Error(0))
}

func (repo *LinkMockRepository[T]) Update(uid uuid.UUID, item T) lerror.ValueError {
	args := repo.Called(uid, item)
	if args.Error(0) == nil {
		return lerror.ValueError{}
	}
	return lerror.GetInternals(repository.WRITE_LINK_UPDATE_ERROR, args.Error(0))
}
