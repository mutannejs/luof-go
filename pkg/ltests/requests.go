package ltests

import (
    "encoding/json"

    "github.com/go-resty/resty/v2"
)

type RequestFuncType func (map[string]string, map[string]string) (*resty.Response, error)
type DeleteFuncType func (map[string]string) (*resty.Response, error)

func GetFormDataPost(c *resty.Client, urlBase string) RequestFuncType {
    return func(
        pathParamsMap map[string]string,
        formData map[string]string,
    ) (*resty.Response, error) {
        return c.R().
            SetPathParams(pathParamsMap).
            SetFormData(formData).
            Post(urlBase)
    }
}

func GetJSONPost(c *resty.Client, urlBase string) RequestFuncType {
    return func(
        pathParamsMap map[string]string,
        dataMap map[string]string,
    ) (*resty.Response, error) {
        return c.R().
            SetPathParams(pathParamsMap).
            SetBody(dataMap).
            Post(urlBase)
    }
}

func GetJSONPut(c *resty.Client, urlBase string) RequestFuncType {
    return func(
        pathParamsMap map[string]string,
        dataMap map[string]string,
    ) (*resty.Response, error) {
        return c.R().
            SetPathParams(pathParamsMap).
            SetBody(dataMap).
            Put(urlBase)
    }
}

func GetGet(c *resty.Client, urlBase string) RequestFuncType {
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

func GetDelete(c *resty.Client, urlBase string) DeleteFuncType {
    return func(
        pathParamsMap map[string]string,
    ) (*resty.Response, error) {
        return c.R().
            SetPathParams(pathParamsMap).
            Delete(urlBase)
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
