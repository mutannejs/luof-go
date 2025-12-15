package belongs_to

import (
	"database/sql"
	"errors"
	"time"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lerror"

	"github.com/google/uuid"
)

type BelongsTo struct {
    DB *sql.DB
}

func New(db *sql.DB) BelongsTo {
    return BelongsTo{db}
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

    exists = !lerror.Equals(err, sql.ErrNoRows)

    if lerror.Equals(err, sql.ErrNoRows) {
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
    var exists bool

    if exists, err = btr.Exists(linkUid, categoryUid); exists {
        return errors.New(repository.ALREADY_BELONGS)
    } else if err != nil {
        return
    }

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
