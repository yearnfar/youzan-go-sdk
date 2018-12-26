package lib

import (
	"encoding/json"
	"errors"
)

const GetTokenApi = "https://open.youzan.com/oauth/token"

// YZGetTokenResponse 返回数据结构
type YZGetTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresTn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

// YZGetTokenClient 客户端数据结构
type YZGetTokenClient struct {
	clientID     string
	clientSecret string
	accessToken  string
	refreshToken string
}

func NewYZGetTokenClient(clientID, clientSecret, accessToken, refreshToken string) *YZGetTokenClient {
	return &YZGetTokenClient{
		clientID:     clientID,
		clientSecret: clientSecret,
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}
}

// GetToken 获取access_token
func (c *YZGetTokenClient) GetToken(typ string, keys map[string]string, resp *YZGetTokenResponse) error {
	params := make(map[string]string)
	switch typ {
	case "oauth":
		params["grant_type"] = "authorization_code"
		params["code"] = keys["code"]
		params["redirect_uri"] = keys["redirect_uri"]

	case "refresh_token":
		params["grant_type"] = "authorization_code"
		params["refresh_token"] = keys["refresh_token"]

	case "self":
		params["grant_type"] = "silent"
		params["kdt_id"] = keys["kdt_id"]

	case "platform_init":
		params["grant_type"] = "authorize_platform"

	case "platform":
		params["grant_type"] = "authorize_platform"
		params["kdt_id"] = keys["kdt_id"]
	default:
		return errors.New("类型错误")
	}

	data, err := new(YZHttpClient).Post(GetTokenApi, params)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, resp)
}
