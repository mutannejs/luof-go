package link_storage

import (
    "fmt"
    "database/sql"
)

type LinkStorage struct {
    DB *sql.DB
}

func NewLinkStorage (db *sql.DB) (LinkStorage, error) {
    ls := LinkStorage{db}
    err := ls.InitTable()
    return ls, err
}

func (ls *LinkStorage) InitTable() error {
    _, err := ls.DB.Exec(`
        CREATE TABLE IF NOT EXISTS link (
            id INTEGER PRIMARY KEY,
            url TEXT,
            name VARCHAR(200) NOT NULL,
            description TEXT,
            created_at DATETIME NOT NULL,
            update_at DATETIME NOT NULL
        );
    `)
    if err != nil {
        fmt.Println("SQL >>", err)
    }
    return err
}
