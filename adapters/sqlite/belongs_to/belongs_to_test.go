package belongs_to

import (
	"database/sql"
	"testing"
	"time"

	"github.com/mutannejs/luof-go/adapters/sqlite"
	"github.com/mutannejs/luof-go/adapters/sqlite/category"
	"github.com/mutannejs/luof-go/adapters/sqlite/link"
	"github.com/mutannejs/luof-go/core/domain"
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
	alternativeMockLink = domain.AlternativeMockLink
	alternativeMockUidLink = domain.AlternativeMockUidLink
)

type TestSuite struct {
	suite.Suite
	env map[string]string
	db *sql.DB
	btr *BelongsTo
}

func (ts *TestSuite) SetupSuite() {
	env, _ := lenv.LoadTest()
	db, _ := sqlite.GetConnection(env)
	lmigration.Migrate(db, belongsToTableMigration, sqlite.GetMigration)
	lr := link.New(db)
	cr := category.New(db)
	btr := New(db)

	ts.env = env
	ts.db = db
	ts.btr = btr

	lr.Create(mockLink)
	lr.Create(alternativeMockLink)
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

	ts.NoError(err, "Tentar relacionar um link a uma categoria, ambos válidos, não deveria retornar erro")
}

func (ts *TestSuite) TestExists() {
	ts.btr.Create(mockUidLink, mockUidCategory, time.Now(), false)

	exists, err := ts.btr.Exists(mockUidLink, mockUidCategory)

	ts.NoError(err, "Exists, se informado uma chave válida não deveria retornar erro")
	ts.Equal(true, exists, "Exists deveria retornar verdadeiro para uma chave válida")
}

func (ts *TestSuite) TestNotExists() {
	exists, err := ts.btr.Exists(mockUidLink, mockUidCategory)

	ts.NoError(err, "Exists, se informado uma chave inválida não deveria retornar erro")
	ts.Equal(false, exists, "Exists deveria retornar falso para uma chave inválida")
}

func (ts *TestSuite) TestGetLinksByCategory_Empty() {
	links, err := ts.btr.GetLinksByCategory(mockUidCategory)

	ts.Empty(links, "Tentar recuperar links de uma categoria vazia deveria retornar uma lista de links vazia")
	ts.NoError(err, "Tentar recuperar links de uma categoria vazia não deveria retornar erro")
}

func (ts *TestSuite) TestGetLinksByCategory_Exists() {
	ts.btr.Create(mockUidLink, mockUidCategory, time.Now(), false)
	ts.btr.Create(alternativeMockUidLink, mockUidCategory, time.Now(), true)
	links, err := ts.btr.GetLinksByCategory(mockUidCategory)

	ts.NoError(err, "Tentar recuperar links de uma categoria válida não deveria retornar erro")
	ts.Len(links, 2)

	var link domain.Link
	if links[0].GetUid() == mockUidLink {
		link = links[0]
	} else if len(links) > 1 && links[1].GetUid() == mockUidLink {
		link = links[1]
	} else {
		ts.Fail("Tentar recuperar links de uma categoria válida deveria retornar todos os dados de seus links")
	}

	ts.Equal(mockUidLink, link.GetUid())
	ts.Equal(mockLink.Name, link.Name)
	ts.Equal(mockLink.Url, link.Url)
	ts.Equal(mockLink.Description.Content, link.Description.Content)
	ts.Equal(mockLink.Description.UseMarkdown, link.Description.UseMarkdown)
	ts.NotZero(link.CreatedAt)
	ts.Zero(link.UpdatedAt)
}

func TestSqliteCategory(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
