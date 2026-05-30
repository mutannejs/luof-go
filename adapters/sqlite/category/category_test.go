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

func (ts *TestSuite) insertAllCategories() {
	for _, category := range categoriesTree {
		ts.cr.Create(category)
	}
}

func (ts *TestSuite) makeTree() {
	ts.insertAllCategories()
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

func (ts *TestSuite) SetupSuite() {
	env, _ := lenv.LoadTest()
	db, _ := sqlite.GetConnection(env)
	lmigration.Migrate(db, categoryTableMigration, sqlite.GetMigration)
	cr := New(db)

	ts.env = env
	ts.db = db
	ts.cr = cr
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
	ts.True( exists, "Exists deveria retornar verdadeiro para um uid válido")
}

func (ts *TestSuite) TestNotExists() {
	uid, err := luuid.New()

	if err != nil {
		ts.Fail(err.Error())
	}

	exists, err := ts.cr.Exists(uid)

	ts.NoError(err, "Exists se informado um uid inválido não deveria retornar erro")
	ts.False( exists, "Exists deveria retornar falso para um uid válido")
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

func (ts *TestSuite) TestHasSubcategories() {
	ts.makeTree()

	hasSubcategories, err := ts.cr.HasSubcategories(
		categoriesTree["filme"].GetUid())

	ts.NoError(err, "HasSubcategories se informado um uid válido não deveria retornar erro")
	ts.True( hasSubcategories, "HasSubcategories deveria retornar verdadeiro para um uid de uma categoria com subcategorias")
}

func (ts *TestSuite) TestHasNoSubcategories() {
	ts.cr.Create(mockCategory)

	hasSubcategories, err := ts.cr.HasSubcategories(mockUidCategory)

	ts.NoError(err, "HasSubcategories se informado um uid válido não deveria retornar erro")
	ts.False( hasSubcategories, "HasSubcategories deveria retornar falso para um uid de uma categoria sem subcategorias")
}

func (ts *TestSuite) TestIsSubcategory() {
	ts.makeTree()

	isSubcategory, err := ts.cr.IsSubcategory(
		categoriesTree["filme"].GetUid(),
		categoriesTree["terror"].GetUid())

	ts.NoError(err, "IsSubcategory, se informado duas chaves válidas não deveria retornar erro")
	ts.True( isSubcategory, "IsSubcategory deveria retornar verdadeiro para duas categorias que são uma subcategoria direta da outra")
}

func (ts *TestSuite) TestNotIsSubcategory() {
	ts.makeTree()

	isSubcategory, err := ts.cr.IsSubcategory(
		categoriesTree["filme"].GetUid(),
		categoriesTree["serial_killer"].GetUid())

	ts.NoError(err, "IsSubcategory, se informado duas chaves válidas não deveria retornar erro")
	ts.False( isSubcategory, "IsSubcategory deveria retornar falso para duas categorias que não são uma subcategoria direta da outra")
}

func (ts *TestSuite) TestIsAncestor() {
	ts.makeTree()

	isAncestor, err := ts.cr.IsAncestor(
		categoriesTree["filme"].GetUid(),
		categoriesTree["terror"].GetUid())

	ts.NoError(err, "IsAncestor, se informado duas chaves válidas não deveria retornar erro")
	ts.True( isAncestor, "IsAncestor deveria retornar verdadeiro para duas categorias que são relacionadas diretamente")

	isAncestor, err = ts.cr.IsAncestor(
		categoriesTree["filme"].GetUid(),
		categoriesTree["serial_killer"].GetUid())

	ts.NoError(err, "IsAncestor, se informado duas chaves válidas não deveria retornar erro")
	ts.True( isAncestor, "IsAncestor deveria retornar verdadeiro para duas categorias que são relacionadas")
}

func (ts *TestSuite) TestNotIsAncestor() {
	ts.makeTree()

	isAncestor, err := ts.cr.IsAncestor(
		categoriesTree["livro"].GetUid(),
		categoriesTree["terror"].GetUid())

	ts.NoError(err, "IsAncestor, se informado duas chaves válidas não deveria retornar erro")
	ts.False( isAncestor, "IsAncestor deveria retornar falso para duas categorias que não são relacionadas")

	isAncestor, err = ts.cr.IsAncestor(
		categoriesTree["serial_killer"].GetUid(),
		categoriesTree["filme"].GetUid())

	ts.NoError(err, "IsAncestor, se informado duas chaves válidas não deveria retornar erro")
	ts.False( isAncestor, "IsAncestor deveria retornar falso para duas categorias que são relacionadas, mas não na ordem ancestral->descendente")
}

func (ts *TestSuite) TestAreRelated() {
	ts.makeTree()

	areRelated, err := ts.cr.AreRelated(
		categoriesTree["filme"].GetUid(),
		categoriesTree["serial_killer"].GetUid())

	ts.NoError(err, "AreRelated, se informado duas chaves válidas não deveria retornar erro")
	ts.True( areRelated, "AreRelated deveria retornar verdadeiro para duas categorias que são relacionadas por parentesco")

	areRelated, err = ts.cr.AreRelated(
		categoriesTree["serial_killer"].GetUid(),
		categoriesTree["filme"].GetUid())

	ts.NoError(err, "AreRelated, se informado duas chaves válidas não deveria retornar erro")
	ts.True( areRelated, "AreRelated deveria retornar verdadeiro para duas categorias que são relacionadas por parentesco")
}

func (ts *TestSuite) TestNotAreRelated() {
	ts.makeTree()

	areRelated, err := ts.cr.AreRelated(
		categoriesTree["livro"].GetUid(),
		categoriesTree["jumpscare"].GetUid())

	ts.NoError(err, "AreRelated, se informado duas chaves válidas não deveria retornar erro")
	ts.False( areRelated, "AreRelated deveria retornar falso para duas categorias que não são uma relacionadas por parentesco")
}

func (ts *TestSuite) TestInsertSubcategory() {
	ts.insertAllCategories()

	err := ts.cr.InsertSubcategory(
		categoriesTree["filme"].GetUid(),
		categoriesTree["terror"].GetUid(),
		time.Now())

	ts.NoError(err, "InsertSubcategory, se informado duas chaves válidas não deveria retornar erro")

	isSubcategory, err := ts.cr.IsSubcategory(
		categoriesTree["filme"].GetUid(),
		categoriesTree["terror"].GetUid())

	ts.True( isSubcategory, "IsSubcategory deveria retornar true para uma relação criada usando InsertSubcategory")
}

func (ts *TestSuite) TestGetAllRootCategories_Empty() {}

func (ts *TestSuite) TestGetSubcategories_Empty() {
	ts.makeTree()

	subcategories, err := ts.cr.GetSubcategories(categoriesTree["jumpscare"].GetUid())

	ts.NoError(err, "Tentar recuperar subcategorias de uma categoria válida não deveria retornar erro")
	ts.Empty(subcategories, "Tentar recuperar subcategorias de uma categoria vazia deveria retornar uma lista de subcategorias vazia")
}

func (ts *TestSuite) TestGetSubcategories_NotEmpty() {
	ts.makeTree()

	subcategories, err := ts.cr.GetSubcategories(categoriesTree["terror"].GetUid())

	ts.NoError(err, "Tentar recuperar subcategorias de uma categoria válida não deveria retornar erro")
	ts.Len(subcategories, 2)

	var mockCategory = categoriesTree["jumpscare"]
	var category domain.Category

	if subcategories[0].GetUid() == categoriesTree["jumpscare"].GetUid() {
		category = subcategories[0]
	} else if len(subcategories) > 1 && subcategories[1].GetUid() == categoriesTree["jumpscare"].GetUid() {
		category = subcategories[1]
	} else {
		ts.Fail("Tentar recuperar subcategorias de uma categoria válida deveria retornar todos os dados de suas subcategorias")
	}

	ts.Equal(mockCategory.GetUid(), category.GetUid())
	ts.Equal(mockCategory.Name, category.Name)
	ts.Equal(mockCategory.Description.Content, category.Description.Content)
	ts.Equal(mockCategory.Description.UseMarkdown, category.Description.UseMarkdown)
	ts.NotZero(category.CreatedAt)
	ts.NotZero(category.UpdatedAt)
}

func TestSqliteCategory(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
