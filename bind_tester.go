package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_AdminManagement/Admin.html

const ApiBindTester = "https://api.weixin.qq.com/wxa/bind_tester"

func (client *WeChatClient) BindTester(data *BindTesterRequest) (*BindTesterResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiBindTester).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		BindTesterResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.BindTesterResponse, nil
}

func (client *WeChatClient) BuildBindTesterRequest() *BindTesterRequest {
	return &BindTesterRequest{}
}

type BindTesterRequest struct {
	AccessToken string `name:"access_token" json:"-"`
	WeChatID    string `name:"wechatid" json:"wechatid"`
}

type BindTesterResponse struct {
	UserString string `json:"userstr"`
}
