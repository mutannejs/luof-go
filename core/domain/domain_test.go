package domain

import (
    "github.com/mutannejs/luof-go/pkg/luuid"
    "fmt"
    "reflect"
    "testing"
    "time"
)

var (
    linkMock = map[string]any{
        "url": "github.com/mutannejs/luof-go",
        "name": "luof",
        "description": "### luof-go repository",
        "useMarkdown": true,
    }
    categoryMock = map[string]any{
        "name": "development",
        "description": "links about development",
        "useMarkdown": false,
    }
)

func isTimeZero(value time.Time) bool {
    var zero time.Time
    return reflect.DeepEqual(value, zero)
}

func TestNewCategory(t *testing.T) {
    category, err := NewCategory(
        categoryMock["name"].(string),
        categoryMock["description"].(string),
        categoryMock["useMarkdown"].(bool),
    )

    if err != nil ||
            luuid.IsZero(category.Uid) ||
            category.Name != categoryMock["name"] ||
            category.Description.Content != categoryMock["description"] ||
            category.Description.UseMarkdown != categoryMock["useMarkdown"] ||
            isTimeZero(category.CreatedAt) ||
            !isTimeZero(category.UpdatedAt) {
        fmt.Println(err)
        t.Fail()
    }
}

func TestNewLink(t *testing.T) {
    link, err := NewLink(
        linkMock["url"].(string),
        linkMock["name"].(string),
        linkMock["description"].(string),
        linkMock["useMarkdown"].(bool),
    )

    if err != nil ||
            luuid.IsZero(link.Uid) ||
            link.Url != linkMock["url"] ||
            link.Name != linkMock["name"] ||
            link.Description.Content != linkMock["description"] ||
            link.Description.UseMarkdown != linkMock["useMarkdown"] ||
            isTimeZero(link.CreatedAt) ||
            !isTimeZero(link.UpdatedAt) {
        fmt.Println(err)
        t.Fail()
    }
}
