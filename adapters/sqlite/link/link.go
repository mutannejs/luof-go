package link

import (
	"database/sql"
)

type Link struct {
	DB *sql.DB
}

func New(db *sql.DB) *Link {
	return &Link{db}
}
