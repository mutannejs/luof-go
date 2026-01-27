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

	ts.ElementsMatch(ltests.GetErrorKeys(res.Body()), []string{"name", "useMarkdown"}) 
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
		resBody,
		expectedJson,
		"Tentar atualizar um link passando um uuid inválido deveria retornar o erro " + domain.LINK_NOT_EXISTS.Error())
}

func TestUpdateLinkAllTests(t *testing.T) {
	suite.Run(t, new(UpdateLinkTestSuite))
}
