package requests

import (
	"database/sql"
	"regexp"
	"strconv"
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/ltests"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
)

const (
    LINK_TABLE_MIGRATION uint = 1764809880
)

var (
	mockLink = domain.MockLink
)

type TestSuite struct {
    suite.Suite
    db *sql.DB
    c *resty.Client
    post ltests.PostFuncType
}

func (ts *TestSuite) SetupSuite() {
    env, _ := lenv.LoadTest()
    urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api/links"

    ts.c = resty.New()
    ts.post = ltests.GetPost(ts.c, urlBase)
}

func (ts *TestSuite) TestCreateLink() {
	res, _ := ts.post(map[string]string{
		"url": mockLink.Url,
		"name": mockLink.Name,
		"description": mockLink.Description.Content,
		"useMarkdown": strconv.FormatBool(mockLink.Description.UseMarkdown),
	})

	ts.Regexp(
		regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12}"),
		res,
		"Tentar criar um link passando parâmetros válidos deveria retornar um uuid válido")
}

func (ts *TestSuite) TestCreateLink_Error() {
	res, _ := ts.post(map[string]string{
		"url": mockLink.Name, // não é URL
		// falta o parâmetro name
		"description": mockLink.Description.Content,
		"useMarkdown": mockLink.Name, // não é boolean
	})

	ts.Regexp(
		regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12}"),
		res,
		"Tentar criar um link passando parâmetros válidos deveria retornar um uuid válido")
}

func TestSqliteCategory(t *testing.T) {
    suite.Run(t, new(TestSuite))
}
