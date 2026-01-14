package ltests

import (
    "encoding/json"

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

func GetErrorKeys(res []byte) (keys []string) {
    var errors map[string]any

    if err := json.Unmarshal(res, &errors); err != nil {
        return nil
    }

    for key, _ := range errors["Errors"].(map[string]any) {
        keys = append(keys, key)
    }

    return
}
