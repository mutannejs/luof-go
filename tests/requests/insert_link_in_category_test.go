package requests

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/ltests"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
)

type InsertLinkInCategoryTestSuite struct {
	suite.Suite
	post ltests.RequestFuncType
	linkUidString string
	alternativeLinkUidString string
	categoryUidString string
}

func (ts *InsertLinkInCategoryTestSuite) SetupSuite() {
	env, _ := lenv.LoadTest()
	urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api"

	c := resty.New()
	ts.post = ltests.GetJSONPost(c, urlBase + "/categories/{categoryUid}/links")

	postLink := ltests.GetJSONPost(c, urlBase + "/links")
	postCategory := ltests.GetJSONPost(c, urlBase + "/categories")

	resLink, _ := postLink(nil, domain.MockLinkMapRequest)
	ts.linkUidString = string(resLink.Body())

	resLink, _ = postLink(nil, domain.AlternativeMockLinkMapRequest)
	ts.alternativeLinkUidString = string(resLink.Body())

	resCategory, _ := postCategory(nil, domain.MockCategoryMapRequest)
	ts.categoryUidString = string(resCategory.Body())
}

func (ts *InsertLinkInCategoryTestSuite) TestInsertLinkInCategory() {
	res, _ := ts.post(
		map[string]string{
			"categoryUid": ts.categoryUidString},
		map[string]string{
			"linkUid": ts.linkUidString,
			"isMain": "true",
		})

	ts.Equal(
		res.StatusCode(),
		200,
		"Tentar inserir um link em uma categoria, ambos ainda não relacionados, deveria retornar status 200")

	ts.Empty(
		res.Body(),
		"Tentar inserir um link em uma categoria, ambos ainda não relacionados, não deveria retornar nada")
}

func (ts *InsertLinkInCategoryTestSuite) TestInsertLinkInCategory_Error() {
	res, _ := ts.post(
		map[string]string{
			"categoryUid": domain.MockCategory.Name},
		map[string]string{
			"linkUid": domain.MockCategory.Name,
			"isMain": domain.MockCategory.Name,
		})

	ts.Equal(
		res.StatusCode(),
		400,
		"Tentar inserir um link em uma categoria passando parâmetros inválidos deveria retornar status 400")

	ts.ElementsMatch([]string{"categoryUid"}, ltests.GetErrorKeys(res.Body()))
}

func (ts *InsertLinkInCategoryTestSuite) TestInsertLinkInCategory_LinkNotExists() {
	res, _ := ts.post(
		map[string]string{
			"categoryUid": ts.categoryUidString},
		map[string]string{
			"linkUid": domain.MockUidLink.String(),
			"isMain": "true",
		})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.LINK_NOT_EXISTS.Error()),
		res.Body())

	ts.Equal(
		res.StatusCode(),
		404,
		"Tentar inserir um link que não existe em uma categoria deveria retornar status 404")

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar inserir um link que não existe em uma categoria deveria retornar erro contendo " + domain.LINK_NOT_EXISTS.Error())
}

func (ts *InsertLinkInCategoryTestSuite) TestInsertLinkInCategory_CategoryNotExists() {
	res, _ := ts.post(
		map[string]string{
			"categoryUid": domain.MockUidCategory.String()},
		map[string]string{
			"linkUid": ts.linkUidString,
			"isMain": "true",
		})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.CATEGORY_NOT_EXISTS.Error()),
		res.Body())

	ts.Equal(
		res.StatusCode(),
		404,
		"Tentar inserir um link em uma categoria que não existe deveria retornar status 404")

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar inserir um link em uma categoria que não existe deveria retornar erro contendo " + domain.CATEGORY_NOT_EXISTS.Error())
}

func (ts *InsertLinkInCategoryTestSuite) TestInsertLinkInCategory_AlreadyExists() {
	res, _ := ts.post(
		map[string]string{
			"categoryUid": ts.categoryUidString},
		map[string]string{
			"linkUid": ts.linkUidString,
			"isMain": "true",
		})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.ALREADY_BELONGS.Error()),
		res.Body())

	ts.Equal(
		res.StatusCode(),
		409,
		"Tentar inserir um link em uma categoria, ambos já relacionados, deveria retornar status 409")

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar inserir um link em uma categoria, ambos já relacionados, deveria retornar erro contendo " + domain.ALREADY_BELONGS.Error())
}

func (ts *InsertLinkInCategoryTestSuite) TestInsertLinkInCategory_DefaultIsMain() {
	res, _ := ts.post(
		map[string]string{
			"categoryUid": ts.categoryUidString},
		map[string]string{
			"linkUid": ts.alternativeLinkUidString,
		})

	ts.Equal(
		res.StatusCode(),
		200,
		"Tentar inserir um link em uma categoria sem informar se ela é a principal deveria retornar status 200")

	ts.Empty(
		res.Body(),
		"Tentar inserir um link em uma categoria sem informar se ela é a principal não deveria retornar nada")
}

func TestInsertLinkInCategoryAllTests(t *testing.T) {
	suite.Run(t, new(InsertLinkInCategoryTestSuite))
}
