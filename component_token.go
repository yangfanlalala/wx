package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/component_access_token.html

const ApdComponentToken = "https://api.weixin.qq.com/cgi-bin/component/api_component_token"

func (client *WeChatClient) ComponentToken() {

}

type ComponentTokenRequest struct {
	ComponentAppID        string `position:"body" name:"component_appid" json:"component_appid"`
	ComponentAppSecret    string `position:"body" name:"component_appsecret" json:"component_appsecret"`
	ComponentVerifyTicket string `position:"body" name:"component_verify_ticket" json:"component_verify_ticket"`
}

type ComponentTokenResponse struct {
	CommonResponse
	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn            int64  `json:"expires_in"`
}
