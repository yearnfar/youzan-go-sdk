package lib

import (
	"encoding/json"
	"errors"
	"strings"
)

const ApiURL = "https://open.youzan.com/api/oauthentry/"

// YZTokenClient 无sign验证
type YZTokenClient struct {
	accessToken string // 验证Token
	apiVersion  string // 接口版本
}

// Post post方法
func (c *YZTokenClient) Post(method string, params map[string]string, result interface{}) (err error) {
	api := c.MakeURL(method)
	pairs, err := c.MakeParam(method, params)
	if err != nil {
		return
	}

	data, err := new(YZHttpClient).Post(api, pairs)
	if err != nil {
		return
	}

	return json.Unmarshal(data, result)
}

// Get get方法
func (c *YZTokenClient) Get(method string, params map[string]string, result interface{}) (err error) {
	api := c.MakeURL(method)
	pairs, err := c.MakeParam(method, params)
	if err != nil {
		return
	}

	data, err := new(YZHttpClient).Get(api, pairs)
	if err != nil {
		return
	}

	return json.Unmarshal(data, result)
}

// MakeURL 生成url
func (c *YZTokenClient) MakeURL(method string) string {
	arr := strings.Split(method, ".")
	method = strings.Join(arr[:len(arr)-1], ".") + "/" + c.apiVersion + "/" + arr[len(arr)-1]
	return ApiURL + method
}

// MakeParam 生成参数
func (c *YZTokenClient) MakeParam(method string, params map[string]string) (map[string]string, error) {
	pair := c.commonParam(method)
	for k, v := range params {
		if _, ok := pair[k]; ok {
			return nil, errors.New("参数名冲突")
		}

		pair[k] = v
	}
	return pair, nil
}

// 公共参数
func (c *YZTokenClient) commonParam(method string) map[string]string {
	pair := map[string]string{}
	pair[TokenKey] = c.accessToken
	pair[MethodKey] = method
	return pair
}
