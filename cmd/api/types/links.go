package types

type GetLink struct {
    Uid string `param:"linkId" validate:"required,uuid"`
}

type SaveLink struct {
    Url string `form:"url" validate:"url"`
    Name string `form:"name" validate:"required,max=200"`
    Description string `form:"description"`
    UseMarkdown bool `form:"useMarkdown"`
}
