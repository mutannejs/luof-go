package types

type SaveLink struct {
    Url string `json:"url" validate="url"`
    Name string `json:"name" validate="alphanumunicode,required,max=200"`
    Description string `json:"description" validate="alphanumunicode"`
    UseMarkdown bool `json:"useMarkdown" validate="boolean"`
}
