package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/beta_Mini_Programs/fastverify.html

const ApiVerifyBetaWeApp = "https://api.weixin.qq.com/wxa/verifybetaweapp?access_token=ACCESS_TOKEN" //快速认证试用小程序

func (client *WeChatClient) VerifyBetaWeApp() {

}

type VerifyBetaWeAppRequest struct {
	AccessToken string              `position:"query" name:"access_token" json:"-"`
	VerifyInfo  VerifyBetaWeAppInfo `position:"body" name:"verify_info" json:"verify_info"`
}

type VerifyBetaWeAppInfo struct {
	EnterpriseName     string `name:"enterprise_name" json:"enterprise_name"`
	Code               string `name:"code" json:"code"`
	CodeType           int64  `name:"code_type" json:"code_type"`
	LegalPersonaWechat string `name:"legal_persona_wechat" json:"legal_persona_wechat"`
	LegalPersonaName   string `name:"legal_persona_name" json:"legal_persona_name"`
	LegalPersonaIDCard string `name:"legal_persona_idcard" json:"legal_persona_idcard"`
	ComponentPhone     string `name:"component_phone" json:"component_phone"`
}

type VerifyBetaWeAppResponse struct {
CommonResponse
}
