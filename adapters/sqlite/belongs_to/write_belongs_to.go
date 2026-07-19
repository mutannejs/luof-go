package belongs_to

import (
	"time"

	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
)

func (btr *BelongsTo) Create(
	linkUid uuid.UUID,
	categoryUid uuid.UUID,
	insertedAt time.Time,
	isMain bool,
) (err error) {
	_, err = btr.DB.Exec(
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

	lerror.SetInternal(&err)
	return
}

func (btr *BelongsTo) Delete(
	linkUid uuid.UUID,
	categoryUid uuid.UUID,
) (err error) {
	_, err = btr.DB.Exec(
		`
			DELETE FROM belongs_to WHERE uid_link = ? AND uid_category = ?
		`,
		linkUid,
		categoryUid)

	lerror.SetInternal(&err)
	return
}

func (btr *BelongsTo) SetHasNoMainCategory(
	linkUid uuid.UUID,
) (err error) {
	_, err = btr.DB.Exec(
		`
			UPDATE belongs_to
			SET is_main = false
			WHERE uid_link = ? AND is_main = true
		`,
		linkUid)

	lerror.SetInternal(&err)
	return
}

func (btr *BelongsTo) Update(
	linkUid uuid.UUID,
	categoryUid uuid.UUID,
	isMain bool,
) (err error) {
	_, err = btr.DB.Exec(
		`
			UPDATE belongs_to
			SET is_main = ?
			WHERE uid_link = ? AND uid_category = ?
		`,
		isMain,
		linkUid,
		categoryUid)

	lerror.SetInternal(&err)
	return
}
