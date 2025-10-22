package luuid

import (
    "testing"

    "github.com/google/uuid"
    "github.com/stretchr/testify/assert"
)

func TestIsZero(t *testing.T) {
    var uid uuid.UUID
    assert.True(t, IsZero(uid), "um zero deveria ser reconhecido pela função IsZero")
    assert.False(t, IsZero(uuid.New()), "um uuid válido não deveria ser reconhecido pela função IsZero")
}

func TestZero(t *testing.T) {
    var zero = Zero()
    assert.Zero(t, zero, "a função Zero deveria retornar um zero do tipo uuid.UUID")
}

func TestNew(t *testing.T) {
    var uid, err = New()
    if err == nil {
        assert.NotZero(t, uid, "a função New deveria retornar um uuid válido")
    } else {
        assert.EqualError(t, err, UUID_ERROR_NEW.Error(), "o erro retornado por New deveria ser " + UUID_ERROR_NEW.Error())
    }
}
