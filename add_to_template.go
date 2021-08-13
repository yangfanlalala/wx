package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/addtotemplate.html

const ApiAddToTemplate = "https://api.weixin.qq.com/wxa/addtotemplate?access_token=TOKEN"

func (client *WeChatClient) AddToTemplate() {

}

type AddToTemplateRequest struct {
	AccessToken  string `position:"query" name:"access_token" json:"-"`
	DraftID      string `position:"body" name:"draft_id" json:"draft_id"`
	TemplateType int64  `position:"body" name:"template_type" json:"template_type"`
}

type AddToTemplateResponse struct {
	ErrorCode    int64  `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}
