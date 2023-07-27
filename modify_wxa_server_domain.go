package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/thirdparty-management/domain-mgnt/modifyThirdpartyServerDomain.html

const ApiModifyWxaServerDomain = "https://api.weixin.qq.com/cgi-bin/component/modify_wxa_server_domain"

func (client *WeChatClient) ModifyWxaServerDomain(data *ModifyWxaServerDomainRequest) (*ModifyWxaServerDomainResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiModifyWxaServerDomain).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		ModifyWxaServerDomainResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.ModifyWxaServerDomainResponse, nil
}

func (client *WeChatClient) BuildModifyWxaServerDomainRequest() *ModifyWxaServerDomainRequest {
	return &ModifyWxaServerDomainRequest{}
}

type ModifyWxaServerDomainRequest struct {
	AccessToken              string `position:"query" name:"access_token" json:"-"`
	Action                   string `json:"action"`
	WxaServerDomain          string `json:"wxa_server_domain"`            // 以;分隔的域名
	IsModifyPublishedTogeter bool   `json:"is_modify_published_together"` // false 只改test
}

type ModifyWxaServerDomainResponse struct {
	PublishedWxaServerDomain string `json:"published_wxa_server_domain"`
	TestingWxaServerDoamin   string `json:"testing_wxa_server_domain"`
	InvalidWxaServerDomain   string `json:"invalid_wxa_server_domain"`
}
