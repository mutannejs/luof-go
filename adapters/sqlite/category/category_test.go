package category

import (
	"database/sql"
	"testing"
	"time"

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
	categoriesTree = domain.CategoriesTree
)

type TestSuite struct {
	suite.Suite
	env map[string]string
	db *sql.DB
	cr *Category
}

func (ts *TestSuite) SetupSuite() {
	env, _ := lenv.LoadTest()
	db, _ := sqlite.GetConnection(env)
	lmigration.Migrate(db, categoryTableMigration, sqlite.GetMigration)
	cr := New(db)

	ts.env = env
	ts.db = db
	ts.cr = cr

	for _, category := range categoriesTree {
		cr.Create(category)
	}
}

func (ts *TestSuite) TearDownTest() {
	ts.db.Exec("DELETE FROM category")
}

func (ts *TestSuite) TearDownSuite() {
	lmigration.Drop(ts.db, sqlite.GetMigration)
	ts.db.Close()
}

func (ts *TestSuite) TestCreate() {
	err := ts.cr.Create(mockCategory)

	ts.NoError(err, "Tentar criar uma categoria válida não deveria retornar erro")
}

func (ts *TestSuite) TestGetByUid_Exists() {
	ts.cr.Create(mockCategory)

	category, err := ts.cr.GetByUid(mockUidCategory)

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

	category, err := ts.cr.GetByUid(uid)

	ts.Empty(category, "Tentar recuperar uma categoria informando um uid inválido deveria retornar um Category vazio")
	ts.ErrorIs(sql.ErrNoRows, err, "Tentar recuperar uma categoria informando um uid inválido deveria retornar " + sql.ErrNoRows.Error())
}

func (ts *TestSuite) TestExists() {
	ts.cr.Create(mockCategory)

	exists, err := ts.cr.Exists(mockUidCategory)

	ts.NoError(err, "Exists se informado um uid válido não deveria retornar erro")
	ts.Equal(true, exists, "Exists deveria retornar verdadeiro para um uid válido")
}

func (ts *TestSuite) TestNotExists() {
	uid, err := luuid.New()

	if err != nil {
		ts.Fail(err.Error())
	}

	exists, err := ts.cr.Exists(uid)

	ts.NoError(err, "Exists se informado um uid inválido não deveria retornar erro")
	ts.Equal(false, exists, "Exists deveria retornar falso para um uid válido")
}

func (ts *TestSuite) TestUpdate() {
	ts.cr.Create(mockCategory)

	err := ts.cr.Update(mockUidCategory, alternativeMockCategory)

	ts.NoError(err, "Tentar atualizar uma categoria com uid válido não deveria retornar erro")

	category, _ := ts.cr.GetByUid(mockUidCategory)

	ts.Equal(alternativeMockCategory.Name, category.Name)
	ts.Equal(alternativeMockCategory.Description.Content, category.Description.Content)
	ts.Equal(alternativeMockCategory.Description.UseMarkdown, category.Description.UseMarkdown)
}

func (ts *TestSuite) TestDelete() {
	ts.cr.Create(mockCategory)

	err := ts.cr.Delete(mockUidCategory)

	ts.NoError(err, "Tentar deletar uma categoria válida não deveria retornar erro")

	_, err = ts.cr.GetByUid(mockUidCategory)

	ts.ErrorIs(sql.ErrNoRows, err, "Tentar recuperar uma categoria previamente deletado deveria retornar " + sql.ErrNoRows.Error())
}

func (ts *TestSuite) TestIsSubcategory() {
	ts.makeTree()

	areRelated, err := ts.cr.IsSubcategory(
		categoriesTree["filme"].GetUid(),
		categoriesTree["terror"].GetUid())

	ts.NoError(err, "IsSubcategory, se informado duas chaves válidas não deveria retornar erro")
	ts.Equal(true, areRelated, "IsSubcategory deveria retornar verdadeiro para duas categorias que são uma subcategoria direta da outra")
}

func (ts *TestSuite) TestNotIsSubcategory() {
	ts.makeTree()

	areRelated, err := ts.cr.IsSubcategory(
		categoriesTree["filme"].GetUid(),
		categoriesTree["serial_killer"].GetUid())

	ts.NoError(err, "IsSubcategory, se informado duas chaves válidas não deveria retornar erro")
	ts.Equal(false, areRelated, "IsSubcategory deveria retornar falso para duas categorias que não são uma subcategoria direta da outra")
}

/*func (ts *TestSuite) TestGetLinksByCategory_Empty() {
	links, err := ts.cr.GetLinksByCategory(mockUidCategory)

	ts.Empty(links, "Tentar recuperar links de uma categoria vazia deveria retornar uma lista de links vazia")
	ts.NoError(err, "Tentar recuperar links de uma categoria vazia não deveria retornar erro")
}

func (ts *TestSuite) TestGetLinksByCategory_Exists() {
	ts.cr.Create(mockUidLink, mockUidCategory, time.Now(), false)
	ts.cr.Create(alternativeMockUidLink, mockUidCategory, time.Now(), true)
	links, err := ts.cr.GetLinksByCategory(mockUidCategory)

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
}*/

func TestSqliteCategory(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (ts *TestSuite) makeTree() {
	ts.insertSubcategory("leitura", "livro")
	ts.insertSubcategory("filme", "terror")
	ts.insertSubcategory("filme", "acao")
	ts.insertSubcategory("terror", "jumpscare")
	ts.insertSubcategory("terror", "serial_killer")
}

func (ts *TestSuite) insertSubcategory(father, child string) error {
	return ts.cr.InsertSubcategory(
		categoriesTree[father].GetUid(),
		categoriesTree[child].GetUid(),
		time.Now())
}
