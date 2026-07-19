package belongs_to

import (
	"database/sql"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
)

func (btr *BelongsTo) Exists(
	linkUid uuid.UUID,
	categoryUid uuid.UUID,
) (exists bool, err error) {
	err = btr.DB.QueryRow(
		`
			SELECT uid_category
			FROM belongs_to
			WHERE uid_category = ? AND uid_link = ?
		`,
		categoryUid,
		linkUid).Scan(new(string))

	exists = err != sql.ErrNoRows

	if err == sql.ErrNoRows {
		err = nil
	}

	lerror.SetInternal(&err)
	return
}

func (btr *BelongsTo) GetLinksByCategory(uid uuid.UUID) (links []domain.Link, err error) {
	var rows *sql.Rows

	rows, err = btr.DB.Query(
		`
			SELECT
				l.uid_link,
				l.name,
				l.url,
				l.description,
				l.use_markdown,
				l.created_at,
				l.updated_at
			FROM belongs_to bt
			INNER JOIN link l ON l.uid_link = bt.uid_link
			WHERE uid_category = ?
		`,
		uid)

	if err != nil {
		lerror.SetInternal(&err)
		return
	}

	links = make([]domain.Link, 0)
	var l domain.Link
	var uidLinkStr string
	var uidLink uuid.UUID
	var errLoop error

	for rows.Next() {
		errLoop = rows.Scan(
			&uidLinkStr,
			&l.Name,
			&l.Url,
			&l.Description.Content,
			&l.Description.UseMarkdown,
			&l.CreatedAt,
			&l.UpdatedAt)

		if errLoop != nil {
			continue
		}

		uidLink, errLoop = uuid.Parse(uidLinkStr)

		if errLoop != nil {
			continue
		}

		l.SetUid(uidLink)
		links = append(links, l)
	}

	err = rows.Err()
	rows.Close()

	lerror.SetInternal(&err)
	return
}

func (btr *BelongsTo) HasLinks(
	uid uuid.UUID,
) (hasLinks bool, err error) {
	err = btr.DB.QueryRow(
		`SELECT 1 FROM belongs_to WHERE uid_category = ? LIMIT 1`,
		uid).Scan(new(int))

	hasLinks = err != sql.ErrNoRows

	if err == sql.ErrNoRows {
		err = nil
	}

	lerror.SetInternal(&err)
	return
}
