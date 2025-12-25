package get_repositories

import (
	"database/sql"

	"github.com/mutannejs/luof-go/adapters/sqlite/belongs_to"
	"github.com/mutannejs/luof-go/adapters/sqlite/category"
	"github.com/mutannejs/luof-go/adapters/sqlite/link"
	"github.com/mutannejs/luof-go/core/repository"
)

func GetRepositories(db *sql.DB) repository.Repositories {
    return repository.Repositories{
        BelongsTo: belongs_to.New(db),
        Category: category.New(db),
        Link: link.New(db),
    }
}
