package types

import (
	z "github.com/Oudwins/zog"
)

type GetLink struct {
	LinkUid string
}

type SaveLink struct {
	Url string
	Name string
	Description string
	UseMarkdown bool
}

var (
	GetLinkSchema = z.Struct(z.Shape{
		"linkUid": UidValidate,
	})
	SaveLinkSchema = z.Struct(z.Shape{
		"url": z.String().URL(),
		"name": z.String().Max(200).Required(),
		"description": z.String(),
		"useMarkdown": z.Bool(),
	})
)
