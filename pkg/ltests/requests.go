package ltests

import (
    "github.com/go-resty/resty/v2"
)

type PostFuncType func (map[string]string) (*resty.Response, error)

func GetPost(c *resty.Client, urlBase string) PostFuncType {
    return func(formData map[string]string) (*resty.Response, error) {
        return c.R().
            SetFormData(formData).
            Post(urlBase)
    }
}
