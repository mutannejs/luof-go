package requests

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/ltests"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
)

type UpdateLinkTestSuite struct {
	suite.Suite
	linkUid string
	put ltests.RequestFuncType
}

func (ts *UpdateLinkTestSuite) SetupSuite() {
	env, _ := lenv.LoadTest()
	urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api/links"

	c := resty.New()

	post := ltests.GetJSONPost(c, urlBase)
	res, _ := post(nil, domain.MockLinkMapRequest)

	ts.linkUid = res.String()
	ts.put = ltests.GetJSONPut(c, urlBase + "/{linkUid}")
}

func (ts *UpdateLinkTestSuite) TestUpdateLink() {
	res, err := ts.put(
		map[string]string{
			"linkUid": ts.linkUid},
		domain.MockLinkMapRequest)

	ts.NoError(err)

	ts.Equal(
		res.StatusCode(),
		204,
		"Tentar atualizar um link deveria retornar status 204")

	ts.Empty(
		string(res.Body()),
		"Tentar atualizar um link passando parâmetros válidos não deveria retornar nada")
}

func (ts *UpdateLinkTestSuite) TestUpdateLink_ParamRequired() {
	res, _ := ts.put(
		map[string]string{
			"linkUid": ts.linkUid},
		map[string]string{
			"url": domain.AlternativeMockLink.Url,
			// "name": domain.AlternativeMockLink.Name,
			"description": domain.AlternativeMockLink.Description.Content,
			"useMarkdown": domain.AlternativeMockLink.Description.Content,
		})

	ts.Equal(
		res.StatusCode(),
		400,
		"Tentar atualizar um link passando parâmetros inválidos deveria retornar status 400")

	ts.ElementsMatch([]string{"name", "useMarkdown"}, ltests.GetErrorKeys(res.Body()))
}

func (ts *UpdateLinkTestSuite) TestUpdateLink_NotExists() {
	res, _ := ts.put(
		map[string]string{
			"linkUid": domain.MockUidLink.String()},
		domain.MockLinkMapRequest)

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.LINK_NOT_EXISTS.Error()),
		res.Body())

	ts.Equal(
		res.StatusCode(),
		404,
		"Tentar atualizar um link que não existe deveria retornar status 404")

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar atualizar um link passando um uuid inválido deveria retornar o erro " + domain.LINK_NOT_EXISTS.Error())
}

func TestUpdateLinkAllTests(t *testing.T) {
	suite.Run(t, new(UpdateLinkTestSuite))
}
