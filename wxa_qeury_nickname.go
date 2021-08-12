package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/api_wxa_querynickname.html

const ApiWxaQueryNickname = "https://api.weixin.qq.com/wxa/api_wxa_querynickname"

func (client *WeChatClient) WxaQueryNickname() {

}

type WxaQueryNickname struct {
	AccessToken string `position:"access_token" name:"access_token" json:"-"`
	AuditID string `position:"body" name:"audit_id" json:"audit_id"`
}