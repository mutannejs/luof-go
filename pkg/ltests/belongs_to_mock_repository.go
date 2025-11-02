package ltests

import (
    "github.com/google/uuid"
    "github.com/stretchr/testify/mock"
)

type BelongsToMockRepository[T Identifiable] struct {
    mock.Mock
}

func (repo *BelongsToMockRepository[T]) GetLinksByCategory(uid uuid.UUID) ([]T, error) {
    args := repo.Called(uid)
    return args.Get(0).([]T), args.Error(1)
}
