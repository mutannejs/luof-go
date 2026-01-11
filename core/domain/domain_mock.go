package domain

var (
    // Elementos do domínio
    LinkMockMap = map[string]any{
        "url": "http://github.com/mutannejs/luof-go",
        "name": "luof",
        "description": "### luof-go repository",
        "useMarkdown": true,
    }
    CategoryMockMap = map[string]any{
        "name": "development",
        "description": "links about development",
        "useMarkdown": false,
    }
    MockLink, _ = NewLink(
        LinkMockMap["url"].(string),
        LinkMockMap["name"].(string),
        LinkMockMap["description"].(string),
        LinkMockMap["useMarkdown"].(bool),
    )
    MockCategory, _ = NewCategory(
        CategoryMockMap["name"].(string),
        CategoryMockMap["description"].(string),
        CategoryMockMap["useMarkdown"].(bool),
    )
    AlternativeMockLink, _ = NewLink(
        "http://github.com/mutannejs/luof",
        "luof",
        "luof repository",
        false,
    )
    AlternativeMockCategory, _ = NewCategory(
        "design",
        "design links",
        true,
    )

    // Identificadores dos elementos
    MockUidCategory = MockCategory.GetUid()
    MockUidLink = MockLink.GetUid()
    AlternativeMockUidLink = AlternativeMockLink.GetUid()
    AlternativeMockUidCategory = AlternativeMockCategory.GetUid()

    // Conjuntos de elementos do domínio
    MockLinks = []Link{MockLink, AlternativeMockLink}
)
