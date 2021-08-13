package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/pre_auth_code.html

const ApiCreatePreAuthCode = "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode"

func (client *WeChatClient) CreatePreAuthCode() {

}

type CreatePreAuthCodeRequest struct {
	ComponentAccessToken string `position:"query" name:"component_access_token" json:"-"`
	ComponentAppID       string `position:"body" name:"component_appid" json:"component_appid"`
}

type CreatePreAuthCodeResponse struct {
	ErrorCode    int64  `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
	PreAuthCode  string `json:"pre_auth_code"`
	ExpiresIn    int64  `json:"expires_in"` //有效期，单位:秒
}
