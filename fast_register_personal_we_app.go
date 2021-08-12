package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/fastregisterpersonalweapp.html

const ApiFastRegisterPersonalWeApp = "https://api.weixin.qq.com/wxa/component/fastregisterpersonalweapp?action=query&component_access_token=ACCESS_TOKEN" //快速注册个人小程序

func (client *WeChatClient) FastRegisterPersonalWeApp() {

}

type FastRegisterPersonalWeAppRequest struct {
	Action               string `position:"query" name:"action" json:"-"`
	ComponentAccessToken string `position:"query" name:"component_access_token" json:"-"`
	IDName               string `position:"body" name:"idname" json:"idname"`
	WxUser               string `position:"body" name:"wxuser" json:"wxuser"`
	ComponentPhone       string `position:"body" name:"component_phone" json:"component_phone"`
}
