package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/deletetemplate.html

const ApiDeleteTemplate = "https://api.weixin.qq.com/wxa/deletetemplate"

func (client WeChatClient) DeleteTemplate() {

}

type DeleteTemplateRequest struct {
	AccessToken string `position:"query" name:"access_token"`
	TemplateID string `position:"body" name:"template_id"`
}
