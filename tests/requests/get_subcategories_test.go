package requests

import (
	"encoding/json"
	"testing"

	"github.com/mutannejs/luof-go/adapters/sqlite"
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/ltests"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
)

type GetSubcategoriesTestSuite struct {
	suite.Suite
	get ltests.RequestFuncType
	fatherUidString string
	childUidString string
	alternativeChildUidString string
}

func (ts *GetSubcategoriesTestSuite) SetupSuite() {
	env, _ := lenv.LoadTest()
	db, _ := sqlite.GetConnection(env)
	urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api/categories"

	c := resty.New()
	ts.get = ltests.GetGet(c, urlBase + "/{categoryUid}/subcategories")

	postCategory := ltests.GetJSONPost(c, urlBase)
	postSubcategory := ltests.GetJSONPost(c, urlBase + "/{categoryUid}/subcategories")

	ltests.CleanTable(db, "category")
	db.Close()

	resCategory, _ := postCategory(nil, domain.MockCategoryMapRequest)
	ts.fatherUidString = string(resCategory.Body())

	resCategory, _ = postCategory(nil, domain.AlternativeMockCategoryMapRequest)
	ts.childUidString = string(resCategory.Body())

	postSubcategory(
		map[string]string{
			"categoryUid": ts.fatherUidString},
		map[string]string{
			"childUid": ts.childUidString,
		})

	resCategory, _ = postCategory(nil, domain.AlternativeMockCategoryMapRequest)
	ts.alternativeChildUidString = string(resCategory.Body())

	postSubcategory(
		map[string]string{
			"categoryUid": ts.fatherUidString},
		map[string]string{
			"childUid": ts.alternativeChildUidString,
		})
}

func (ts *GetSubcategoriesTestSuite) TestGetSubcategories() {
	res, _ := ts.get(
		map[string]string{
			"categoryUid": ts.fatherUidString,
		},
		nil)

	var categoriesJson []map[string]string
	json.Unmarshal(res.Body(), &categoriesJson)

	ts.Len(
		categoriesJson,
		2,
		"Tentar recuperar todos as subcategorias de uma categoria válida deveria retornar uma lista com todas suas subcategorias")
}

func (ts *GetSubcategoriesTestSuite) TestGetSubcategories_Error() {
	res, _ := ts.get(
		map[string]string{
			"categoryUid": domain.MockCategory.Name,
		},
		nil)

	ts.ElementsMatch(ltests.GetErrorKeys(res.Body()), []string{"categoryUid"}) 
}

func (ts *GetSubcategoriesTestSuite) TestGetSubcategories_NotExists() {
	res, _ := ts.get(
		map[string]string{
			"categoryUid": domain.AlternativeMockUidCategory.String(),
		},
		nil)

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.CATEGORY_NOT_EXISTS.Error()),
		res.Body())

	ts.Equal(
		resBody,
		expectedJson,
		"Tentar recuperar as subcategorias de uma categoria inválida deveria retornar erro contendo " + domain.CATEGORY_NOT_EXISTS.Error()) 
}

func TestGetSubcategoriesAllTests(t *testing.T) {
	suite.Run(t, new(GetSubcategoriesTestSuite))
}
