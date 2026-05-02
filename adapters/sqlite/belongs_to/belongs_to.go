package belongs_to

import (
	"database/sql"
	"time"

	"github.com/mutannejs/luof-go/core/domain"

	"github.com/google/uuid"
)

type BelongsTo struct {
	DB *sql.DB
}

func New(db *sql.DB) *BelongsTo {
	return &BelongsTo{db}
}

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
		return
	}

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

	if len(links) == 0 {
		links = make([]domain.Link, 0, 0)
	}
	
	err = rows.Err()
	rows.Close()

	return
}

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
			WHERE uid_link = ? AND uid_Category = ?
		`,
		isMain,
		linkUid,
		categoryUid)

	return
}
