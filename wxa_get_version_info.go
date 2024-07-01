package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/code-management/getVersionInfo.html

const ApiWxaGetVersionInfo = "https://api.weixin.qq.com/wxa/getversioninfo"

func (client *WeChatClient) WxaGetVersionInfo(data *WxaGetVersionInfoRequest) (*WxaGetVersionInfoResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiWxaGetVersionInfo).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		WxaGetVersionInfoResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.WxaGetVersionInfoResponse, nil
}

func (client *WeChatClient) BuildWxaGetVersionInfoRequest() *WxaGetVersionInfoRequest {
	return &WxaGetVersionInfoRequest{}
}

type WxaGetVersionInfoRequest struct {
	ComponentAccessToken string `position:"query" name:"component_access_token" json:"-"`
	ComponentAppID       string `position:"body" name:"component_appid" json:"component_appid"`
	AuthorizerAppID      string `position:"body" name:"authorizer_appid" json:"authorizer_appid"`
}

type WxaGetVersionInfoResponse struct {
}
