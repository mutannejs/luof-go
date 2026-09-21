package belongs_to

import (
	"time"

	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
)

func (btr *BelongsTo) Create(
	linkUid uuid.UUID,
	categoryUid uuid.UUID,
	insertedAt time.Time,
	isMain bool,
) (lerror.ValueError) {
	var _, err = btr.DB.Exec(
		`
			INSERT INTO belongs_to (
				uid_link,
				uid_category,
				inserted_at,
				is_main
			) VALUES (
				?, ?, ?, ?
			)
		`,
		linkUid,
		categoryUid,
		insertedAt,
		isMain)

	if err != nil {
		return lerror.GetInternals(repository.WRITE_BELONGS_TO_CREATE_ERROR, err)
	} else {
		return lerror.ValueError{}
	}
}

func (btr *BelongsTo) Delete(
	linkUid uuid.UUID,
	categoryUid uuid.UUID,
) (lerror.ValueError) {
	var _, err = btr.DB.Exec(
		`
			DELETE FROM belongs_to WHERE uid_link = ? AND uid_category = ?
		`,
		linkUid,
		categoryUid)

	if err != nil {
		return lerror.GetInternals(repository.WRITE_BELONGS_TO_DELETE_ERROR, err)
	} else {
		return lerror.ValueError{}
	}
}

func (btr *BelongsTo) SetHasNoMainCategory(
	linkUid uuid.UUID,
) (lerror.ValueError) {
	var _, err = btr.DB.Exec(
		`
			UPDATE belongs_to
			SET is_main = false
			WHERE uid_link = ? AND is_main = true
		`,
		linkUid)

	if err != nil {
		return lerror.GetInternals(repository.WRITE_BELONGS_TO_SET_HAS_NO_MAIN_CATEGORY_ERROR, err)
	} else {
		return lerror.ValueError{}
	}
}

func (btr *BelongsTo) Update(
	linkUid uuid.UUID,
	categoryUid uuid.UUID,
	isMain bool,
) (lerror.ValueError) {
	var _, err = btr.DB.Exec(
		`
			UPDATE belongs_to
			SET is_main = ?
			WHERE uid_link = ? AND uid_category = ?
		`,
		isMain,
		linkUid,
		categoryUid)

	if err != nil {
		return lerror.GetInternals(repository.WRITE_BELONGS_TO_UPDATE_ERROR, err)
	} else {
		return lerror.ValueError{}
	}
}
