package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/deletetemplate.html

const ApiDeleteTemplate = "https://api.weixin.qq.com/wxa/deletetemplate"

func (client *WeChatClient) DeleteTemplate(data *DeleteTemplateRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiDeleteTemplate).
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

type DeleteTemplateRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	TemplateID  string `position:"body" name:"template_id" json:"template_id"`
}

type DeleteTemplateResponse struct {
}
