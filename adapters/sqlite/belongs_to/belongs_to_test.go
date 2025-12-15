package belongs_to

import (
	"database/sql"
	"testing"
	"time"

	"github.com/mutannejs/luof-go/adapters/sqlite"
	"github.com/mutannejs/luof-go/adapters/sqlite/category"
	"github.com/mutannejs/luof-go/adapters/sqlite/link"
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/core/repository"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/lmigration"

	"github.com/stretchr/testify/suite"
)

var (
    belongsToTableMigration uint = 1765827457
    mockLink = domain.MockLink
    mockUidLink = domain.MockUidLink
    mockCategory = domain.MockCategory
    mockUidCategory = domain.MockUidCategory
    alternativeMockUidLink = domain.AlternativeMockUidLink
    alternativeMockUidCategory = domain.AlternativeMockUidCategory
)

type TestSuite struct {
    suite.Suite
    env map[string]string
    db *sql.DB
    btr BelongsTo
}

func (ts *TestSuite) SetupSuite() {
    env, _ := lenv.Load(true)
    db, _ := sqlite.GetConnection(env)
    lmigration.Migrate(db, belongsToTableMigration, sqlite.GetMigration)
    lr := link.New(db)
    cr := category.New(db)
    btr := New(db)

    ts.env = env
    ts.db = db
    ts.btr = btr

    lr.Create(mockLink)
    cr.Create(mockCategory)
}

func (ts *TestSuite) TearDownTest() {
    ts.db.Exec("DELETE FROM belongs_to")
}

func (ts *TestSuite) TearDownSuite() {
    lmigration.Drop(ts.db, sqlite.GetMigration)
    ts.db.Close()
}

func (ts *TestSuite) TestCreate() {
    err := ts.btr.Create(mockUidLink, mockUidCategory, time.Now(), false)

    ts.NoError(err, "Tentar relacionar um link a uma categoria, ainda não relacionados e ambos válidos, não deveria retornar erro")
}

func (ts *TestSuite) TestCreate_Exists() {
    ts.btr.Create(mockUidLink, mockUidCategory, time.Now(), false)
    err := ts.btr.Create(mockUidLink, mockUidCategory, time.Now(), true)

    ts.EqualError(err, repository.ALREADY_BELONGS, "Tentar relacionar um link já pertencente a uma categoria novamente, deveria retornar o erro " + repository.ALREADY_BELONGS)
}

func (ts *TestSuite) TestCreate_Invalid() {
    err := ts.btr.Create(alternativeMockUidLink, alternativeMockUidCategory, time.Now(), false)

    ts.Error(err, "Tentar relacionar um link a uma categoria, sendo um deles inválido, deveria retornar erro")
}

// func (ts *TestSuite) TestExists() {
//     ts.btr.Create(mockUidLink, mockUidCategory, time.Now(), false)

//     exists, err := ts.btr.Exists(mockUidLink, mockUidCategory)

//     ts.NoError(err, "Exists se informado um uid válido não deveria retornar erro")
//     ts.Equal(true, exists, "Exists deveria retornar verdadeiro para um uid válido")
// }

// func (ts *TestSuite) TestNotExists() {
//     uid, err := luuid.New()

//     if err != nil {
//         ts.Fail(err.Error())
//     }

//     exists, err := ts.btr.Exists(uid)

//     ts.NoError(err, "Exists se informado um uid inválido não deveria retornar erro")
//     ts.Equal(false, exists, "Exists deveria retornar falso para um uid válido")
// }

// func (ts *TestSuite) TestGetByUid_NotExists() {
//     uid, err := luuid.New()

//     if err != nil {
//         ts.Fail(err.Error())
//     }

//     category, err := ts.lr.GetByUid(uid)

//     ts.Empty(category, "Tentar recuperar uma categoria informando um uid inválido deveria retornar um Category vazio")
//     ts.ErrorIs(sql.ErrNoRows, err, "Tentar recuperar uma categoria informando um uid inválido deveria retornar " + sql.ErrNoRows.Error())
// }

// func (ts *TestSuite) TestGetByUid_Exists() {
//     category, err := ts.btr.GetByUid(mockUidCategory)

//     ts.NoError(err, "Tentar recuperar uma categoria informando um uid válido não deveria retornar erro")
//     ts.Equal(mockUidCategory, category.GetUid())
//     ts.Equal(mockCategory.Name, category.Name)
//     ts.Equal(mockCategory.Description.Content, category.Description.Content)
//     ts.Equal(mockCategory.Description.UseMarkdown, category.Description.UseMarkdown)
//     ts.NotZero(category.CreatedAt)
//     ts.Zero(category.UpdatedAt)
// }

func TestSqliteCategory(t *testing.T) {
    suite.Run(t, new(TestSuite))
}
