package lmigration

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(db *sql.DB, version uint, getMigration func (*sql.DB) (*migrate.Migrate, error)) (err error) {
    var m *migrate.Migrate

    if m, err = getMigration(db); err != nil {
        return err
    }

    return m.Migrate(version)
}

func Up(db *sql.DB, getMigration func (*sql.DB) (*migrate.Migrate, error)) (err error) {
    var m *migrate.Migrate

    if m, err = getMigration(db); err != nil {
        return err
    }

    return m.Up()
}

func Down(db *sql.DB, getMigration func (*sql.DB) (*migrate.Migrate, error)) (err error) {
    var m *migrate.Migrate

    if m, err = getMigration(db); err != nil {
        return err
    }

    return m.Down()
}
