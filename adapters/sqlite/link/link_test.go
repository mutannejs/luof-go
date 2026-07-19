package link

import (
	"database/sql"
	"testing"

	"github.com/mutannejs/luof-go/adapters/sqlite"
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/lmigration"

	"github.com/stretchr/testify/suite"
)

const (
	linkTableMigration = lmigration.LINK_TABLE_MIGRATION
)

var (
	mockLink = domain.MockLink
	mockUidLink = domain.MockUidLink
	alternativeMockLink = domain.AlternativeMockLink
)

type TestSuite struct {
	suite.Suite
	env map[string]string
	db *sql.DB
	lr *Link
}

func (ts *TestSuite) SetupSuite() {
	env, _ := lenv.LoadTest()
	db, _ := sqlite.GetConnection(env)
	lmigration.Migrate(db, linkTableMigration, sqlite.GetMigration)
	lr := New(db)

	ts.env = env
	ts.db = db
	ts.lr = lr
}

func (ts *TestSuite) TearDownTest() {
	ts.db.Exec("DELETE FROM link")
}

func (ts *TestSuite) TearDownSuite() {
	lmigration.Drop(ts.db, sqlite.GetMigration)
	ts.db.Close()
}

func (ts *TestSuite) TestCreate() {
	err := ts.lr.Create(mockLink)

	ts.NoError(err, "Tentar criar um link válido não deveria retornar erro")
}

func (ts *TestSuite) TestGetByUid_Exists() {
	ts.lr.Create(mockLink)

	link, err := ts.lr.GetByUid(mockUidLink)

	ts.NoError(err, "Tentar recuperar um link informando um uid válido não deveria retornar erro")
	ts.Equal(mockUidLink, link.GetUid())
	ts.Equal(mockLink.Url, link.Url)
	ts.Equal(mockLink.Name, link.Name)
	ts.Equal(mockLink.Description.Content, link.Description.Content)
	ts.Equal(mockLink.Description.UseMarkdown, link.Description.UseMarkdown)
	ts.NotZero(link.CreatedAt)
	ts.Zero(link.UpdatedAt)
}

func (ts *TestSuite) TestGetByUid_NotExists() {
	uid := mockUidLink

	link, err := ts.lr.GetByUid(uid)

	ts.Empty(link, "Tentar recuperar um link informando um uid inválido deveria retornar um Link vazio")
	ts.ErrorIs(sql.ErrNoRows, err, "Tentar recuperar um link informando um uid inválido deveria retornar " + sql.ErrNoRows.Error())
}

func (ts *TestSuite) TestExists() {
	ts.lr.Create(mockLink)

	exists, err := ts.lr.Exists(mockUidLink)

	ts.NoError(err, "Exists se informado um uid válido não deveria retornar erro")
	ts.True( exists, "Exists deveria retornar verdadeiro para um uid válido")
}

func (ts *TestSuite) TestNotExists() {
	uid := mockUidLink

	exists, err := ts.lr.Exists(uid)

	ts.NoError(err, "Exists se informado um uid inválido não deveria retornar erro")
	ts.False( exists, "Exists deveria retornar falso para um uid válido")
}

func (ts *TestSuite) TestUpdate() {
	ts.lr.Create(mockLink)

	err := ts.lr.Update(mockUidLink, alternativeMockLink)

	ts.NoError(err, "Tentar atualizar um link com uid válido não deveria retornar erro")

	link, _ := ts.lr.GetByUid(mockUidLink)

	ts.Equal(alternativeMockLink.Url, link.Url)
	ts.Equal(alternativeMockLink.Name, link.Name)
	ts.Equal(alternativeMockLink.Description.Content, link.Description.Content)
	ts.Equal(alternativeMockLink.Description.UseMarkdown, link.Description.UseMarkdown)
}

func (ts *TestSuite) TestDelete() {
	ts.lr.Create(mockLink)

	err := ts.lr.Delete(mockUidLink)

	ts.NoError(err, "Tentar deletar um link válido não deveria retornar erro")

	_, err = ts.lr.GetByUid(mockUidLink)

	ts.ErrorIs(sql.ErrNoRows, err, "Tentar recuperar um link previamente deletado deveria retornar " + sql.ErrNoRows.Error())
}

func TestSqliteLink(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
