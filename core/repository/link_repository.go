package repository

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
)

type ReadLink interface {
	Exists(uuid.UUID) (bool, lerror.ValueError)
	GetByUid(uuid.UUID) (domain.Link, lerror.ValueError)
}

type WriteLink interface {
	Create(domain.Link) lerror.ValueError
	Delete(uuid.UUID) lerror.ValueError
	Update(uuid.UUID, domain.Link) lerror.ValueError
}

type Link interface {
	ReadLink
	WriteLink
}
