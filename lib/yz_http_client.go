package lib

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

type YZHttpClient struct {
}

// Get http请求 get方法
func (c *YZHttpClient) Get(api string, param map[string]string) ([]byte, error) {
	query := &strings.Builder{}
	for k, v := range param {
		if query.Len() != 0 {
			query.WriteString("&")
		}

		query.WriteString(k + "=" + v)
	}

	if query.Len() > 0 {
		api += "?" + query.String()
	}

	resp, err := http.Get(api)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// Get http请求 post方法
func (c *YZHttpClient) Post(api string, param map[string]string) (data []byte, err error) {
	body := bytes.NewBuffer(nil)
	for k, v := range param {
		if body.Len() != 0 {
			body.WriteString("&")
		}

		body.WriteString(k + "=" + v)
	}

	resp, err := http.Post(api, "multipart/form-data", body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
