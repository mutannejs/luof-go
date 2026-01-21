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

type DeleteLinkTestSuite struct {
    suite.Suite
    db *sql.DB
    c *resty.Client
    delete ltests.DeleteFuncType
    post ltests.RequestFuncType
}

func (ts *DeleteLinkTestSuite) SetupSuite() {
    env, _ := lenv.LoadTest()
    urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api/links"

    ts.c = resty.New()
    ts.post = ltests.GetJSONPost(ts.c, urlBase)
    ts.delete = ltests.GetDelete(ts.c, urlBase + "/{linkUid}")

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

	expectedMap := map[string]string{"message": domain.LINK_NOT_EXISTS.Error()}
	expectedJson, _ := json.Marshal(expectedMap)

	ts.Equal(
		strings.TrimSpace(string(expectedJson)),
		strings.TrimSpace(string(res.Body())),
		"Tentar deletar um link passando um uuid inválido deveria retornar o erro " + domain.LINK_NOT_EXISTS.Error())
}

func TestDeleteLinkAllTests(t *testing.T) {
    suite.Run(t, new(DeleteLinkTestSuite))
}
