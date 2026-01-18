package requests

import (
	"database/sql"
	"encoding/json"
	"strconv"
	"strings"
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/ltests"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
)

type GetLinkTestSuite struct {
    suite.Suite
    db *sql.DB
    c *resty.Client
    get ltests.GetFuncType
    linkUid string
}

func (ts *GetLinkTestSuite) SetupSuite() {
    env, _ := lenv.LoadTest()
    urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api/links"

    ts.c = resty.New()

    post := ltests.GetJSONPost(ts.c, urlBase)
    res, _ := post(map[string]string{
		"url": domain.MockLink.Url,
		"name": domain.MockLink.Name,
		"description": domain.MockLink.Description.Content,
		"useMarkdown": strconv.FormatBool(domain.MockLink.Description.UseMarkdown),
	})

	ts.linkUid = res.String()
    ts.get = ltests.GetGet(ts.c, urlBase + "/{linkUid}")
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

	ts.ElementsMatch(ltests.GetErrorKeys(res.Body()), []string{"linkUid"}) 
}

func (ts *GetLinkTestSuite) TestGetLink_NotExists() {
	res, _ := ts.get(map[string]string{"linkUid": domain.MockUidLink.String()}, nil)

	expectedMap := map[string]string{"message": domain.LINK_NOT_EXISTS.Error()}
	expectedJson, _ := json.Marshal(expectedMap)

	ts.Equal(
		strings.TrimSpace(string(expectedJson)),
		strings.TrimSpace(string(res.Body())),
		"Tentar recuperar um link passando um uuid válido deveria retornar o link correspondente")
}

func TestGetLinkAllTests(t *testing.T) {
    suite.Run(t, new(GetLinkTestSuite))
}
