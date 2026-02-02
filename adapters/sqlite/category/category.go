package category

import (
	"database/sql"
	"time"

	"github.com/mutannejs/luof-go/core/domain"

	"github.com/google/uuid"
)

type Category struct {
	DB *sql.DB
}

func New(db *sql.DB) *Category {
	return &Category{db}
}

func (sr *Category) AreRelated(
	fatherUid uuid.UUID,
	childUid uuid.UUID,
) (areRelated bool, err error) {
	err = sr.DB.QueryRow(
		`
			WITH RECURSIVE relateds
				(father, child)
			AS (
				SELECT uid_father, uid_category
				FROM category
				WHERE uid_father = ?

				UNION

				SELECT c.uid_father, c.uid_category
				FROM category c
				INNER JOIN relateds d
				WHERE d.child = c.uid_father

				UNION

				SELECT c.uid_father, c.uid_child
				FROM category c
				INNER JOIN relateds d
				WHERE d.father = c.uid_child
			)

			SELECT father as uid_category
			FROM relateds
			WHERE uid_category = ?

			UNION

			SELECT child as uid_category
			FROM relateds
			WHERE uid_category = ?
		`,
		fatherUid,
		childUid,
		childUid,
	).Scan(new(string))

	areRelated = err != sql.ErrNoRows

	if err == sql.ErrNoRows {
		err = nil
	}

	return
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

func (sr *Category) GetSubcategories(
	uid uuid.UUID,
) (categories []domain.Category, err error) {
	var rows *sql.Rows

	rows, err = sr.DB.Query(
		`
			SELECT
				name,
				description,
				use_markdown,
				created_at,
				updated_at
			FROM category
			WHERE uid_father = ?
		`,
		uid)

	if err != nil {
		return
	}

	var c domain.Category
	var categoryUidStr string
	var categoryUid uuid.UUID
	var loopErr error

	for rows.Next() {
		loopErr = rows.Scan(
			&categoryUidStr,
			&c.Name,
			&c.Description.Content,
			&c.Description.UseMarkdown,
			&c.CreatedAt,
			&c.UpdatedAt)

		if loopErr != nil {
			continue
		}

		categoryUid, loopErr = uuid.Parse(categoryUidStr)

		if loopErr != nil {
			continue
		}

		c.SetUid(categoryUid)
		categories = append(categories, c)
	}
	
	err = rows.Err()
	rows.Close()

	return
}

func (sr *Category) IsSubcategory(
	fatherUid uuid.UUID,
	childUid uuid.UUID,
) (isCategory bool, err error) {
	err = sr.DB.QueryRow(
		`
			SELECT uid_father
			FROM category
			WHERE uid_father = ? AND uid_category = ?
		`,
		fatherUid,
		childUid,
	).Scan(new(string))

	isCategory = err != sql.ErrNoRows

	if err == sql.ErrNoRows {
		err = nil
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

func (sr *Category) DeleteSubcategory(
	childUid uuid.UUID,
) (err error) {
	_, err = sr.DB.Exec(
		`
			UPDATE category
			SET uid_father = null
			WHERE uid_category = ?
		`,
		childUid)

	return
}

func (sr *Category) InsertSubcategory(
	fatherUid uuid.UUID,
	childUid uuid.UUID,
	updatedAt time.Time,
) (err error) {
	_, err = sr.DB.Exec(
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
