package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/component_access_token.html

const ApdComponentToken = "https://api.weixin.qq.com/cgi-bin/component/api_component_token"

func (client *WeChatClient) ComponentToken() {

}

type ComponentTokenRequest struct {
	ComponentAppID string `position:"body" name:"component_appid"`
	ComponentAppSecret string `position:"body" name:"component_appsecret"`
	ComponentVerifyTicket string `position:"body" name:"component_verify_ticket"`
}
