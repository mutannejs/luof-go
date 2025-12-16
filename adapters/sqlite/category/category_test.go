package category

import (
	"database/sql"
	"testing"

	"github.com/mutannejs/luof-go/adapters/sqlite"
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/lmigration"
	"github.com/mutannejs/luof-go/pkg/luuid"

	"github.com/stretchr/testify/suite"
)

var (
    categoryTableMigration uint = 1765719599
    mockCategory = domain.MockCategory
    mockUidCategory = domain.MockUidCategory
    alternativeMockCategory = domain.AlternativeMockCategory
)

type TestSuite struct {
    suite.Suite
    env map[string]string
    db *sql.DB
    lr *Category
}

func (ts *TestSuite) SetupSuite() {
    env, _ := lenv.Load(true)
    db, _ := sqlite.GetConnection(env)
    lmigration.Migrate(db, categoryTableMigration, sqlite.GetMigration)
    lr := New(db)

    ts.env = env
    ts.db = db
    ts.lr = lr
}

func (ts *TestSuite) TearDownTest() {
    ts.db.Exec("DELETE FROM category")
}

func (ts *TestSuite) TearDownSuite() {
    lmigration.Drop(ts.db, sqlite.GetMigration)
    ts.db.Close()
}

func (ts *TestSuite) TestCreate() {
    err := ts.lr.Create(mockCategory)

    ts.NoError(err, "Tentar criar uma categoria válida não deveria retornar erro")
}

func (ts *TestSuite) TestGetByUid_Exists() {
    ts.lr.Create(mockCategory)

    category, err := ts.lr.GetByUid(mockUidCategory)

    ts.NoError(err, "Tentar recuperar uma categoria informando um uid válido não deveria retornar erro")
    ts.Equal(mockUidCategory, category.GetUid())
    ts.Equal(mockCategory.Name, category.Name)
    ts.Equal(mockCategory.Description.Content, category.Description.Content)
    ts.Equal(mockCategory.Description.UseMarkdown, category.Description.UseMarkdown)
    ts.NotZero(category.CreatedAt)
    ts.Zero(category.UpdatedAt)
}

func (ts *TestSuite) TestGetByUid_NotExists() {
    uid, err := luuid.New()

    if err != nil {
        ts.Fail(err.Error())
    }

    category, err := ts.lr.GetByUid(uid)

    ts.Empty(category, "Tentar recuperar uma categoria informando um uid inválido deveria retornar um Category vazio")
    ts.ErrorIs(sql.ErrNoRows, err, "Tentar recuperar uma categoria informando um uid inválido deveria retornar " + sql.ErrNoRows.Error())
}

func (ts *TestSuite) TestExists() {
    ts.lr.Create(mockCategory)

    exists, err := ts.lr.Exists(mockUidCategory)

    ts.NoError(err, "Exists se informado um uid válido não deveria retornar erro")
    ts.Equal(true, exists, "Exists deveria retornar verdadeiro para um uid válido")
}

func (ts *TestSuite) TestNotExists() {
    uid, err := luuid.New()

    if err != nil {
        ts.Fail(err.Error())
    }

    exists, err := ts.lr.Exists(uid)

    ts.NoError(err, "Exists se informado um uid inválido não deveria retornar erro")
    ts.Equal(false, exists, "Exists deveria retornar falso para um uid válido")
}

func (ts *TestSuite) TestUpdate() {
    ts.lr.Create(mockCategory)

    err := ts.lr.Update(mockUidCategory, alternativeMockCategory)

    ts.NoError(err, "Tentar atualizar uma categoria com uid válido não deveria retornar erro")

    category, _ := ts.lr.GetByUid(mockUidCategory)

    ts.Equal(alternativeMockCategory.Name, category.Name)
    ts.Equal(alternativeMockCategory.Description.Content, category.Description.Content)
    ts.Equal(alternativeMockCategory.Description.UseMarkdown, category.Description.UseMarkdown)
}

func (ts *TestSuite) TestDelete() {
    ts.lr.Create(mockCategory)

    err := ts.lr.Delete(mockUidCategory)

    ts.NoError(err, "Tentar deletar uma categoria válida não deveria retornar erro")

    _, err = ts.lr.GetByUid(mockUidCategory)

    ts.ErrorIs(sql.ErrNoRows, err, "Tentar recuperar uma categoria previamente deletado deveria retornar " + sql.ErrNoRows.Error())
}

func TestSqliteCategory(t *testing.T) {
    suite.Run(t, new(TestSuite))
}
