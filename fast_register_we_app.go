package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/Fast_Registration_Interface_document.html

const ApiFastRegisterWeApp = "https://api.weixin.qq.com/cgi-bin/component/fastregisterweapp?action=create&component_access_token=TOKEN"

// FastRegisterWeApp 快速注册小程序
func (client WeChatClient) FastRegisterWeApp() {

}

type FastRegisterWeAppRequest struct {
	Action             string `position:"query" name:"action"`
	AccessCode         string `position:"query" name:"access_code"`
	Name               string `position:"body" name:"name"`
	Code               string `position:"body" name:"code"`
	CodeType           int64  `position:"body" name:"code_type"`
	LegalPersonaWechat string `position:"body" name:"legal_persona_wechat"`
	LegalPersonaName   string `position:"body" name:"legal_persona_name"`
	ComponentPhone     string `position:"body" name:"component_phone"`
}
