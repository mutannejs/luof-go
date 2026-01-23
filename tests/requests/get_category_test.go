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

type GetCategoryTestSuite struct {
    suite.Suite
    db *sql.DB
    c *resty.Client
    get ltests.RequestFuncType
    categoryUid string
}

func (ts *GetCategoryTestSuite) SetupSuite() {
    env, _ := lenv.LoadTest()
    urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api/categories"

    ts.c = resty.New()

    post := ltests.GetJSONPost(ts.c, urlBase)
    res, _ := post(nil, domain.MockCategoryMapRequest)

	ts.categoryUid = res.String()
    ts.get = ltests.GetGet(ts.c, urlBase + "/{categoryUid}")
}

func (ts *GetCategoryTestSuite) TestGetCategory() {
	res, _ := ts.get(map[string]string{"categoryUid": ts.categoryUid}, nil)

	mockCategoryJson, _ := json.Marshal(domain.MockCategory)
	mockCategoryJson = ltests.DeleteKeyInByteSlice(mockCategoryJson, "CreatedAt")
	resBody := ltests.DeleteKeyInByteSlice(res.Body(), "CreatedAt")

	ts.JSONEq(
		string(mockCategoryJson),
		string(resBody),
		"Tentar recuperar uma categoria passando um uuid válido deveria retornar a categoria correspondente")
}

func (ts *GetCategoryTestSuite) TestGetCategory_ParamRequired() {
	res, _ := ts.get(map[string]string{"categoryId": domain.MockUidCategory.String()}, nil)

	ts.ElementsMatch(ltests.GetErrorKeys(res.Body()), []string{"categoryUid"}) 
}

func (ts *GetCategoryTestSuite) TestGetCategory_NotExists() {
	res, _ := ts.get(map[string]string{"categoryUid": domain.MockUidCategory.String()}, nil)

	expectedMap := map[string]string{"message": domain.CATEGORY_NOT_EXISTS.Error()}
	expectedJson, _ := json.Marshal(expectedMap)

	ts.Equal(
		strings.TrimSpace(string(expectedJson)),
		strings.TrimSpace(string(res.Body())),
		"Tentar recuperar uma categoria passando um uuid inválido deveria retornar o erro " + domain.CATEGORY_NOT_EXISTS.Error())
}

func TestGetCategoryAllTests(t *testing.T) {
    suite.Run(t, new(GetCategoryTestSuite))
}
