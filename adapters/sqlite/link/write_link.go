package link

import (
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
)

func (lr *Link) Create(l domain.Link) (err error) {
	_, err = lr.DB.Exec(
		`
			INSERT INTO link (
				uid_link,
				url,
				name,
				description,
				use_markdown,
				created_at,
				updated_at
			) VALUES (
				?, ?, ?, ?, ?, ?, ?
			)
		`,
		l.GetUid(),
		l.Url,
		l.Name,
		l.Description.Content,
		l.Description.UseMarkdown,
		l.CreatedAt,
		l.UpdatedAt)
		
	lerror.Internal(&err)
	return
}

func (lr *Link) Delete(uid uuid.UUID) (err error) {
	_, err = lr.DB.Exec(
		`DELETE FROM link WHERE uid_link = ?`,
		uid)
		
	lerror.Internal(&err)
	return
}

func (lr *Link) Update(uid uuid.UUID, l domain.Link) (err error) {
	_, err = lr.DB.Exec(
		`
			UPDATE link
			SET url = ?,
			name = ?,
			description = ?,
			use_markdown = ?,
			created_at = ?,
			updated_at = ?
			WHERE uid_link = ?
		`,
		l.Url,
		l.Name,
		l.Description.Content,
		l.Description.UseMarkdown,
		l.CreatedAt,
		l.UpdatedAt,
		uid)
		
	lerror.Internal(&err)
	return
}
