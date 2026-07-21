package requests

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/ltests"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
)

type RemoveLinkFromCategoryTestSuite struct {
	suite.Suite
	post ltests.RequestFuncType
	delete ltests.DeleteFuncType
	linkUidString string
	alternativeLinkUidString string
	categoryUidString string
	alternativeCategoryUidString string
}

func (ts *RemoveLinkFromCategoryTestSuite) SetupSuite() {
	env, _ := lenv.LoadTest()
	urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api"

	c := resty.New()
	ts.post = ltests.GetJSONPost(c, urlBase + "/categories/{categoryUid}/links")
	ts.delete = ltests.GetDelete(c, urlBase + "/categories/{categoryUid}/links/{linkUid}")

	postLink := ltests.GetJSONPost(c, urlBase + "/links")
	postCategory := ltests.GetJSONPost(c, urlBase + "/categories")

	resLink, _ := postLink(nil, domain.MockLinkMapRequest)
	ts.linkUidString = string(resLink.Body())

	resLink, _ = postLink(nil, domain.MockLinkMapRequest)
	ts.alternativeLinkUidString = string(resLink.Body())

	resCategory, _ := postCategory(nil, domain.MockCategoryMapRequest)
	ts.categoryUidString = string(resCategory.Body())

	resCategory, _ = postCategory(nil, domain.MockCategoryMapRequest)
	ts.alternativeCategoryUidString = string(resCategory.Body())
}

func (ts *RemoveLinkFromCategoryTestSuite) TestRemoveLinkFromCategory() {
	ts.post(
		map[string]string{
			"categoryUid": ts.categoryUidString},
		map[string]string{
			"linkUid": ts.linkUidString,
			"isMain": "true",
		})

	res, _ := ts.delete(
		map[string]string{
			"categoryUid": ts.categoryUidString,
			"linkUid": ts.linkUidString,
		})

	ts.Equal(
		res.StatusCode(),
		204,
		"Tentar remover um link de uma categoria, ambos relacionados, deveria retornar status 204")

	ts.Empty(
		res.Body(),
		"Tentar remover um link de uma categoria, ambos relacionados, não deveria retornar nada")
}

func (ts *RemoveLinkFromCategoryTestSuite) TestRemoveLinkFromCategory_Error() {
	ts.post(
		map[string]string{
			"categoryUid": ts.categoryUidString},
		map[string]string{
			"linkUid": ts.linkUidString,
			"isMain": "true",
		})

	res, _ := ts.delete(
		map[string]string{
			"categoryUid": domain.MockCategory.Name,
			"linkUid": domain.MockCategory.Name,
		})

	ts.Equal(
		res.StatusCode(),
		400,
		"Tentar remover um link de uma categoria passando parâmetros inválidos deveria retornar status 400")

	ts.ElementsMatch([]string{"categoryUid", "linkUid"}, ltests.GetErrorKeys(res.Body()))
}

func (ts *RemoveLinkFromCategoryTestSuite) TestRemoveLinkFromCategory_LinkNotExists() {
	res, _ := ts.delete(
		map[string]string{
			"categoryUid": ts.categoryUidString,
			"linkUid": domain.MockUidLink.String(),
		})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.LINK_NOT_EXISTS.Error()),
		res.Body())

	ts.Equal(
		res.StatusCode(),
		404,
		"Tentar remover um link de uma categoria que não existe deveria retornar status 404")

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar remover um link de uma categoria que não existe deveria retornar erro contendo " + domain.LINK_NOT_EXISTS.Error())
}

func (ts *RemoveLinkFromCategoryTestSuite) TestRemoveLinkFromCategory_CategoryNotExists() {
	res, _ := ts.delete(
		map[string]string{
			"categoryUid": domain.MockUidCategory.String(),
			"linkUid": ts.linkUidString,
		})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.CATEGORY_NOT_EXISTS.Error()),
		res.Body())

	ts.Equal(
		res.StatusCode(),
		404,
		"Tentar remover um link que não existe de uma categoria deveria retornar status 404")

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar remover um link que não existe de uma categoria deveria retornar erro contendo " + domain.CATEGORY_NOT_EXISTS.Error())
}

func (ts *RemoveLinkFromCategoryTestSuite) TestRemoveLinkFromCategory_NotExists() {
	res, _ := ts.delete(
		map[string]string{
			"categoryUid": ts.alternativeCategoryUidString,
			"linkUid": ts.alternativeLinkUidString,
		})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.NOT_BELONGS.Error()),
		res.Body())

	ts.Equal(
		res.StatusCode(),
		404,
		"Tentar remover um link de uma categoria, ambos não relacionados, deveria retornar status 404")

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar remover um link de uma categoria, ambos não relacionados, deveria retornar erro contendo " + domain.NOT_BELONGS.Error())
}

func TestRemoveLinkFromCategoryAllTests(t *testing.T) {
	suite.Run(t, new(RemoveLinkFromCategoryTestSuite))
}
