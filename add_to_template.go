package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/addtotemplate.html

const ApiAddToTemplate = "https://api.weixin.qq.com/wxa/addtotemplate"

func (client *WeChatClient) AddToTemplate(data *AddToTemplateRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiAddToTemplate).
		WithMethod(http.MethodPost).
		WithContentType("application/json").
		WithData(data)
	rsp := CommonResponse{}
	err := client.DoRequest(req, rsp)
	if err != nil {
		return err
	}
	if err = rsp.Error(); err != nil {
		return err
	}
	return nil
}

type AddToTemplateRequest struct {
	AccessToken  string `position:"query" name:"access_token" json:"-"`
	DraftID      string `position:"body" name:"draft_id" json:"draft_id"`
	TemplateType int64  `position:"body" name:"template_type" json:"template_type"`
}

type AddToTemplateResponse struct {
}
