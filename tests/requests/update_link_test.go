package requests

import (
	"database/sql"
	"encoding/json"
	"strings"
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/ltests"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
)

type UpdateLinkTestSuite struct {
    suite.Suite
    db *sql.DB
    c *resty.Client
    linkUid string
    put ltests.RequestFuncType
}

func (ts *UpdateLinkTestSuite) SetupSuite() {
    env, _ := lenv.LoadTest()
    urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api/links"

    ts.c = resty.New()

    post := ltests.GetJSONPost(ts.c, urlBase)
    res, _ := post(nil, domain.MockLinkMapRequest)

	ts.linkUid = res.String()
    ts.put = ltests.GetJSONPut(ts.c, urlBase + "/{linkUid}")
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

	expectedMap := map[string]string{"message": domain.LINK_NOT_EXISTS.Error()}
	expectedJson, _ := json.Marshal(expectedMap)

	ts.Equal(
		strings.TrimSpace(string(expectedJson)),
		strings.TrimSpace(string(res.Body())),
		"Tentar atualizar um link passando um uuid inválido deveria retornar o erro " + domain.LINK_NOT_EXISTS.Error())
}

func TestUpdateLinkAllTests(t *testing.T) {
    suite.Run(t, new(UpdateLinkTestSuite))
}
