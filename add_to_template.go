package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/addtotemplate.html

const ApiAddToTemplate = "https://api.weixin.qq.com/wxa/addtotemplate"
const TemplateTypeNormal = 0
const TemplateTypeStandard = 1

func (client *WeChatClient) AddToTemplate(data *AddToTemplateRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiAddToTemplate).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
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

func (client *WeChatClient) BuildAddToTemplateRequest() *AddToTemplateRequest {
	return &AddToTemplateRequest{}
}

type AddToTemplateRequest struct {
	AccessToken  string `position:"query" name:"access_token" json:"-"`
	DraftID      int64 	`position:"body" name:"draft_id" json:"draft_id"`
	TemplateType int64  `position:"body" name:"template_type" json:"template_type"`
}

type AddToTemplateResponse struct {
}
