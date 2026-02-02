package ltests

import (
	"database/sql"
)

func CleanTable(db *sql.DB, table string) (err error) {
	_, err = db.Exec("DELETE FROM " + table)
	return
}
