package wx

import "net/http"

// app document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/weapp-wxverify/secwxaapi_wxaauth.html

const ApiSecWxaReauth = " https://api.weixin.qq.com/wxa/sec/reauth"

func (client *WeChatClient) SecWxaReauth(data *SecWxaReauthRequest) (*SecWxaReauthResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiSecWxaReauth).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		SecWxaReauthResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.SecWxaReauthResponse, nil
}

func (client *WeChatClient) BuildSecWxaReauthRequest() *SecWxaReauthRequest {
	return &SecWxaReauthRequest{}
}

type SecWxaReauthRequest SecWxaAuthRequest

type SecWxaReauthResponse SecWxaAuthResponse
