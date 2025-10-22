package ltests

import (
    "errors"
    "fmt"
    "github.com/google/uuid"
    "github.com/stretchr/testify/mock"
    "testing"
    _ "github.com/mattn/go-sqlite3"
)

const (
    DB_PATH = "./luof_test.db"
    DB_DRIVER = "sqlite3"
    PRINT_PREFIX = "ltests >>"
)

func FailIfError(t *testing.T, err error) {
    if err != nil {
        fmt.Println(PRINT_PREFIX, err)
        t.Fail()
    }
}

func PrintAndFail(t *testing.T, message string, err error) {
    fmt.Println(PRINT_PREFIX, message)
    if err != nil {
        fmt.Printf("\t  Error: %s\n", err)
    }
    t.Fail()
}

type Identifiable interface {
    GetUid() uuid.UUID
}

type MockCRUDRepository[T Identifiable] struct {
    items map[uuid.UUID]T
    populate bool
    mockItem T
}

func (repo *MockCRUDRepository[T]) Length() int {
    return len(repo.items)
}

func (repo *MockCRUDRepository[T]) GetItem(uid uuid.UUID) T {
    return repo.items[uid]
}

func NewMockCRUDRepository[T Identifiable](populate bool, mockItem T) *MockCRUDRepository[T] {
    var repo = MockCRUDRepository[T]{
        make(map[uuid.UUID]T),
        populate,
        mockItem,
    }
    if populate {
        repo.items[mockItem.GetUid()] = mockItem
    }
    return &repo
}

func (repo *MockCRUDRepository[T]) TestRepository(t *testing.T, testName string) {
    var qttitems int = len(repo.items)
    if repo.populate && qttitems != 1 || !repo.populate && qttitems != 0 {
        fmt.Printf("O teste %v não pode inicializar direito por erro na criação do repositório\n", testName)
        t.Fail()
    }
}

func (repo *MockCRUDRepository[T]) Exists(uid uuid.UUID) (bool, error) {
    _, ok := repo.items[uid]
    return ok == true, nil
}

func (repo *MockCRUDRepository[T]) GetByUid(uid uuid.UUID) (T, error) {
    item, ok := repo.items[uid]
    var err error
    if ok {
        err = nil
    } else {
        err = errors.New("Not found")
    }
    return item, err
}

func (repo *MockCRUDRepository[T]) Create(item T) error {
    repo.items[item.GetUid()] = item
    return nil
}

func (repo *MockCRUDRepository[T]) Delete(uid uuid.UUID) error {
    delete(repo.items, uid)
    return nil
}

func (repo *MockCRUDRepository[T]) Update(uid uuid.UUID, item T) error {
    repo.items[uid] = item
    return nil
}



type MockTestifyRepository[T Identifiable] struct {
    mock.Mock
}

func (repo *MockTestifyRepository[T]) Exists(uid uuid.UUID) (bool, error) {
    args := repo.Called(uid)
    return args.Bool(0), args.Error(1)
}

func (repo *MockTestifyRepository[T]) GetByUid(uid uuid.UUID) (T, error) {
    args := repo.Called(uid)
    return args.Get(0).(T), args.Error(1)
}

func (repo *MockTestifyRepository[T]) Create(item T) error {
    args := repo.Called(item)
    return args.Error(0)
}

func (repo *MockTestifyRepository[T]) Delete(uid uuid.UUID) error {
    args := repo.Called(uid)
    return args.Error(0)
}

func (repo *MockTestifyRepository[T]) Update(uid uuid.UUID, item T) error {
    args := repo.Called(uid, item)
    return args.Error(0)
}
