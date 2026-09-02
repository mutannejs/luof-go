package requests

import (
	"testing"

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
	urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api/categories"

	c := resty.New()
	ts.post = ltests.GetJSONPost(c, urlBase + "/{categoryUid}/subcategories")
	ts.postCategory = ltests.GetJSONPost(c, urlBase)
	ts.deleteCategory = ltests.GetDelete(c, urlBase + "/{categoryUid}")

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

	ts.Equal(
		res.StatusCode(),
		204,
		"Tentar inserir uma subcategoria, ambas ainda não relacionadas, deveria retornar status 204")

	ts.Empty(
		string(res.Body()),
		"Tentar inserir uma subcategoria, ambas ainda não relacionadas, não deveria retornar nada")
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

	ts.Equal(
		res.StatusCode(),
		204,
		"Tentar inserir uma subcategoria, ambas não diretamente relacionadas, deveria retornar status 204")

	ts.Empty(
		string(res.Body()),
		"Tentar inserir uma subcategoria, ambas não diretamente relacionadas, não deveria retornar nada")
}

func (ts *InsertSubcategoryTestSuite) TestInsertSubcategory_Error() {
	res, _ := ts.post(
		map[string]string{
			"categoryUid": domain.MockCategory.Name},
		map[string]string{
			"childUid": domain.MockCategory.Name,
		})

	ts.Equal(
		res.StatusCode(),
		400,
		"Tentar inserir uma subcategoria passando parâmetros inválidos deveria retornar status 400")

	ts.ElementsMatch([]string{"categoryUid"}, ltests.GetErrorKeys(res.Body()))
}

func (ts *InsertSubcategoryTestSuite) TestInsertSubcategory_FatherNotEXists() {
	res, _ := ts.post(
		map[string]string{
			"categoryUid": domain.MockUidCategory.String()},
		map[string]string{
			"childUid": ts.childUidString,
		})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.FATHER_NOT_EXISTS),
		res.Body())

	ts.Equal(
		res.StatusCode(),
		404,
		"Tentar inserir uma subcategoria em uma categoria pai que não existe deveria retornar status 404")

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar inserir uma subcategoria em uma categoria pai que não existe deveria retornar erro contendo " + domain.FATHER_NOT_EXISTS)
}

func (ts *InsertSubcategoryTestSuite) TestInsertSubcategory_ChildNotEXists() {
	res, _ := ts.post(
		map[string]string{
			"categoryUid": ts.fatherUidString},
		map[string]string{
			"childUid": domain.MockUidCategory.String(),
		})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.CHILD_NOT_EXISTS),
		res.Body())

	ts.Equal(
		res.StatusCode(),
		404,
		"Tentar inserir uma categoria que não existe em outra deveria retornar status 404")

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar inserir uma categoria que não existe em outra deveria retornar erro contendo " + domain.CHILD_NOT_EXISTS)
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
		ltests.GetResponseMessage(domain.IS_SUBCATEGORY),
		res.Body())

	ts.Equal(
		res.StatusCode(),
		409,
		"Tentar inserir uma subcategoria, ambas já relacionadas, deveria retornar status 409")

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar inserir uma subcategoria, ambas já relacionadas, deveria retornar erro contendo " + domain.IS_SUBCATEGORY)
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
		ltests.GetResponseMessage(domain.ANCESTOR_NOT_BECOME_A_SUBCATEGORY),
		res.Body())

	ts.Equal(
		res.StatusCode(),
		409,
		"Tentar inserir uma subcategoria, sendo que a relação inversa já existe, deveria retornar status 409")

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar inserir uma subcategoria, sendo que a relação inversa já existe, deveria retornar erro contendo " + domain.ANCESTOR_NOT_BECOME_A_SUBCATEGORY)
}

func TestInsertSubcategoryAllTests(t *testing.T) {
	suite.Run(t, new(InsertSubcategoryTestSuite))
}
