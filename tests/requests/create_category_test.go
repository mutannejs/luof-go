package requests

import (
	"testing"

	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/ltests"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
)

type CreateCategoryTestSuite struct {
	suite.Suite
	post ltests.RequestFuncType
}

func (ts *CreateCategoryTestSuite) SetupSuite() {
	env, _ := lenv.LoadTest()
	urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api/categories"

	c := resty.New()
	ts.post = ltests.GetJSONPost(c, urlBase)
}

func (ts *CreateCategoryTestSuite) TestCreateCategory() {
	res, _ := ts.post(nil, domain.MockCategoryMapRequest)

	ts.Regexp(
		ltests.UidRegex,
		res,
		"Tentar criar uma categoria passando parâmetros válidos deveria retornar um uuid válido")
}

func (ts *CreateCategoryTestSuite) TestCreateCategory_Error() {
	res, _ := ts.post(
		nil,
		map[string]string{
			// falta o parâmetro name
			"description": domain.MockCategory.Description.Content,
			"useMarkdown": domain.MockCategory.Name, // não é boolean
		})

	ts.ElementsMatch([]string{"name", "useMarkdown"}, ltests.GetErrorKeys(res.Body())) 
}

func TestCreateCategoryAllTests(t *testing.T) {
	suite.Run(t, new(CreateCategoryTestSuite))
}
