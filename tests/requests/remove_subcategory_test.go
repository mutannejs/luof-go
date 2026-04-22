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

type RemoveSubcategoryTestSuite struct {
	suite.Suite
	post ltests.RequestFuncType
	delete ltests.DeleteFuncType
	postCategory ltests.RequestFuncType
	deleteCategory ltests.DeleteFuncType
	fatherUidString string
	childUidString string
}

func (ts *RemoveSubcategoryTestSuite) SetupSuite() {
	env, _ := lenv.LoadTest()
	db, _ := sqlite.GetConnection(env)
	urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api/categories"

	c := resty.New()
	ts.post = ltests.GetJSONPost(c, urlBase + "/{categoryUid}/subcategories")
	ts.delete = ltests.GetDelete(c, urlBase + "/{categoryUid}/subcategories/{childUid}")
	ts.postCategory = ltests.GetJSONPost(c, urlBase)
	ts.deleteCategory = ltests.GetDelete(c, urlBase + "/{categoryUid}")

	ltests.CleanTable(db, "category")
	db.Close()
}

func (ts *RemoveSubcategoryTestSuite) SetupTest() {
	resCategory, _ := ts.postCategory(nil, domain.MockCategoryMapRequest)
	ts.fatherUidString = string(resCategory.Body())

	resCategory, _ = ts.postCategory(nil, domain.AlternativeMockCategoryMapRequest)
	ts.childUidString = string(resCategory.Body())

	ts.post(
		map[string]string{
			"categoryUid": ts.fatherUidString},
		map[string]string{
			"childUid": ts.childUidString,
		})
}

func (ts *RemoveSubcategoryTestSuite) TearDownTest() {
	ts.deleteCategory(
		map[string]string{
			"categoryUid": ts.childUidString,
		})
	ts.deleteCategory(
		map[string]string{
			"categoryUid": ts.fatherUidString,
		})
}

func (ts *RemoveSubcategoryTestSuite) TestRemoveSubcategory() {
	res, _ := ts.delete(
		map[string]string{
			"categoryUid": ts.fatherUidString,
			"childUid": ts.childUidString,
		})

	ts.Empty(
		string(res.Body()),
		"Tentar remover uma subcategoria, ambas relacionadas não deveria retornar nada")
}

func (ts *RemoveSubcategoryTestSuite) TestRemoveSubcategory_Error() {
	res, _ := ts.delete(
		map[string]string{
			"categoryUid": domain.MockCategory.Name,
			"childUid": domain.MockCategory.Name,
		})

	ts.ElementsMatch(ltests.GetErrorKeys(res.Body()), []string{"categoryUid", "childUid"}) 
}

func (ts *RemoveSubcategoryTestSuite) TestRemoveSubcategory_NotExists() {
	res, _ := ts.delete(
		map[string]string{
			"categoryUid": ts.fatherUidString,
			"childUid": domain.MockUidCategory.String(),
		})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.NOT_IS_SUBCATEGORY.Error()),
		res.Body())

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar remover uma subcategoria tal que a relação não existe, deveria retornar erro contendo " + domain.NOT_IS_SUBCATEGORY.Error()) 
}

func TestRemoveSubcategoryAllTests(t *testing.T) {
	suite.Run(t, new(RemoveSubcategoryTestSuite))
}
