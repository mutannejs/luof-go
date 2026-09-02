package requests

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/ltests"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
)

type DeleteLinkTestSuite struct {
	suite.Suite
	delete ltests.DeleteFuncType
	post ltests.RequestFuncType
}

func (ts *DeleteLinkTestSuite) SetupSuite() {
	env, _ := lenv.LoadTest()
	urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api/links"

	c := resty.New()
	ts.post = ltests.GetJSONPost(c, urlBase)
	ts.delete = ltests.GetDelete(c, urlBase + "/{linkUid}")

	ts.post(nil, domain.MockLinkMapRequest)
}

func (ts *DeleteLinkTestSuite) TestDeleteLink() {
	res, _ := ts.post(nil, domain.MockLinkMapRequest)
	res, err := ts.delete(map[string]string{"linkUid": res.String()})

	ts.NoError(err)

	ts.Equal(
		res.StatusCode(),
		204,
		"Tentar deletar um link passando um uuid válido deveria retornar status 204")

	ts.Empty(
		string(res.Body()),
		"Tentar deletar um link passando um uuid válido não deveria retornar nada")
}

func (ts *DeleteLinkTestSuite) TestDeleteLink_ParamRequired() {
	res, _ := ts.delete(map[string]string{"linkId": domain.MockUidLink.String()})

	ts.Equal(
		res.StatusCode(),
		400,
		"Tentar deletar um link passando parâmetros inválidos deveria retornar status 400")

	ts.ElementsMatch([]string{"linkUid"}, ltests.GetErrorKeys(res.Body()))
}

func (ts *DeleteLinkTestSuite) TestDeleteLink_NotExists() {
	res, _ := ts.delete(map[string]string{"linkUid": domain.MockUidLink.String()})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.LINK_NOT_EXISTS),
		res.Body())

	ts.Equal(
		res.StatusCode(),
		404,
		"Tentar deletar um link que não existe deveria retornar status 404")

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar deletar um link passando um uuid inválido deveria retornar o erro " + domain.LINK_NOT_EXISTS)
}

func TestDeleteLinkAllTests(t *testing.T) {
	suite.Run(t, new(DeleteLinkTestSuite))
}
