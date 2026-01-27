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

type ToggleMainCategoryTestSuite struct {
	suite.Suite
	post ltests.RequestFuncType
	patch ltests.RequestFuncType
	linkUidString string
	categoryUidString string
}

func (ts *ToggleMainCategoryTestSuite) SetupSuite() {
	env, _ := lenv.LoadTest()
	db, _ := sqlite.GetConnection(env)
	urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api"
	c := resty.New()

	ltests.CleanTable(db, "belongs_to")
	db.Close()

	ts.patch = ltests.GetJSONPatch(c, urlBase + "/categories/{categoryUid}/links/{linkUid}")

	postLink := ltests.GetJSONPost(c, urlBase + "/links")
	postCategory := ltests.GetJSONPost(c, urlBase + "/categories")
	postBelongsTo := ltests.GetJSONPost(c, urlBase + "/categories/{categoryUid}/links")

	resLink, _ := postLink(nil, domain.MockLinkMapRequest)
	ts.linkUidString = string(resLink.Body())

	resCategory, _ := postCategory(nil, domain.MockCategoryMapRequest)
	ts.categoryUidString = string(resCategory.Body())

	postBelongsTo(
		map[string]string{
			"categoryUid": ts.categoryUidString},
		map[string]string{
			"linkUid": ts.linkUidString,
			"isMain": "true",
		})
}

func (ts *ToggleMainCategoryTestSuite) TestToggleMainCategory() {
	res, _ := ts.patch(
		map[string]string{
			"categoryUid": ts.categoryUidString,
			"linkUid": ts.linkUidString},
		map[string]string{
			"isMain": "false",
		})

	ts.Empty(
		string(res.Body()),
		"Tentar alternar a categoria principal de um link não deveria retornar nada")
}

func (ts *ToggleMainCategoryTestSuite) TestToggleMainCategory_Error() {
	res, _ := ts.patch(
		map[string]string{
			"categoryUid": domain.MockLink.Name,
			"linkUid": ts.linkUidString},
		map[string]string{
			"isMain": domain.MockLink.Name,
		})

	ts.ElementsMatch(ltests.GetErrorKeys(res.Body()), []string{"categoryUid"}) 
}

func (ts *ToggleMainCategoryTestSuite) TestToggleMainCategory_NotExists() {
	res, _ := ts.patch(
		map[string]string{
			"categoryUid": domain.MockUidCategory.String(),
			"linkUid": domain.MockUidLink.String()},
		map[string]string{
			"isMain": "false",
		})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.NOT_BELONGS.Error()),
		res.Body())

	ts.Equal(
		resBody,
		expectedJson,
		"Tentar alternar a categoria principal de um link, ambos não relacionados, deveria retornar erro contendo " + domain.NOT_BELONGS.Error())
}

func TestToggleMainCategoryAllTests(t *testing.T) {
	suite.Run(t, new(ToggleMainCategoryTestSuite))
}
