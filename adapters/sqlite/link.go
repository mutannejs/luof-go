package sqlite

import (
	"database/sql"
	"errors"

	"github.com/mutannejs/luof-go/core/domain"

	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

type Link struct {
    DB *sql.DB
}

func New(db *sql.DB) Link {
    return Link{db}
}

func (lr *Link) Exists(uid uuid.UUID) (exists bool, err error) {
    err = lr.DB.QueryRow(
        `SELECT name FROM link WHERE uid_link = ?`,
        uid).Scan(new(string))

    exists = !errors.Is(err, sql.ErrNoRows)

    if errors.Is(err, sql.ErrNoRows) {
        err = nil
    }

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

    return
}

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

    return
}

func (lr *Link) Delete(uid uuid.UUID) (err error) {
    _, err = lr.DB.Exec(
        `DELETE FROM link WHERE uid_link = ?`,
        uid)

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

    return
}
