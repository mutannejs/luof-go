package domain

import (
	"strconv"
)

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

	// Corpo de requests
	MockLinkMapRequest = map[string]string{
		"url": MockLink.Url,
		"name": MockLink.Name,
		"description": MockLink.Description.Content,
		"useMarkdown": strconv.FormatBool(MockLink.Description.UseMarkdown),
	}
	MockCategoryMapRequest = map[string]string{
		"name": MockCategory.Name,
		"description": MockCategory.Description.Content,
		"useMarkdown": strconv.FormatBool(MockCategory.Description.UseMarkdown),
	}
	AlternativeMockLinkMapRequest = map[string]string{
		"url": AlternativeMockLink.Url,
		"name": AlternativeMockLink.Name,
		"description": AlternativeMockLink.Description.Content,
		"useMarkdown": strconv.FormatBool(AlternativeMockLink.Description.UseMarkdown),
	}
)
