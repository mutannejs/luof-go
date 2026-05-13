package luuid

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestIsZero(t *testing.T) {
	var assert = assert.New(t)
	var uid uuid.UUID

	assert.True(
		IsZero(uid),
		"Um zero deveria ser reconhecido pela função IsZero")
	assert.False(
		IsZero(uuid.New()),
		"Um uuid válido não deveria ser reconhecido pela função IsZero")
}

func TestZero(t *testing.T) {
	var assert = assert.New(t)
	var zero = Zero()

	assert.Zero(
		zero,
		"A função Zero deveria retornar um zero do tipo uuid.UUID")
}

func TestNew(t *testing.T) {
	var assert = assert.New(t)
	var uid, err = New()

	if err == nil {
		assert.NotZero(
			uid,
			"A função New deveria retornar um uuid válido se nenhum erro ocorrer")
	} else {
		assert.EqualError(
			UUID_ERROR_NEW,
			err.Error(),
			"O erro retornado por New deveria ser " + UUID_ERROR_NEW.Error())
	}
}
