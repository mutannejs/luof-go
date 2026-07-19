package link

import (
	"database/sql"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
)

func (lr *Link) Exists(uid uuid.UUID) (exists bool, err error) {
	err = lr.DB.QueryRow(
		`SELECT name FROM link WHERE uid_link = ?`,
		uid).Scan(new(string))

	exists = err != sql.ErrNoRows

	if err == sql.ErrNoRows {
		err = nil
	}
	
	lerror.SetInternal(&err)
	return
}

func (lr *Link) GetByUid(uid uuid.UUID) (l domain.Link, err error) {
	err = lr.DB.QueryRow(
			`SELECT
				url,
				name,
				description,
				use_markdown,
				created_at,
				updated_at
			FROM link WHERE uid_link = ?`,
			uid).
		Scan(
			&l.Url,
			&l.Name,
			&l.Description.Content,
			&l.Description.UseMarkdown,
			&l.CreatedAt,
			&l.UpdatedAt)

	if err == nil {
		l.SetUid(uid)
	}

	lerror.SetInternal(&err)
	return
}
