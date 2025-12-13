package sqlite

import (
	"database/sql"
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/lmigration"
	"github.com/mutannejs/luof-go/pkg/luuid"

	"github.com/stretchr/testify/suite"
)

var (
    MockLink = domain.MockLink
    MockUidLink = domain.MockUidLink
    AlternativeMockLink = domain.AlternativeMockLink
)

type TestSuite struct {
    suite.Suite
    env map[string]string
    db *sql.DB
    lr Link
}

func (ts *TestSuite) SetupSuite() {
    env, _ := lenv.Load(true)
    db, _ := GetConnection(env)
    lmigration.Migrate(db, 1764809880, GetMigration)
    lr := New(db)

    ts.env = env
    ts.db = db
    ts.lr = lr
}

func (ts *TestSuite) TearDownTest() {
    ts.db.Exec("DELETE FROM link")
}

func (ts *TestSuite) TearDownSuite() {
    lmigration.Down(ts.db, GetMigration)
}

func (ts *TestSuite) TestCreate() {
    err := ts.lr.Create(MockLink)

    ts.NoError(err, "Tentar criar um link válido não deveria retornar erro")
}

func (ts *TestSuite) TestGetByUid_Exists() {
    ts.lr.Create(MockLink)

    link, err := ts.lr.GetByUid(MockUidLink)

    ts.NoError(err, "Tentar recuperar um link informando um uid válido não deveria retornar erro")
    ts.Equal(MockUidLink, link.GetUid())
    ts.Equal(MockLink.Url, link.Url)
    ts.Equal(MockLink.Name, link.Name)
    ts.Equal(MockLink.Description.Content, link.Description.Content)
    ts.Equal(MockLink.Description.UseMarkdown, link.Description.UseMarkdown)
    ts.NotZero(link.CreatedAt)
    ts.Zero(link.UpdatedAt)
}

func (ts *TestSuite) TestGetByUid_NotExists() {
    uid, err := luuid.New()

    if err != nil {
        ts.Fail(err.Error())
    }

    link, err := ts.lr.GetByUid(uid)

    ts.Empty(link, "Tentar recuperar um link informando um uid inválido deveria retornar um Link vazio")
    ts.ErrorIs(sql.ErrNoRows, err, "Tentar recuperar um link informando um uid inválido deveria retornar " + sql.ErrNoRows.Error())
}

func (ts *TestSuite) TestExists() {
    ts.lr.Create(MockLink)

    exists, err := ts.lr.Exists(MockUidLink)

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
    ts.lr.Create(MockLink)

    err := ts.lr.Update(MockUidLink, AlternativeMockLink)

    ts.NoError(err, "Tentar atualizar um link com uid válido não deveria retornar erro")

    link, _ := ts.lr.GetByUid(MockUidLink)

    ts.Equal(AlternativeMockLink.Url, link.Url)
    ts.Equal(AlternativeMockLink.Name, link.Name)
    ts.Equal(AlternativeMockLink.Description.Content, link.Description.Content)
    ts.Equal(AlternativeMockLink.Description.UseMarkdown, link.Description.UseMarkdown)
}

func (ts *TestSuite) TestDelete() {
    ts.lr.Create(MockLink)

    err := ts.lr.Delete(MockUidLink)

    ts.NoError(err, "Tentar deletar um link válido não deveria retornar erro")

    _, err = ts.lr.GetByUid(MockUidLink)

    ts.ErrorIs(sql.ErrNoRows, err, "Tentar recuperar um link previamente deletado deveria retornar " + sql.ErrNoRows.Error())
}

func TestSqliteLink(t *testing.T) {
    suite.Run(t, new(TestSuite))
}
