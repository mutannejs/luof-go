package requests

import (
	"testing"

	"github.com/mutannejs/luof-go/adapters/sqlite"
	"github.com/mutannejs/luof-go/core/domain"
	"github.com/mutannejs/luof-go/pkg/lenv"
	"github.com/mutannejs/luof-go/pkg/ltests"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
)

type RemoveLinkFromCategoryTestSuite struct {
    suite.Suite
    post ltests.RequestFuncType
    delete ltests.DeleteFuncType
    linkUidString string
    categoryUidString string
}

func (ts *RemoveLinkFromCategoryTestSuite) SetupSuite() {
    env, _ := lenv.LoadTest()
    db, _ := sqlite.GetConnection(env)
    urlBase := "http://localhost:" + env["SERVER_PORT"] + "/api"

    c := resty.New()
    ts.post = ltests.GetJSONPost(c, urlBase + "/categories/{categoryUid}/links")
    ts.delete = ltests.GetDelete(c, urlBase + "/categories/{categoryUid}/links/{linkUid}")

    postLink := ltests.GetJSONPost(c, urlBase + "/links")
    postCategory := ltests.GetJSONPost(c, urlBase + "/categories")

    ltests.CleanTable(db, "belongs_to")
    db.Close()

    resLink, _ := postLink(nil, domain.MockLinkMapRequest)
    ts.linkUidString = string(resLink.Body())

    resCategory, _ := postCategory(nil, domain.MockCategoryMapRequest)
    ts.categoryUidString = string(resCategory.Body())
}

func (ts *RemoveLinkFromCategoryTestSuite) TestRemoveLinkFromCategory() {
	ts.post(
		map[string]string{
			"categoryUid": ts.categoryUidString},
		map[string]string{
			"linkUid": ts.linkUidString,
			"isMain": "true",
		})

	res, _ := ts.delete(
		map[string]string{
			"categoryUid": ts.categoryUidString,
			"linkUid": ts.linkUidString,
		})

	ts.Empty(
		res.Body(),
		"Tentar remover um link de uma categoria, ambos relacionados, não deveria retornar nada")
}

func (ts *RemoveLinkFromCategoryTestSuite) TestRemoveLinkFromCategory_Error() {
	ts.post(
		map[string]string{
			"categoryUid": ts.categoryUidString},
		map[string]string{
			"linkUid": ts.linkUidString,
			"isMain": "true",
		})

	res, _ := ts.delete(
		map[string]string{
			"categoryUid": domain.MockCategory.Name,
			"linkUid": domain.MockCategory.Name,
		})

	ts.ElementsMatch(ltests.GetErrorKeys(res.Body()), []string{"categoryUid", "linkUid"}) 
}

func (ts *RemoveLinkFromCategoryTestSuite) TestRemoveLinkFromCategory_NotExists() {
	res, _ := ts.delete(
		map[string]string{
			"categoryUid": domain.MockUidCategory.String(),
			"linkUid": domain.MockUidLink.String(),
		})

	expectedJson, resBody := ltests.TrimResponse(
		ltests.GetResponseMessage(domain.NOT_BELONGS.Error()),
		res.Body())

	ts.Equal(
		resBody,
		expectedJson,
        "Tentar remover um link de uma categoria, ambos não relacionados, deveria retornar erro contendo " + domain.NOT_BELONGS.Error()) 
}

func TestRemoveLinkFromCategoryAllTests(t *testing.T) {
    suite.Run(t, new(RemoveLinkFromCategoryTestSuite))
}
