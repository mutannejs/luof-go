package requests

import (
	"testing"

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
	alternativeCategoryUidString string
}

func (ts *RemoveSubcategoryTestSuite) SetupSuite() {
	env, _ := lenv.LoadTest()
	urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api/categories"

	c := resty.New()
	ts.post = ltests.GetJSONPost(c, urlBase + "/{categoryUid}/subcategories")
	ts.delete = ltests.GetDelete(c, urlBase + "/{categoryUid}/subcategories/{childUid}")
	ts.postCategory = ltests.GetJSONPost(c, urlBase)
	ts.deleteCategory = ltests.GetDelete(c, urlBase + "/{categoryUid}")

	resCategory, _ := ts.postCategory(nil, domain.MockCategoryMapRequest)
	ts.alternativeCategoryUidString = string(resCategory.Body())
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

	ts.Equal(
		res.StatusCode(),
		204,
		"Tentar remover uma subcategoria, ambos relacionadas, deveria retornar status 204")

	ts.Empty(
		string(res.Body()),
		"Tentar remover uma subcategoria, ambas relacionadas, não deveria retornar nada")
}

func (ts *RemoveSubcategoryTestSuite) TestRemoveSubcategory_Error() {
	res, _ := ts.delete(
		map[string]string{
			"categoryUid": domain.MockCategory.Name,
			"childUid": domain.MockCategory.Name,
		})

	ts.Equal(
		res.StatusCode(),
		400,
		"Tentar remover uma subcategoria passando parâmetros inválidos deveria retornar status 400")

	ts.ElementsMatch([]string{"categoryUid", "childUid"}, ltests.GetErrorKeys(res.Body()))
}

func (ts *RemoveSubcategoryTestSuite) TestRemoveSubcategory_FatherNotExists() {
	res, _ := ts.delete(
		map[string]string{
			"categoryUid": domain.MockUidCategory.String(),
			"childUid": ts.childUidString,
		})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.FATHER_NOT_EXISTS),
		res.Body())

	ts.Equal(
		res.StatusCode(),
		404,
		"Tentar remover uma categoria de outra que não existe deveria retornar status 404")

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar remover uma categoria de outra que não existe deveria retornar erro contendo " + domain.FATHER_NOT_EXISTS)
}

func (ts *RemoveSubcategoryTestSuite) TestRemoveSubcategory_ChildNotExists() {
	res, _ := ts.delete(
		map[string]string{
			"categoryUid": ts.fatherUidString,
			"childUid": domain.MockUidCategory.String(),
		})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.CHILD_NOT_EXISTS),
		res.Body())

	ts.Equal(
		res.StatusCode(),
		404,
		"Tentar remover uma categoria que não existe de outra deveria retornar status 404")

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar remover uma subcategoria tal que a relação não existe, deveria retornar erro contendo " + domain.CHILD_NOT_EXISTS)
}

func (ts *RemoveSubcategoryTestSuite) TestRemoveSubcategory_NotExists() {
	res, _ := ts.delete(
		map[string]string{
			"categoryUid": ts.fatherUidString,
			"childUid": ts.alternativeCategoryUidString,
		})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.NOT_IS_SUBCATEGORY),
		res.Body())

	ts.Equal(
		res.StatusCode(),
		404,
		"Tentar remover uma subcategoria tal que a relação não existe deveria retornar status 404")

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar remover uma subcategoria tal que a relação não existe, deveria retornar erro contendo " + domain.NOT_IS_SUBCATEGORY)
}

func TestRemoveSubcategoryAllTests(t *testing.T) {
	suite.Run(t, new(RemoveSubcategoryTestSuite))
}
