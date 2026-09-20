package ltests

import (
	"time"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func NewCategoryMockRepository() *CategoryMockRepository[domain.Category] {
	return &CategoryMockRepository[domain.Category]{}
}

type CategoryMockRepository[T Identifiable] struct {
	mock.Mock
}

func (repo *CategoryMockRepository[T]) AreRelated(
	fatherUid uuid.UUID,
	childUid uuid.UUID,
) (bool, lerror.ValueError) {
	args := repo.Called(fatherUid, childUid)
	return args.Bool(0), lerror.GetInternals(repository.READ_CATEGORY_ARE_RELATED_ERROR, args.Error(1))
}

func (repo *CategoryMockRepository[T]) Create(item T) lerror.ValueError {
	args := repo.Called(item)
	return lerror.GetInternals(repository.WRITE_CATEGORY_CREATE_ERROR, args.Error(0))
}

func (repo *CategoryMockRepository[T]) Delete(uid uuid.UUID) lerror.ValueError {
	args := repo.Called(uid)
	return lerror.GetInternals(repository.WRITE_CATEGORY_DELETE_ERROR, args.Error(0))
}

func (repo *CategoryMockRepository[T]) DeleteSubcategory(
	childUid uuid.UUID,
) lerror.ValueError {
	args := repo.Called(childUid)
	return lerror.GetInternals(repository.WRITE_CATEGORY_DELETE_SUBCATEGORY_ERROR, args.Error(0))
}

func (repo *CategoryMockRepository[T]) Exists(uid uuid.UUID) (bool, lerror.ValueError) {
	args := repo.Called(uid)
	return args.Bool(0), lerror.GetInternals(repository.READ_CATEGORY_EXISTS_ERROR, args.Error(1))
}

func (repo *CategoryMockRepository[T]) GetAllRootCategories() ([]T, lerror.ValueError) {
	args := repo.Called()
	return args.Get(0).([]T), lerror.GetInternals(repository.READ_CATEGORY_GET_ALL_ROOT_CATEGORIES_ERROR, args.Error(1))
}

func (repo *CategoryMockRepository[T]) GetByUid(uid uuid.UUID) (T, lerror.ValueError) {
	args := repo.Called(uid)
	return args.Get(0).(T), lerror.GetInternals(repository.READ_CATEGORY_GET_BY_UID_ERROR, args.Error(1))
}

func (repo *CategoryMockRepository[T]) GetSubcategories(
	uid uuid.UUID,
) ([]T, lerror.ValueError) {
	args := repo.Called(uid)
	return args.Get(0).([]T), lerror.GetInternals(repository.READ_CATEGORY_GET_SUBCATEGORIES_ERROR, args.Error(1))
}

func (repo *CategoryMockRepository[T]) HasSubcategories(uid uuid.UUID) (bool, lerror.ValueError) {
	args := repo.Called(uid)
	return args.Bool(0), lerror.GetInternals(repository.READ_CATEGORY_HAS_SUBCATEGORIES_ERROR, args.Error(1))
}

func (repo *CategoryMockRepository[T]) InsertSubcategory(
	fatherUid uuid.UUID,
	childUid uuid.UUID,
	updatedAt time.Time,
) lerror.ValueError {
	args := repo.Called(fatherUid, childUid, updatedAt)
	return lerror.GetInternals(repository.WRITE_CATEGORY_INSERT_SUBCATEGORY_ERROR, args.Error(0))
}

func (repo *CategoryMockRepository[T]) IsAncestor(
	ancestorUid uuid.UUID,
	categoryUid uuid.UUID,
) (bool, lerror.ValueError) {
	args := repo.Called(ancestorUid, categoryUid)
	return args.Bool(0), lerror.GetInternals(repository.READ_CATEGORY_IS_ANCESTOR_ERROR, args.Error(1))
}

func (repo *CategoryMockRepository[T]) IsSubcategory(
	fatherUid uuid.UUID,
	childUid uuid.UUID,
) (bool, lerror.ValueError) {
	args := repo.Called(fatherUid, childUid)
	return args.Bool(0), lerror.GetInternals(repository.READ_CATEGORY_IS_SUBCATEGORY_ERROR, args.Error(1))
}

func (repo *CategoryMockRepository[T]) Update(uid uuid.UUID, item T) lerror.ValueError {
	args := repo.Called(uid, item)
	return lerror.GetInternals(repository.READ_CATEGORY_ARE_RELATED_ERROR, args.Error(0))
}
