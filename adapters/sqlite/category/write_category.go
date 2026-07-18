package category

import (
	"time"

	"github.com/mutannejs/luof-go/core/domain"

	"github.com/google/uuid"
)

func (cr *Category) Create(l domain.Category) (err error) {
	_, err = cr.DB.Exec(
		`
			INSERT INTO category (
				uid_category,
				name,
				description,
				use_markdown,
				created_at,
				updated_at
			) VALUES (
				?, ?, ?, ?, ?, ?
			)
		`,
		l.GetUid(),
		l.Name,
		l.Description.Content,
		l.Description.UseMarkdown,
		l.CreatedAt,
		l.UpdatedAt)

	return
}

func (cr *Category) Delete(uid uuid.UUID) (err error) {
	_, err = cr.DB.Exec(
		`DELETE FROM category WHERE uid_category = ?`,
		uid)

	return
}

func (cr *Category) DeleteSubcategory(
	childUid uuid.UUID,
) (err error) {
	_, err = cr.DB.Exec(
		`
			UPDATE category
			SET uid_father = null
			WHERE uid_category = ?
		`,
		childUid)

	return
}

func (cr *Category) InsertSubcategory(
	fatherUid uuid.UUID,
	childUid uuid.UUID,
	updatedAt time.Time,
) (err error) {
	_, err = cr.DB.Exec(
		`
			UPDATE category
			SET uid_father = ?,
				updated_at = ?
			WHERE uid_category = ?
		`,
		fatherUid,
		updatedAt,
		childUid)

	return
}

func (cr *Category) Update(uid uuid.UUID, l domain.Category) (err error) {
	_, err = cr.DB.Exec(
		`
			UPDATE category
			SET name = ?,
				description = ?,
				use_markdown = ?,
				created_at = ?,
				updated_at = ?
			WHERE uid_category = ?
		`,
		l.Name,
		l.Description.Content,
		l.Description.UseMarkdown,
		l.CreatedAt,
		l.UpdatedAt,
		uid)

	return
}
