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

type GetLinkTestSuite struct {
	suite.Suite
	get ltests.RequestFuncType
	linkUid string
}

func (ts *GetLinkTestSuite) SetupSuite() {
	env, _ := lenv.LoadTest()
	urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api/links"

	c := resty.New()

	post := ltests.GetJSONPost(c, urlBase)
	res, _ := post(nil, domain.MockLinkMapRequest)

	ts.linkUid = res.String()
	ts.get = ltests.GetGet(c, urlBase + "/{linkUid}")
}

func (ts *GetLinkTestSuite) TestGetLink() {
	res, _ := ts.get(map[string]string{"linkUid": ts.linkUid}, nil)

	mockLinkJson, _ := json.Marshal(domain.MockLink)
	mockLinkJson = ltests.DeleteKeyInByteSlice(mockLinkJson, "CreatedAt")
	resBody := ltests.DeleteKeyInByteSlice(res.Body(), "CreatedAt")

	ts.JSONEq(
		string(mockLinkJson),
		string(resBody),
		"Tentar recuperar um link passando um uuid válido deveria retornar o link correspondente")
}

func (ts *GetLinkTestSuite) TestGetLink_ParamRequired() {
	res, _ := ts.get(map[string]string{"linkId": domain.MockUidLink.String()}, nil)

	ts.ElementsMatch([]string{"linkUid"}, ltests.GetErrorKeys(res.Body())) 
}

func (ts *GetLinkTestSuite) TestGetLink_NotExists() {
	res, _ := ts.get(map[string]string{"linkUid": domain.MockUidLink.String()}, nil)

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.LINK_NOT_EXISTS.Error()),
		res.Body())

	ts.Equal(
		expectedJson,
		resBody,
		"Tentar recuperar um link passando um uuid inválido deveria retornar o erro " + domain.LINK_NOT_EXISTS.Error())
}

func TestGetLinkAllTests(t *testing.T) {
	suite.Run(t, new(GetLinkTestSuite))
}
