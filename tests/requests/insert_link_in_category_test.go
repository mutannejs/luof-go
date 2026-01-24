package requests

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/mutannejs/luof-go/adapters/sqlite"
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/ltests"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
)

type InsertLinkInCategoryTestSuite struct {
    suite.Suite
    post ltests.RequestFuncType
    linkUidString string
    categoryUidString string
}

func (ts *InsertLinkInCategoryTestSuite) SetupSuite() {
    env, _ := lenv.LoadTest()
    db, _ := sqlite.GetConnection(env)
    urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api"

    c := resty.New()
    ts.post = ltests.GetJSONPost(c, urlBase + "/categories/{categoryUid}/links")

    postLink := ltests.GetJSONPost(c, urlBase + "/links")
    postCategory := ltests.GetJSONPost(c, urlBase + "/categories")

    ltests.CleanTable(db, "belongs_to")
    db.Close()

    resLink, _ := postLink(nil, domain.MockLinkMapRequest)
    ts.linkUidString = string(resLink.Body())

    resCategory, _ := postCategory(nil, domain.MockCategoryMapRequest)
    ts.categoryUidString = string(resCategory.Body())
}

func (ts *InsertLinkInCategoryTestSuite) TestInsertLinkInCategory() {
	res, _ := ts.post(
		map[string]string{
			"categoryUid": ts.categoryUidString},
		map[string]string{
			"linkUid": ts.linkUidString,
			"isMain": "true",
		})

	ts.Empty(
		res.Body(),
		"Tentar inserir um link em uma categoria, ambos ainda não relacionados não deveria retornar nada")
}

func (ts *InsertLinkInCategoryTestSuite) TestInsertLinkInCategory_Error() {
	res, _ := ts.post(
		map[string]string{
			"categoryUid": domain.MockCategory.Name},
		map[string]string{
			"linkUid": domain.MockCategory.Name,
			"isMain": domain.MockCategory.Name,
		})

	ts.ElementsMatch(ltests.GetErrorKeys(res.Body()), []string{"categoryUid"}) 
}

func (ts *InsertLinkInCategoryTestSuite) TestInsertLinkInCategory_AlreadyExists() {
	res, _ := ts.post(
		map[string]string{
			"categoryUid": ts.categoryUidString},
		map[string]string{
			"linkUid": ts.linkUidString,
			"isMain": "true",
		})

	expectedMap := map[string]string{"message": domain.ALREADY_BELONGS.Error()}
	expectedJson, _ := json.Marshal(expectedMap)

	ts.Equal(
		strings.TrimSpace(string(res.Body())),
		strings.TrimSpace(string(expectedJson)),
        "Tentar inserir um link em uma categoria, ambos já relacionados, deveria retornar erro contendo " + domain.ALREADY_BELONGS.Error()) 
}

func TestInsertLinkInCategoryAllTests(t *testing.T) {
    suite.Run(t, new(InsertLinkInCategoryTestSuite))
}
