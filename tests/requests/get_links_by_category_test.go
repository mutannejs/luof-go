package requests

import (
	"encoding/json"
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/ltests"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
)

type GetLinksByCategoryTestSuite struct {
	suite.Suite
	get ltests.RequestFuncType
	uidLinkString string
	alternativeUidLinkString string
	categoryUidString string
	emptyCategoryUidString string
}

func (ts *GetLinksByCategoryTestSuite) SetupSuite() {
	env, _ := lenv.LoadTest()
	urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api"
	c := resty.New()

	ts.get = ltests.GetGet(c, urlBase + "/categories/{categoryUid}/links")

	postLink := ltests.GetJSONPost(c, urlBase + "/links")
	postCategory := ltests.GetJSONPost(c, urlBase + "/categories")
	postBelongsTo := ltests.GetJSONPost(c, urlBase + "/categories/{categoryUid}/links")

	resCategory, _ := postCategory(nil, domain.MockCategoryMapRequest)
	ts.categoryUidString = string(resCategory.Body())

	resCategory, _ = postCategory(nil, domain.AlternativeMockCategoryMapRequest)
	ts.emptyCategoryUidString = string(resCategory.Body())

	resLink, _ := postLink(nil, domain.MockLinkMapRequest)
	ts.uidLinkString = string(resLink.Body())

	postBelongsTo(
		map[string]string{
			"categoryUid": ts.categoryUidString},
		map[string]string{
			"linkUid": ts.uidLinkString,
			"isMain": "true",
		})

	resLink, _ = postLink(nil, domain.AlternativeMockLinkMapRequest)
	ts.alternativeUidLinkString = string(resLink.Body())

	postBelongsTo(
		map[string]string{
			"categoryUid": ts.categoryUidString},
		map[string]string{
			"linkUid": ts.alternativeUidLinkString,
			"isMain": "false",
		})
}

func (ts *GetLinksByCategoryTestSuite) TestGetLinksByCategory() {
	res, _ := ts.get(
		map[string]string{
			"categoryUid": ts.categoryUidString,
		},
		nil)

	var linksJson []map[string]string
	json.Unmarshal(res.Body(), &linksJson)

	ts.Equal(
		res.StatusCode(),
		200,
		"Tentar recuperar todos os links de uma categoria válida deveria retornar status 200")

	ts.Len(
		linksJson,
		2,
		"Tentar recuperar todos os links de uma categoria válida deveria retornar uma lista com todos seus links")
}

func (ts *GetLinksByCategoryTestSuite) TestGetLinksByCategory_Error() {
	res, _ := ts.get(
		map[string]string{
			"categoryUid": domain.MockLink.Name,
		},
		nil)

	ts.Equal(
		res.StatusCode(),
		400,
		"Tentar recuperar todos os links de uma categoria passando parâmetros errados deveria retornar status 400")

	ts.ElementsMatch([]string{"categoryUid"}, ltests.GetErrorKeys(res.Body())) 
}

func (ts *GetLinksByCategoryTestSuite) TestGetLinksByCategory_NotExists() {
	res, _ := ts.get(
		map[string]string{
			"categoryUid": domain.AlternativeMockUidCategory.String(),
		},
		nil)

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.CATEGORY_NOT_EXISTS),
		res.Body())

	ts.Equal(
		res.StatusCode(),
		404,
		"Tentar recuperar todos os links de uma categoria que não existe deveria retornar status 404")

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar recuperar os link de uma categoria inválida deveria retornar erro contendo " + domain.CATEGORY_NOT_EXISTS) 
}

func (ts *GetLinksByCategoryTestSuite) TestGetLinksByCategory_Empty() {
	res, _ := ts.get(
		map[string]string{
			"categoryUid": ts.emptyCategoryUidString,
		},
		nil)

	expectedJson, resBody := ltests.TrimResponse(
		[]byte("[]"),
		res.Body())

	ts.Equal(
		res.StatusCode(),
		200,
		"Tentar recuperar todos os links de uma categoria vazia deveria retornar status 200")

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar recuperar os link de uma categoria vazia deveria retornar um array vazio") 
}

func TestGetLinksByCategoryAllTests(t *testing.T) {
	suite.Run(t, new(GetLinksByCategoryTestSuite))
}
