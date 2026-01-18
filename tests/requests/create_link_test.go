package requests

import (
	"database/sql"
	"strconv"
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/ltests"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
)

type CreateLinkTestSuite struct {
    suite.Suite
    db *sql.DB
    c *resty.Client
    post ltests.PostFuncType
}

func (ts *CreateLinkTestSuite) SetupSuite() {
    env, _ := lenv.LoadTest()
    urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api/links"

    ts.c = resty.New()
    ts.post = ltests.GetJSONPost(ts.c, urlBase)
}

func (ts *CreateLinkTestSuite) TestCreateLink() {
	res, _ := ts.post(map[string]string{
		"url": domain.MockLink.Url,
		"name": domain.MockLink.Name,
		"description": domain.MockLink.Description.Content,
		"useMarkdown": strconv.FormatBool(domain.MockLink.Description.UseMarkdown),
	})

	ts.Regexp(
		ltests.UidRegex,
		res,
		"Tentar criar um link passando parâmetros válidos deveria retornar um uuid válido")
}

func (ts *CreateLinkTestSuite) TestCreateLink_Error() {
	res, _ := ts.post(map[string]string{
		"url": domain.MockLink.Name, // não é URL
		// falta o parâmetro name
		"description": domain.MockLink.Description.Content,
		"useMarkdown": domain.MockLink.Name, // não é boolean
	})

	ts.ElementsMatch(ltests.GetErrorKeys(res.Body()), []string{"url", "name", "useMarkdown"}) 
}

func TestCreateLinkAllTests(t *testing.T) {
    suite.Run(t, new(CreateLinkTestSuite))
}
