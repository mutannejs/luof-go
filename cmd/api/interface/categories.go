package interfaces

import (
	z "github.com/Oudwins/zog"
)

type GetCategory struct {
	CategoryUid string
}

type SaveCategory struct {
	Name string
	Description string
	UseMarkdown bool
}

type RemoveSubcategory struct {
	CategoryUid string
	ChildUid string
}

type SaveSubcategory struct {
	ChildUid string
}

type GetBelongsTo struct {
	CategoryUid string
	LinkUid string
}

type CreateBelongsTo struct {
	LinkUid string
	IsMain bool
}

type UpdateBelongsTo struct {
	IsMain bool
}

var (
	GetCategorySchema = z.Struct(z.Shape{
		"categoryUid": UidValidate,
	})
	SaveCategorySchema = z.Struct(z.Shape{
		"name": z.String().Max(200).Required(),
		"description": z.String(),
		"useMarkdown": z.Bool(),
	})
	RemoveSubcategorySchema = z.Struct(z.Shape{
		"categoryUid": UidValidate,
		"childUid": UidValidate,
	})
	SaveSubcategorySchema = z.Struct(z.Shape{
		"childUid": UidValidate,
	})
	GetBelongsToSchema = z.Struct(z.Shape{
		"categoryUid": UidValidate,
		"linkUid": UidValidate,
	})
	CreateBelongsToSchema = z.Struct(z.Shape{
		"linkUid": UidValidate,
		"isMain": z.Bool(),
	})
	UpdateBelongsToSchema = z.Struct(z.Shape{
		"isMain": z.Bool().Required(),
	})
)
