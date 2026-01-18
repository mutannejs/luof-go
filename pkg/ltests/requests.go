package ltests

import (
    "encoding/json"

    "github.com/go-resty/resty/v2"
)

type PostFuncType func (map[string]string) (*resty.Response, error)
type GetFuncType func (map[string]string, map[string]string) (*resty.Response, error)

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

func GetGet(c *resty.Client, urlBase string) GetFuncType {
    return func(
        pathParamsMap map[string]string,
        queryParamsMap map[string]string,
    ) (*resty.Response, error) {
        return c.R().
            SetPathParams(pathParamsMap).
            SetQueryParams(queryParamsMap).
            Get(urlBase)
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

func DeleteKeyInByteSlice(value []byte, key string) []byte {
    var valueMap map[string]string

    json.Unmarshal(value, &valueMap)
    delete(valueMap, key)

    valueByteSlice, _ := json.Marshal(valueMap)

    return valueByteSlice
}
