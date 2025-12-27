package category

import (
	"database/sql"

	"github.com/mutannejs/luof-go/core/domain"

	"github.com/google/uuid"
)

type Category struct {
    DB *sql.DB
}

func New(db *sql.DB) *Category {
    return &Category{db}
}

func (lr *Category) Exists(uid uuid.UUID) (exists bool, err error) {
    err = lr.DB.QueryRow(
        `SELECT name FROM category WHERE uid_category = ?`,
        uid).Scan(new(string))

    exists = err != sql.ErrNoRows

    if err == sql.ErrNoRows {
        err = nil
    }

    return
}

func (lr *Category) GetByUid(uid uuid.UUID) (l domain.Category, err error) {
    err = lr.DB.QueryRow(
            `SELECT
                name,
                description,
                use_markdown,
                created_at,
                updated_at
            FROM category WHERE uid_category = ?`,
            uid).
        Scan(
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

func (lr *Category) Create(l domain.Category) (err error) {
    _, err = lr.DB.Exec(
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

func (lr *Category) Delete(uid uuid.UUID) (err error) {
    _, err = lr.DB.Exec(
        `DELETE FROM category WHERE uid_category = ?`,
        uid)

    return
}

func (lr *Category) Update(uid uuid.UUID, l domain.Category) (err error) {
    _, err = lr.DB.Exec(
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
