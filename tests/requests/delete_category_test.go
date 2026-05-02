package requests

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/ltests"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
)

type DeleteCategoryTestSuite struct {
	suite.Suite
	delete ltests.DeleteFuncType
	post ltests.RequestFuncType
	postLink ltests.RequestFuncType
	postBelongsTo ltests.RequestFuncType
	postSubcategory ltests.RequestFuncType
}

func (ts *DeleteCategoryTestSuite) SetupSuite() {
	env, _ := lenv.LoadTest()
	urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api"

	c := resty.New()
	ts.post = ltests.GetJSONPost(c, urlBase + "/categories")
	ts.delete = ltests.GetDelete(c, urlBase + "/categories/{categoryUid}")
	ts.postLink = ltests.GetJSONPost(c, urlBase + "/links")
	ts.postBelongsTo = ltests.GetJSONPost(c, urlBase + "/categories/{categoryUid}/links")
	ts.postSubcategory = ltests.GetJSONPost(c, urlBase + "/categories/{categoryUid}/subcategories")

	ts.post(nil, domain.MockCategoryMapRequest)
}

func (ts *DeleteCategoryTestSuite) TestDeleteCategory() {
	res, _ := ts.post(nil, domain.MockCategoryMapRequest)
	res, err := ts.delete(map[string]string{"categoryUid": res.String()})

	ts.NoError(err)

	ts.Empty(
		string(res.Body()),
		"Tentar deletar uma categoria passando um uuid válido não deveria retornar nada")
}

func (ts *DeleteCategoryTestSuite) TestDeleteCategory_ParamRequired() {
	res, _ := ts.delete(map[string]string{"categoryId": domain.MockUidCategory.String()})

	ts.ElementsMatch(ltests.GetErrorKeys(res.Body()), []string{"categoryUid"}) 
}

func (ts *DeleteCategoryTestSuite) TestDeleteCategory_NotExists() {
	res, _ := ts.delete(map[string]string{"categoryUid": domain.MockUidCategory.String()})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.CATEGORY_NOT_EXISTS.Error()),
		res.Body())

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar deletar uma categoria passando um uuid inválido deveria retornar o erro " + domain.LINK_NOT_EXISTS.Error())
}

func (ts *DeleteCategoryTestSuite) TestDeleteCategory_HasLinks() {
	res, _ := ts.post(nil, domain.MockCategoryMapRequest)
	linkUidString, _ := ts.postLink(nil, domain.MockLinkMapRequest)
	ts.postBelongsTo(
		map[string]string{
			"categoryUid": res.String()},
		map[string]string{
			"linkUid": string(linkUidString.Body()),
			"isMain": "false"},
		)
	res, _ = ts.delete(map[string]string{"categoryUid": res.String()})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.HAS_LINKS.Error()),
		res.Body())

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar deletar uma categoria com um ou mais links deveria retornar o erro " + domain.HAS_LINKS.Error())
}

func (ts *DeleteCategoryTestSuite) TestDeleteCategory_HasSubcategories() {
	res, _ := ts.post(nil, domain.MockCategoryMapRequest)
	childCategoryUidString, _ := ts.post(nil, domain.AlternativeMockCategoryMapRequest)
	ts.postSubcategory(
		map[string]string{
			"categoryUid": res.String()},
		map[string]string{
			"childUid": string(childCategoryUidString.Body())},
		)
	res, _ = ts.delete(map[string]string{"categoryUid": res.String()})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.HAS_SUBCATEGORIES.Error()),
		res.Body())

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar deletar uma categoria com uma ou mais subcategorias deveria retornar o erro " + domain.HAS_SUBCATEGORIES.Error())
}

func TestDeleteCategoryAllTests(t *testing.T) {
	suite.Run(t, new(DeleteCategoryTestSuite))
}
