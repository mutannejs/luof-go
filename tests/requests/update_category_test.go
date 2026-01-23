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

type UpdateCategoryTestSuite struct {
    suite.Suite
    db *sql.DB
    c *resty.Client
    categoryUid string
    put ltests.RequestFuncType
}

func (ts *UpdateCategoryTestSuite) SetupSuite() {
    env, _ := lenv.LoadTest()
    urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api/categories"

    ts.c = resty.New()

    post := ltests.GetJSONPost(ts.c, urlBase)
    res, _ := post(nil, domain.MockCategoryMapRequest)

	ts.categoryUid = res.String()
    ts.put = ltests.GetJSONPut(ts.c, urlBase + "/{categoryUid}")
}

func (ts *UpdateCategoryTestSuite) TestUpdateCategory() {
	res, err := ts.put(
		map[string]string{
			"categoryUid": ts.categoryUid},
		domain.MockCategoryMapRequest)

	ts.NoError(err)

	ts.Empty(
		string(res.Body()),
		"Tentar atualizar uma categoria passando parâmetros válidos não deveria retornar nada")
}

func (ts *UpdateCategoryTestSuite) TestUpdateCategory_ParamRequired() {
	res, _ := ts.put(
		map[string]string{
			"categoryUid": ts.categoryUid},
		map[string]string{
			// "name": domain.AlternativeMockCategory.Name,
			"description": domain.AlternativeMockCategory.Description.Content,
			"useMarkdown": domain.AlternativeMockCategory.Description.Content,
		})

	ts.ElementsMatch(ltests.GetErrorKeys(res.Body()), []string{"name", "useMarkdown"}) 
}

func (ts *UpdateCategoryTestSuite) TestUpdateCategory_NotExists() {
	res, _ := ts.put(
		map[string]string{
			"categoryUid": domain.MockUidCategory.String()},
		domain.MockCategoryMapRequest)

	expectedMap := map[string]string{"message": domain.CATEGORY_NOT_EXISTS.Error()}
	expectedJson, _ := json.Marshal(expectedMap)

	ts.Equal(
		strings.TrimSpace(string(expectedJson)),
		strings.TrimSpace(string(res.Body())),
		"Tentar atualizar uma categoria passando um uuid inválido deveria retornar o erro " + domain.CATEGORY_NOT_EXISTS.Error())
}

func TestUpdateCategoryAllTests(t *testing.T) {
    suite.Run(t, new(UpdateCategoryTestSuite))
}
