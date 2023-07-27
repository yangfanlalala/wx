package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/thirdparty-management/domain-mgnt/modifyThirdpartyJumpDomain.html
const ApiModifyWxaJumpDomain = "https://api.weixin.qq.com/cgi-bin/component/modify_wxa_jump_domain"

func (client *WeChatClient) ModifyWxaJumpDomain(data *ModifyWxaJumpDomainRequest) (*ModifyWxaJumpDomainResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiModifyWxaJumpDomain).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		ModifyWxaJumpDomainResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.ModifyWxaJumpDomainResponse, nil
}

func (client *WeChatClient) BuildModifyWxaJumpDomain() *ModifyWxaJumpDomainRequest {
	return &ModifyWxaJumpDomainRequest{}
}

type ModifyWxaJumpDomainRequest struct {
	AccessToken               string `position:"query" name:"access_token" json:"-"`
	Action                    string `json:"action"`
	WxaJumpH5Domain           string `json:"wxa_jump_h5_domain"`
	IsModifyPublishedTogether bool   `json:"is_modify_published_together"`
}

type ModifyWxaJumpDomainResponse struct {
	PublishedWxaJumpH5Domain string `json:"published_wxa_jump_h5_domain"`
	TestingWxaJumpH5Domain   string `json:"testing_wxa_jump_h5_domain"`
	InvalidWxaJumpH5Domain   string `json:"invalid_wxa_jump_h5_domain"`
}
