package belongs_to

import (
	"database/sql"
)

type BelongsTo struct {
	DB *sql.DB
}

func New(db *sql.DB) *BelongsTo {
	return &BelongsTo{db}
}
