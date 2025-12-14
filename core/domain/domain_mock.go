package domain

var (
    // Elementos do domínio
    linkMockMap = map[string]any{
        "url": "github.com/mutannejs/luof-go",
        "name": "luof",
        "description": "### luof-go repository",
        "useMarkdown": true,
    }
    categoryMockMap = map[string]any{
        "name": "development",
        "description": "links about development",
        "useMarkdown": false,
    }
    MockLink, _ = NewLink(
        linkMockMap["url"].(string),
        linkMockMap["name"].(string),
        linkMockMap["description"].(string),
        linkMockMap["useMarkdown"].(bool),
    )
    MockCategory, _ = NewCategory(
        categoryMockMap["name"].(string),
        categoryMockMap["description"].(string),
        categoryMockMap["useMarkdown"].(bool),
    )
    AlternativeMockLink, _ = NewLink(
        "github.com/mutannejs/luof",
        "luof",
        "luof repository",
        false,
    )
    AlternativeMockCategory, _ = NewCategory(
        "design",
        "design links",
        true,
    )

    // Conjuntos de elementos do domínio
    MockLinks = []Link{MockLink, AlternativeMockLink}

    // Identificadores dos elementos
    MockUidCategory = MockCategory.GetUid()
    MockUidLink = MockLink.GetUid()
)
