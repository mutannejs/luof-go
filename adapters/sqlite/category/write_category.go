package category

import (
	"time"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
)

func (cr *Category) Create(l domain.Category) (lerror.ValueError) {
	var _, err = cr.DB.Exec(
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
		
	return lerror.GetInternal(err)
}

func (cr *Category) Delete(uid uuid.UUID) (lerror.ValueError) {
	var _, err = cr.DB.Exec(
		`DELETE FROM category WHERE uid_category = ?`,
		uid)
		
	return lerror.GetInternal(err)
}

func (cr *Category) DeleteSubcategory(
	childUid uuid.UUID,
) (lerror.ValueError) {
	var _, err = cr.DB.Exec(
		`
			UPDATE category
			SET uid_father = null
			WHERE uid_category = ?
		`,
		childUid)
		
	return lerror.GetInternal(err)
}

func (cr *Category) InsertSubcategory(
	fatherUid uuid.UUID,
	childUid uuid.UUID,
	updatedAt time.Time,
) (lerror.ValueError) {
	var _, err = cr.DB.Exec(
		`
			UPDATE category
			SET uid_father = ?,
				updated_at = ?
			WHERE uid_category = ?
		`,
		fatherUid,
		updatedAt,
		childUid)
		
	return lerror.GetInternal(err)
}

func (cr *Category) Update(
	uid uuid.UUID,
	l domain.Category,
) (lerror.ValueError) {
	var _, err = cr.DB.Exec(
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
		
	return lerror.GetInternal(err)
}
