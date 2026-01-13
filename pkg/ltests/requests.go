package ltests

import (
    "github.com/go-resty/resty/v2"
)

type PostFuncType func (map[string]string) (*resty.Response, error)

func GetFormDataPost(c *resty.Client, urlBase string) PostFuncType {
    return func(formData map[string]string) (*resty.Response, error) {
        return c.R().
            SetFormData(formData).
            Post(urlBase)
    }
}

func GetJSONPost(c *resty.Client, urlBase string) PostFuncType {
    return func(dataMap map[string]string) (*resty.Response, error) {
        return c.R().
            SetBody(dataMap).
            Post(urlBase)
    }
}
