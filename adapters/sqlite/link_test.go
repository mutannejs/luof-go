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
    MockLink, _ = domain.NewLink(
        "github.com/mutannejs/luof",
        "luof",
        "luof repository",
        false,
    )
    MockUidLink = MockLink.GetUid()
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
    lmigration.Up(db, GetMigration)
    lr := New(db)

    ts.env = env
    ts.db = db
    ts.lr = lr
}

func (ts *TestSuite) TearDownSuite() {
    lmigration.Down(ts.db, GetMigration)
}

func (ts *TestSuite) TestCreate() {
    err := ts.lr.Create(MockLink)

    ts.NoError(err, "Tentar criar um link válido não deveria retornar erro")
}

func (ts *TestSuite) TestGetByUid_Exists() {
    link, err := ts.lr.GetByUid(MockUidLink)

    ts.NoError(err, "Tentar recuperar um link informando um uid válido não deveria retornar erro")

    ts.Equal(link.GetUid(), MockUidLink)
    ts.Equal(link.Url, MockLink.Url)
    ts.Equal(link.Name, MockLink.Name)
    ts.Equal(link.Description.Content, MockLink.Description.Content)
    ts.Equal(link.Description.UseMarkdown, MockLink.Description.UseMarkdown)
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

func TestSqliteLink(t *testing.T) {
    suite.Run(t, new(TestSuite))
}
