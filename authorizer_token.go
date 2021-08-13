package wx

import (
	"net/http"
)

//api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/api_authorizer_token.html

const ApiAuthorizerToken = "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token"

//AuthorizerToken 获取授权Token
func (client *WeChatClient) AuthorizerToken(req *AuthorizerTokenRequest) (*AuthorizerTokenResponse, error) {
	req.WithMethod(http.MethodPost).WithURL(ApiAuthorizerToken).WithContentType("application/json")
	rsp := &AuthorizerTokenResponse{}
	err := client.DoRequest(req, rsp)
	return rsp, err
}

type AuthorizerTokenRequest struct {
	CommonRequest
	ComponentAccessToken   string `position:"query" name:"component_access_token" json:"-"`
	ComponentAppID         string `position:"body" name:"component_appid" json:"component_app_id"`
	AuthorizerAppID        string `position:"body" name:"authorizer_appid" json:"authorizer_app_id"`
	AuthorizerRefreshToken string `position:"body" name:"authorizer_refresh_token" json:"authorizer_refresh_token"`
}

type AuthorizerTokenResponse struct {
	CommonResponse
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	ExpiresIn              int64  `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}
