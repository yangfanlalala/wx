package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/authorization_info.html

const ApiQueryAuth = "https://api.weixin.qq.com/cgi-bin/component/api_query_auth"

func (client *WeChatClient) QueryAuth() {

}

type QueryAuthRequest struct {
	ComponentAccessToken string `position:"query" name:"component_access_token" json:"-"`
	ComponentAppID       string `position:"body" name:"component_appid" json:"component_appid"`
	AuthorizationCode    string `position:"body" name:"authorization_code" json:"authorization_code"`
}

type QueryAuthResponse struct {
	AuthorizationInfo AuthorizerInfo `json:"authorization_info"`
}
