package types

import (
    z "github.com/Oudwins/zog"
)

type GetCategory struct {
    CategoryUid string
}

var GetCategorySchema = z.Struct(z.Shape{
    "categoryUid": z.String().UUID().Required(),
})

type SaveCategory struct {
    Name string
    Description string
    UseMarkdown bool
}

var SaveCategorySchema = z.Struct(z.Shape{
    "name": z.String().Max(200).Required(),
    "description": z.String(),
    "useMarkdown": z.Bool(),
})
