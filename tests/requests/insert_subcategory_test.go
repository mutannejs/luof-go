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
	fatherUidString string
	childUidString string
}

func (ts *InsertSubcategoryTestSuite) SetupSuite() {
	env, _ := lenv.LoadTest()
	db, _ := sqlite.GetConnection(env)
	urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api/categories"

	c := resty.New()
	ts.post = ltests.GetJSONPost(c, urlBase + "/{categoryUid}/subcategories")

	postCategory := ltests.GetJSONPost(c, urlBase)

	ltests.CleanTable(db, "category")
	db.Close()

	resCategory, _ := postCategory(nil, domain.MockCategoryMapRequest)
	ts.fatherUidString = string(resCategory.Body())

	resCategory, _ = postCategory(nil, domain.AlternativeMockCategoryMapRequest)
	ts.childUidString = string(resCategory.Body())
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
