package wx

import (
	"log"
	"net/http"
)

//api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/api_authorizer_token.html

const ApiAuthorizerToken = "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token"

//AuthorizerToken 获取授权Token
func (client *WeChatClient) AuthorizerToken(data *AuthorizerTokenRequest) (*AuthorizerTokenResponse, error) {
	req := &CommonRequest{}
	req.WithMethod(http.MethodPost).
		WithURL(ApiAuthorizerToken).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		AuthorizerTokenResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		log.Println(rsp.ErrorMessage)
		return nil, err
	}
	return &rsp.AuthorizerTokenResponse, nil
}

type AuthorizerTokenRequest struct {
	ComponentAccessToken   string `position:"query" name:"component_access_token" json:"-"`
	ComponentAppID         string `position:"body" name:"component_appid" json:"component_app_id"`
	AuthorizerAppID        string `position:"body" name:"authorizer_appid" json:"authorizer_app_id"`
	AuthorizerRefreshToken string `position:"body" name:"authorizer_refresh_token" json:"authorizer_refresh_token"`
}

type AuthorizerTokenResponse struct {
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	ExpiresIn              int64  `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}
