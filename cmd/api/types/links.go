package types

import (
    z "github.com/Oudwins/zog"
)

type GetLink struct {
    Uid string
}

var GetLinkSchema = z.Struct(z.Shape{
    "linkUid": z.String().UUID().Required(),
})

type SaveLink struct {
    Url string
    Name string
    Description string
    UseMarkdown bool
}

var SaveLinkSchema = z.Struct(z.Shape{
    "url": z.String().
        URL(z.Message("'url' is not a valid URL")),
    "name": z.String().
        Max(200, z.Message("'name' must be less than 200 characters")).
        Required(),
    "description": z.String(),
    "useMarkdown": z.Bool(),
})
