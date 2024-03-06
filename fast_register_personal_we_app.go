package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/fastregisterpersonalweapp.html

const ApiFastRegisterPersonalWeApp = "https://api.weixin.qq.com/wxa/component/fastregisterpersonalweapp" //快速注册个人小程序

func (client *WeChatClient) FastRegisterPersonalWeApp(data *FastRegisterPersonalWeAppRequest) (*FastRegisterPersonalWeAppResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiFastRegisterPersonalWeApp).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		FastRegisterPersonalWeAppResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.FastRegisterPersonalWeAppResponse, nil
}

func (client *WeChatClient) BuildFastRegisterPersonalWeAppRequest() *FastRegisterPersonalWeAppRequest {
	return &FastRegisterPersonalWeAppRequest{
		Action: "create",
	}
}

type FastRegisterPersonalWeAppRequest struct {
	Action               string `position:"query" name:"action" json:"-"`
	ComponentAccessToken string `position:"query" name:"component_access_token" json:"-"`
	IDName               string `position:"body" name:"idname" json:"idname"`
	WxUser               string `position:"body" name:"wxuser" json:"wxuser"`
	ComponentPhone       string `position:"body" name:"component_phone" json:"component_phone,omitempty"`
}

type FastRegisterPersonalWeAppResponse struct {
	TaskID       string `json:"taskid"`
	AuthorizeURL string `json:"authorize_url"`
	Status       int64  `json:"status"`
}

func (client *WeChatClient) FastRegisterPersonalWeAppQuery(data *FastRegisterPersonalWeAppQueryRequest) (*FastRegisterPersonalWeAppQueryResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiFastRegisterPersonalWeApp).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		FastRegisterPersonalWeAppQueryResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.FastRegisterPersonalWeAppQueryResponse, nil
}

func (client *WeChatClient) BuildFastRegisterPersonalWeAppQueryRequest() *FastRegisterPersonalWeAppQueryRequest {
	return &FastRegisterPersonalWeAppQueryRequest{
		Action: "query",
	}
}

type FastRegisterPersonalWeAppQueryRequest struct {
	Action               string `position:"query" name:"action" json:"-"`
	ComponentAccessToken string `position:"query" name:"component_access_token" json:"-"`
	TaskID               string `position:"body" name:"taskid" json:"taskid"`
}

type FastRegisterPersonalWeAppQueryResponse FastRegisterPersonalWeAppResponse
