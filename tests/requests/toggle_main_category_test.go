package requests

import (
	"testing"

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
	alternativeCategoryUidString string
}

func (ts *ToggleMainCategoryTestSuite) SetupSuite() {
	env, _ := lenv.LoadTest()
	urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api"
	c := resty.New()

	ts.patch = ltests.GetJSONPatch(c, urlBase + "/categories/{categoryUid}/links/{linkUid}")

	postLink := ltests.GetJSONPost(c, urlBase + "/links")
	postCategory := ltests.GetJSONPost(c, urlBase + "/categories")
	postBelongsTo := ltests.GetJSONPost(c, urlBase + "/categories/{categoryUid}/links")

	resLink, _ := postLink(nil, domain.MockLinkMapRequest)
	ts.linkUidString = string(resLink.Body())

	resCategory, _ := postCategory(nil, domain.MockCategoryMapRequest)
	ts.categoryUidString = string(resCategory.Body())

	resCategory, _ = postCategory(nil, domain.AlternativeMockCategoryMapRequest)
	ts.alternativeCategoryUidString = string(resCategory.Body())

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

	ts.Equal(
		res.StatusCode(),
		204,
		"Tentar alternar a categoria principal de um link deveria retornar status 204")

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

	ts.Equal(
		res.StatusCode(),
		400,
		"Tentar alternar a categoria principal de um link passando parâmetros inválidos deveria retornar status 400")

	ts.ElementsMatch([]string{"categoryUid"}, ltests.GetErrorKeys(res.Body()))
}

func (ts *ToggleMainCategoryTestSuite) TestToggleMainCategory_LinkNotExists() {
	res, _ := ts.patch(
		map[string]string{
			"categoryUid": ts.categoryUidString,
			"linkUid": domain.MockUidLink.String()},
		map[string]string{
			"isMain": "false",
		})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.LINK_NOT_EXISTS.Error()),
		res.Body())

	ts.Equal(
		res.StatusCode(),
		404,
		"Tentar alternar a categoria principal de um link que não existe deveria retornar status 404")

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar alternar a categoria principal de um link, ambos não relacionados, deveria retornar erro contendo " + domain.LINK_NOT_EXISTS.Error())
}

func (ts *ToggleMainCategoryTestSuite) TestToggleMainCategory_CategoryNotExists() {
	res, _ := ts.patch(
		map[string]string{
			"categoryUid": domain.MockUidCategory.String(),
			"linkUid": ts.linkUidString},
		map[string]string{
			"isMain": "false",
		})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.CATEGORY_NOT_EXISTS.Error()),
		res.Body())

	ts.Equal(
		res.StatusCode(),
		404,
		"Tentar alternar a categoria principal de um link, sendo que a categoria não existe, deveria retornar status 404")

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar alternar a categoria principal de um link, sendo que a categoria não existe, deveria retornar erro contendo " + domain.CATEGORY_NOT_EXISTS.Error())
}

func (ts *ToggleMainCategoryTestSuite) TestToggleMainCategory_NotExists() {
	res, _ := ts.patch(
		map[string]string{
			"categoryUid": ts.alternativeCategoryUidString,
			"linkUid": ts.linkUidString},
		map[string]string{
			"isMain": "false",
		})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.NOT_BELONGS.Error()),
		res.Body())

	ts.Equal(
		res.StatusCode(),
		404,
		"Tentar alternar a categoria principal de um link, ambos não relacionados, deveria retornar status 404")

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar alternar a categoria principal de um link, ambos não relacionados, deveria retornar erro contendo " + domain.NOT_BELONGS.Error())
}

func TestToggleMainCategoryAllTests(t *testing.T) {
	suite.Run(t, new(ToggleMainCategoryTestSuite))
}
