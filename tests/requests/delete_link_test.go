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

	ts.Empty(
		string(res.Body()),
		"Tentar deletar um link passando um uuid válido não deveria retornar nada")
}

func (ts *DeleteLinkTestSuite) TestDeleteLink_ParamRequired() {
	res, _ := ts.delete(map[string]string{"linkId": domain.MockUidLink.String()})

	ts.ElementsMatch(ltests.GetErrorKeys(res.Body()), []string{"linkUid"}) 
}

func (ts *DeleteLinkTestSuite) TestDeleteLink_NotExists() {
	res, _ := ts.delete(map[string]string{"linkUid": domain.MockUidLink.String()})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.LINK_NOT_EXISTS.Error()),
		res.Body())

	ts.Equal(
		resBody,
		expectedJson,
		"Tentar deletar um link passando um uuid inválido deveria retornar o erro " + domain.LINK_NOT_EXISTS.Error())
}

func TestDeleteLinkAllTests(t *testing.T) {
	suite.Run(t, new(DeleteLinkTestSuite))
}
