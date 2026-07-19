package requests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/ltests"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
)

type GetAllRootCategoriesTestSuite struct {
	suite.Suite
	get ltests.RequestFuncType
	postCategory ltests.RequestFuncType
	postSubcategory ltests.RequestFuncType
	env map[string]string
}

func (ts *GetAllRootCategoriesTestSuite) SetupSuite() {
	ts.env, _ = lenv.LoadTest()
	urlBase := "http://localhost:" + ts.env["SERVER_PORT"] + "/api/categories"

	c := resty.New()
	ts.get = ltests.GetGet(c, urlBase)
	ts.postCategory = ltests.GetJSONPost(c, urlBase)
	ts.postSubcategory = ltests.GetJSONPost(c, urlBase + "/{categoryUid}/subcategories")
}

func (ts *GetAllRootCategoriesTestSuite) TestGetAllRootCategories() {
	res, _ := ts.get(nil, nil)

	var categoriesJson []map[string]string
	json.Unmarshal(res.Body(), &categoriesJson)

	if len(categoriesJson) > 0 {
		fmt.Println("SKIPED: O banco precisa estar vazio para rodar esse teste de forma previsível")
		return
	}

	expectedJson, resBody := ltests.TrimResponse(
		[]byte("[]"),
		res.Body())
		ts.Equal(
			expectedJson,
			resBody,
			"Tentar recuperar todas subcategorias raízes sem que existam categorias salvas deveria retornar um array vazio")

	resCategory, _ := ts.postCategory(nil, domain.MockCategoryMapRequest)
	fatherUidString := string(resCategory.Body())

	resCategory, _ = ts.postCategory(nil, domain.AlternativeMockCategoryMapRequest)
	childUidString := string(resCategory.Body())

	ts.postCategory(nil, domain.AlternativeMockCategoryMapRequest)
	ts.postSubcategory(
		map[string]string{
			"categoryUid": fatherUidString},
		map[string]string{
			"childUid": childUidString,
		})

	res, _ = ts.get(nil,nil)

	json.Unmarshal(res.Body(), &categoriesJson)

	ts.Len(
		categoriesJson,
		2,
		"Tentar recuperar todos as categorias raízes deveria retornar uma lista com e somente com todas categorias raízes")
}

func TestGetAllRootCategoriesAllTests(t *testing.T) {
	suite.Run(t, new(GetAllRootCategoriesTestSuite))
}
