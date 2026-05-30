package category

import (
	"database/sql"

	"github.com/mutannejs/luof-go/core/domain"

	"github.com/google/uuid"
)

func (cr *Category) AreRelated(
	firstCategoryUid uuid.UUID,
	secondCategoryUid uuid.UUID,
) (areRelated bool, err error) {
	err = cr.DB.QueryRow(
		`
			SELECT father FROM (
				WITH RECURSIVE descendants
					(father, child)
				AS (
					SELECT uid_father, uid_category
					FROM category
					WHERE uid_father = ?

					UNION

					SELECT c.uid_father, c.uid_category
					FROM category c
					JOIN descendants d ON d.child = c.uid_father
				)
				SELECT *
				FROM descendants
				WHERE father is not null

				UNION

				SELECT * FROM (
					WITH RECURSIVE ancestors
						(father, child)
					AS (
						SELECT uid_father, uid_category
						FROM category
						WHERE uid_category = ?

						UNION

						SELECT c.uid_father, c.uid_category
						FROM category c
						JOIN ancestors d ON d.father = c.uid_category
					)
					SELECT father AS uid_father, child AS uid_category
					FROM ancestors
					WHERE father is not null
				)
			)
			WHERE father = ? OR child = ?
		`,
		firstCategoryUid,
		firstCategoryUid,
		secondCategoryUid,
		secondCategoryUid,
	).Scan(new(string))

	areRelated = err != sql.ErrNoRows

	if err == sql.ErrNoRows {
		err = nil
	}

	return
}

func (cr *Category) Exists(uid uuid.UUID) (exists bool, err error) {
	err = cr.DB.QueryRow(
		`SELECT name FROM category WHERE uid_category = ?`,
		uid).Scan(new(string))

	exists = err != sql.ErrNoRows

	if err == sql.ErrNoRows {
		err = nil
	}

	return
}

func (cr *Category) GetByUid(uid uuid.UUID) (l domain.Category, err error) {
	err = cr.DB.QueryRow(
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

func (cr *Category) GetAllRootCategories() (
	categories []domain.Category, err error,
) {
	var rows *sql.Rows

	rows, err = cr.DB.Query(
		`
			SELECT
				uid_category,
				name,
				description,
				use_markdown,
				created_at,
				updated_at
			FROM category
			WHERE uid_father IS NULL
		`)

	if err != nil {
		return
	}

	categories = make([]domain.Category, 0, 0)
	var c domain.Category
	var categoryUid uuid.UUID
	var categoryUidStr string
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

func (cr *Category) GetSubcategories(
	uid uuid.UUID,
) (categories []domain.Category, err error) {
	var rows *sql.Rows

	rows, err = cr.DB.Query(
		`
			SELECT
				uid_category,
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

	categories = make([]domain.Category, 0, 0)
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

func (cr *Category) HasSubcategories(
	uid uuid.UUID,
) (hasSubcategories bool, err error) {
	err = cr.DB.QueryRow(
		`SELECT 1 FROM category WHERE uid_father = ? LIMIT 1`,
		uid).Scan(new(int))

	hasSubcategories = err != sql.ErrNoRows

	if err == sql.ErrNoRows {
		err = nil
	}

	return
}

func (cr *Category) IsAncestor(
	ancestorUid uuid.UUID,
	categoryUid uuid.UUID,
) (isAncestor bool, err error) {
	err = cr.DB.QueryRow(
		`
			WITH RECURSIVE ancestors
				(father, child)
			AS (
				SELECT uid_father, uid_category
				FROM category
				WHERE uid_category = ?

				UNION

				SELECT c.uid_father, c.uid_category
				FROM category c
				JOIN ancestors d ON d.father = c.uid_category
			)
			SELECT father
			FROM ancestors
			WHERE father = ?
		`,
		categoryUid,
		ancestorUid,
	).Scan(new(string))

	isAncestor = err != sql.ErrNoRows

	if err == sql.ErrNoRows {
		err = nil
	}

	return
}

func (cr *Category) IsSubcategory(
	fatherUid uuid.UUID,
	childUid uuid.UUID,
) (isSubcategory bool, err error) {
	err = cr.DB.QueryRow(
		`
			SELECT name
			FROM category
			WHERE uid_father = ? AND uid_category = ?
		`,
		fatherUid,
		childUid,
	).Scan(new(string))

	isSubcategory = err != sql.ErrNoRows

	if err == sql.ErrNoRows {
		err = nil
	}

	return
}
