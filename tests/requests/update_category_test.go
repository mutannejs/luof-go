package requests

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/ltests"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
)

type UpdateCategoryTestSuite struct {
	suite.Suite
	categoryUid string
	put ltests.RequestFuncType
}

func (ts *UpdateCategoryTestSuite) SetupSuite() {
	env, _ := lenv.LoadTest()
	urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api/categories"

	c := resty.New()

	post := ltests.GetJSONPost(c, urlBase)
	res, _ := post(nil, domain.MockCategoryMapRequest)

	ts.categoryUid = res.String()
	ts.put = ltests.GetJSONPut(c, urlBase + "/{categoryUid}")
}

func (ts *UpdateCategoryTestSuite) TestUpdateCategory() {
	res, err := ts.put(
		map[string]string{
			"categoryUid": ts.categoryUid},
		domain.MockCategoryMapRequest)

	ts.NoError(err)

	ts.Equal(
		res.StatusCode(),
		204,
		"Tentar atualizar uma categoria deveria retornar status 204")

	ts.Empty(
		string(res.Body()),
		"Tentar atualizar uma categoria passando parâmetros válidos não deveria retornar nada")
}

func (ts *UpdateCategoryTestSuite) TestUpdateCategory_ParamRequired() {
	res, _ := ts.put(
		map[string]string{
			"categoryUid": ts.categoryUid},
		map[string]string{
			// "name": domain.AlternativeMockCategory.Name,
			"description": domain.AlternativeMockCategory.Description.Content,
			"useMarkdown": domain.AlternativeMockCategory.Description.Content,
		})

	ts.Equal(
		res.StatusCode(),
		400,
		"Tentar atualizar uma categoria passando parâmetros inválidosdeveria retornar status 400")

	ts.ElementsMatch([]string{"name", "useMarkdown"}, ltests.GetErrorKeys(res.Body()))
}

func (ts *UpdateCategoryTestSuite) TestUpdateCategory_NotExists() {
	res, _ := ts.put(
		map[string]string{
			"categoryUid": domain.MockUidCategory.String()},
		domain.MockCategoryMapRequest)

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.CATEGORY_NOT_EXISTS),
		res.Body())

	ts.Equal(
		res.StatusCode(),
		404,
		"Tentar atualizar uma categoria que não existe deveria retornar status 404")

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar atualizar uma categoria passando um uuid inválido deveria retornar o erro " + domain.CATEGORY_NOT_EXISTS)
}

func TestUpdateCategoryAllTests(t *testing.T) {
	suite.Run(t, new(UpdateCategoryTestSuite))
}
