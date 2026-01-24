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

type GetCategoryTestSuite struct {
    suite.Suite
    get ltests.RequestFuncType
    categoryUid string
}

func (ts *GetCategoryTestSuite) SetupSuite() {
    env, _ := lenv.LoadTest()
    urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api/categories"

    c := resty.New()

    post := ltests.GetJSONPost(c, urlBase)
    res, _ := post(nil, domain.MockCategoryMapRequest)

	ts.categoryUid = res.String()
    ts.get = ltests.GetGet(c, urlBase + "/{categoryUid}")
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

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.CATEGORY_NOT_EXISTS.Error()),
		res.Body())

	ts.Equal(
		resBody,
		expectedJson,
		"Tentar recuperar uma categoria passando um uuid inválido deveria retornar o erro " + domain.CATEGORY_NOT_EXISTS.Error())
}

func TestGetCategoryAllTests(t *testing.T) {
    suite.Run(t, new(GetCategoryTestSuite))
}
