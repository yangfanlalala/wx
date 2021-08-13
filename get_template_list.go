package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/gettemplatelist.html

const ApiGetTemplateList = "https://api.weixin.qq.com/wxa/gettemplatelist?access_token=ACCESS_TOKEN"

func (client *WeChatClient) GetTemplateList() {

}

type GetTemplateListRequest struct {
	AccessToken  string `position:"query" name:"access_token" json:"-"`
	TemplateType string `position:"query" name:"template_type" json:"-"`
}

type GetTemplateListResponse struct {
CommonResponse
	TemplateList []struct {
		CreateTime   int64  `json:"create_time"`
		UserVersion  string `json:"user_version"`
		UserDesc     string `json:"user_desc"`
		TemplateID   int64  `json:"template_id"`
		TemplateType int64  `json:"template_type"`
	} `json:"template_list"`
}
