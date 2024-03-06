package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/Fast_Registration_Interface_document.html

const ApiFastRegisterWeApp = "https://api.weixin.qq.com/cgi-bin/component/fastregisterweapp"

// FastRegisterWeApp 快速注册小程序
func (client *WeChatClient) FastRegisterWeApp(data *FastRegisterWeAppRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiFastRegisterWeApp).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &CommonResponse{}
	if err := client.DoRequest(req, rsp); err != nil {
		return err
	}
	if err := rsp.Error(); err != nil {
		return err
	}
	return nil
}

func (client *WeChatClient) BuildFastRegisterWeAppRequest() *FastRegisterWeAppRequest {
	return &FastRegisterWeAppRequest{
		Action: "create",
	}
}

type FastRegisterWeAppRequest struct {
	Action               string `position:"query" name:"action" json:"-"`
	ComponentAccessToken string `position:"query" name:"component_access_token" json:"-"`
	Name                 string `position:"body" name:"name" json:"name"`
	Code                 string `position:"body" name:"code" json:"code"`
	CodeType             int64  `position:"body" name:"code_type" json:"code_type"`
	LegalPersonaWechat   string `position:"body" name:"legal_persona_wechat" json:"legal_persona_wechat"`
	LegalPersonaName     string `position:"body" name:"legal_persona_name" json:"legal_persona_name"`
	ComponentPhone       string `position:"body" name:"component_phone" json:"component_phone,omitempty"`
}

type FastRegisterWeAppResponse struct {
}
