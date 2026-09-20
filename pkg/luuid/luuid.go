package luuid

import (
	"reflect"

	"github.com/mutannejs/luof-go/pkg/li18n"

	"github.com/google/uuid"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	UUID_ERROR_NEW = li18n.Msg2Error(&i18n.Message{
		ID: "UUID_ERROR_NEW",
		Description: "Erro retornado quando uuid.New falha",
		Other: "error generating new uuid",
	})
)

func Zero() uuid.UUID {
	var zero uuid.UUID
	return zero
}

func IsZero(uid uuid.UUID) bool {
	var zero uuid.UUID
	return reflect.DeepEqual(uid, zero)
}
