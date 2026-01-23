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

type DeleteCategoryTestSuite struct {
    suite.Suite
    db *sql.DB
    c *resty.Client
    delete ltests.DeleteFuncType
    post ltests.RequestFuncType
}

func (ts *DeleteCategoryTestSuite) SetupSuite() {
    env, _ := lenv.LoadTest()
    urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api/categories"

    ts.c = resty.New()
    ts.post = ltests.GetJSONPost(ts.c, urlBase)
    ts.delete = ltests.GetDelete(ts.c, urlBase + "/{categoryUid}")

    ts.post(nil, domain.MockCategoryMapRequest)
}

func (ts *DeleteCategoryTestSuite) TestDeleteCategory() {
    res, _ := ts.post(nil, domain.MockCategoryMapRequest)
	res, err := ts.delete(map[string]string{"categoryUid": res.String()})

	ts.NoError(err)

	ts.Empty(
		string(res.Body()),
		"Tentar deletar uma categoria passando um uuid válido não deveria retornar nada")
}

func (ts *DeleteCategoryTestSuite) TestDeleteCategory_ParamRequired() {
	res, _ := ts.delete(map[string]string{"categoryId": domain.MockUidCategory.String()})

	ts.ElementsMatch(ltests.GetErrorKeys(res.Body()), []string{"categoryUid"}) 
}

func (ts *DeleteCategoryTestSuite) TestDeleteCategory_NotExists() {
	res, _ := ts.delete(map[string]string{"categoryUid": domain.MockUidCategory.String()})

	expectedMap := map[string]string{"message": domain.CATEGORY_NOT_EXISTS.Error()}
	expectedJson, _ := json.Marshal(expectedMap)

	ts.Equal(
		strings.TrimSpace(string(expectedJson)),
		strings.TrimSpace(string(res.Body())),
		"Tentar deletar uma categoria passando um uuid inválido deveria retornar o erro " + domain.LINK_NOT_EXISTS.Error())
}

func TestDeleteCategoryAllTests(t *testing.T) {
    suite.Run(t, new(DeleteCategoryTestSuite))
}
