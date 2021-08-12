package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/api_authorizer_token.html

const ApiAuthorizerToken = "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token"

func (client *WeChatClient) AuthorizerToken() {

}

type AuthorizerTokenRequest struct {
	ComponentAccessToken  string `position:"query" name:"component_access_token"`
	ComponentAppID string `position:"body" name:"component_appid"`
	AuthorizerAppID string `position:"body" name:"authorizer_appid"`
	AuthorizerRefreshToken string `position:"body" name:"authorizer_refresh_token"`
}
