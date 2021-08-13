package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/api_authorizer_token.html

const ApiAuthorizerToken = "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token"

func (client *WeChatClient) AuthorizerToken() {

}

type AuthorizerTokenRequest struct {
	ComponentAccessToken   string `position:"query" name:"component_access_token" json:"-"`
	ComponentAppID         string `position:"body" name:"component_appid" json:"component_app_id"`
	AuthorizerAppID        string `position:"body" name:"authorizer_appid" json:"authorizer_app_id"`
	AuthorizerRefreshToken string `position:"body" name:"authorizer_refresh_token" json:"authorizer_refresh_token"`
}

type AuthorizerTokenResponse struct {
	ErrorCode              int64  `json:"errcode"`
	ErrorMessage           string `json:"errmsg"`
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	ExpiresIn              int64  `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}
