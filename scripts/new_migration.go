package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mutannejs/luof-go/pkg/lpath"
)

func main() {
    if len(os.Args) == 1 {
        fmt.Println("Erro: Passe ao menos um argumento à função (nome de uma migration)")
        return
    }

    var migrationsPathBase string = lpath.GetAbsolutetPath("migrations")
    var databases []string

    files, err := os.ReadDir(migrationsPathBase)
    if err != nil {
        fmt.Println("Erro: Não foi possível acessar os arquivos em migrations/")
        return
    }

    for _, file := range files {
        if file.IsDir() {
            databases = append(databases, file.Name())
        }
    }

    var migration string
    for _, arg := range os.Args[1:] {
        for _, db := range databases {
            for _, typ := range []string{"up", "down"} {
                migration = fmt.Sprint(time.Now().Unix(), "_", arg, ".", typ, ".sql")

                if _, err = os.OpenFile(
                    lpath.Join(migrationsPathBase, db, migration),
                    os.O_CREATE,
                    0644,
                ); err != nil {
                    fmt.Println("Erro: Não foi possível criar a migration ", migration)
                    break
                }

                fmt.Println("Migration criada: ", lpath.Join(db, migration))
            }
        }
    }
}
