package category

import (
	"database/sql"
)

type Category struct {
	DB *sql.DB
}

func New(db *sql.DB) *Category {
	return &Category{db}
}
