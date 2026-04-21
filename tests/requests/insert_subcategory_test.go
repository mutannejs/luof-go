package requests

import (
	"testing"

	"github.com/mutannejs/luof-go/adapters/sqlite"
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/ltests"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
)

type InsertSubcategoryTestSuite struct {
	suite.Suite
	post ltests.RequestFuncType
	postCategory ltests.RequestFuncType
	deleteCategory ltests.DeleteFuncType
	fatherUidString string
	childUidString string
	netoUidString string
}

func (ts *InsertSubcategoryTestSuite) SetupSuite() {
	env, _ := lenv.LoadTest()
	db, _ := sqlite.GetConnection(env)
	urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api/categories"

	c := resty.New()
	ts.post = ltests.GetJSONPost(c, urlBase + "/{categoryUid}/subcategories")
	ts.postCategory = ltests.GetJSONPost(c, urlBase)
	ts.deleteCategory = ltests.GetDelete(c, urlBase + "/{categoryUid}")

	ltests.CleanTable(db, "category")
	db.Close()
}

func (ts *InsertSubcategoryTestSuite) SetupTest() {
	resCategory, _ := ts.postCategory(nil, domain.MockCategoryMapRequest)
	ts.fatherUidString = string(resCategory.Body())

	resCategory, _ = ts.postCategory(nil, domain.AlternativeMockCategoryMapRequest)
	ts.childUidString = string(resCategory.Body())

	resCategory, _ = ts.postCategory(nil, domain.AlternativeMockCategoryMapRequest)
	ts.netoUidString = string(resCategory.Body())
}

func (ts *InsertSubcategoryTestSuite) TearDownTest() {
	ts.deleteCategory(
		map[string]string{
			"categoryUid": ts.netoUidString,
		})
	ts.deleteCategory(
		map[string]string{
			"categoryUid": ts.childUidString,
		})
	ts.deleteCategory(
		map[string]string{
			"categoryUid": ts.fatherUidString,
		})
}

func (ts *InsertSubcategoryTestSuite) TestInsertSubcategory() {
	res, _ := ts.post(
		map[string]string{
			"categoryUid": ts.fatherUidString},
		map[string]string{
			"childUid": ts.childUidString,
		})

	ts.Empty(
		string(res.Body()),
		"Tentar inserir uma subcategoria, ambas ainda não relacionadas não deveria retornar nada")
}

func (ts *InsertSubcategoryTestSuite) TestInsertSubcategory_Relateds() {
	res, _ := ts.post(
		map[string]string{
			"categoryUid": ts.fatherUidString},
		map[string]string{
			"childUid": ts.childUidString,
		})
	res, _ = ts.post(
		map[string]string{
			"categoryUid": ts.childUidString},
		map[string]string{
			"childUid": ts.netoUidString,
		})

	res, _ = ts.post(
		map[string]string{
			"categoryUid": ts.fatherUidString},
		map[string]string{
			"childUid": ts.netoUidString,
		})

	ts.Empty(
		string(res.Body()),
		"Tentar inserir uma subcategoria, ambas ainda não relacionadas não deveria retornar nada")
}

func (ts *InsertSubcategoryTestSuite) TestInsertSubcategory_Error() {
	res, _ := ts.post(
		map[string]string{
			"categoryUid": domain.MockCategory.Name},
		map[string]string{
			"childUid": domain.MockCategory.Name,
		})

	ts.ElementsMatch(ltests.GetErrorKeys(res.Body()), []string{"categoryUid"}) 
}

func (ts *InsertSubcategoryTestSuite) TestInsertSubcategory_AlreadyExists() {
	res, _ := ts.post(
		map[string]string{
			"categoryUid": ts.fatherUidString},
		map[string]string{
			"childUid": ts.childUidString,
		})

	res, _ = ts.post(
		map[string]string{
			"categoryUid": ts.fatherUidString},
		map[string]string{
			"childUid": ts.childUidString,
		})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.IS_SUBCATEGORY.Error()),
		res.Body())

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar inserir uma subcategoria, ambas já relacionadas, deveria retornar erro contendo " + domain.IS_SUBCATEGORY.Error()) 
}

func (ts *InsertSubcategoryTestSuite) TestInsertSubcategory_AncestorNotBecomeASubcategory() {
	res, _ := ts.post(
		map[string]string{
			"categoryUid": ts.fatherUidString},
		map[string]string{
			"childUid": ts.childUidString,
		})
	res, _ = ts.post(
		map[string]string{
			"categoryUid": ts.childUidString},
		map[string]string{
			"childUid": ts.netoUidString,
		})

	res, _ = ts.post(
		map[string]string{
			"categoryUid": ts.netoUidString},
		map[string]string{
			"childUid": ts.fatherUidString,
		})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.ANCESTOR_NOT_BECOME_A_SUBCATEGORY.Error()),
		res.Body())

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar inserir uma subcategoria, ambas já relacionadas, deveria retornar erro contendo " + domain.ANCESTOR_NOT_BECOME_A_SUBCATEGORY.Error()) 
}

func TestInsertSubcategoryAllTests(t *testing.T) {
	suite.Run(t, new(InsertSubcategoryTestSuite))
}
