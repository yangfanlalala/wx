package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/api_get_authorizer_info.html

const ApiGetAuthorizerInfo = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info"

func (client *WeChatClient) GetAuthorizerInfo() {

}

type GetAuthorizerInfoRequest struct {
	ComponentAccessToken string `position:"query" name:"component_access_token" json:"-"`
	ComponentAppID string `position:"body" name:"component_appid" json:"component_appid"`
	AuthorizerAppID string `position:"body" name:"authorizer_appid" json:"authorizer_appid"`
}