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
	db, _ := sqlite.GetConnection(env)
	urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api"
	c := resty.New()

	ltests.CleanTable(db, "belongs_to")
	db.Close()

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

	ts.ElementsMatch(ltests.GetErrorKeys(res.Body()), []string{"categoryUid"}) 
}

func (ts *GetLinksByCategoryTestSuite) TestGetLinksByCategory_NotExists() {
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
		"Tentar recuperar os link de uma categoria inválida deveria retornar erro contendo " + domain.CATEGORY_NOT_EXISTS.Error()) 
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
		resBody,
		expectedJson,
		"Tentar recuperar os link de uma categoria vazia deveria retornar um array vazio") 
}

func TestGetLinksByCategoryAllTests(t *testing.T) {
	suite.Run(t, new(GetLinksByCategoryTestSuite))
}
