package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/pre_auth_code.html

const ApiCreatePreAuthCode = "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode"

func (client *WeChatClient) CreatePreAuthCode() {

}

type CreatePreAuthCodeRequest struct {
	ComponentAccessToken string `position:"query" name:"component_access_token" json:"-"`
	ComponentAppID string `position:"body" name:"component_appid" json:"component_appid"`
}
