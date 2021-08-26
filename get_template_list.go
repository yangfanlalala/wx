package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/gettemplatelist.html

const ApiGetTemplateList = "https://api.weixin.qq.com/wxa/gettemplatelist"

func (client *WeChatClient) GetTemplateList(data *GetTemplateListRequest) (*GetTemplateListResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiGetTemplateList).
		WithMethod(http.MethodGet).
		WithData(data)
	rsp := &struct {
		CommonResponse
		GetTemplateListResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.GetTemplateListResponse, nil
}

func (client *WeChatClient) BuildGetTemplateListRequest() *GetTemplateListRequest {
	return &GetTemplateListRequest{}
}

type GetTemplateListRequest struct {
	AccessToken  string `position:"query" name:"access_token" json:"-"`
	TemplateType string `position:"query" name:"template_type" json:"-"`
}

type GetTemplateListResponse struct {
	TemplateList []Template `json:"template_list"`
}
