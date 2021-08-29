package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/revertcoderelease.html

const ApiWxaRevertCodeRelease = "https://api.weixin.qq.com/wxa/revertcoderelease"

func (client *WeChatClient) WxaRevertCodeRelease(data *WxaRevertCodeReleaseRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiWxaRevertCodeRelease).
		WithMethod(http.MethodGet).
		WithData(data)
	rsp := &CommonResponse{}
	if err := client.DoRequest(req, rsp); err != nil {
		return err
	}
	if err := rsp.Error(); err != nil {
		return err
	}
	return nil
}

func (client *WeChatClient) BuildWxaRevertCodeReleaseRequest() *WxaRevertCodeReleaseRequest {
	return &WxaRevertCodeReleaseRequest{}
}

type WxaRevertCodeReleaseRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	AppVersion  int64 `position:"query" name:"app_version" json:"app_version"`
}

type WxaRevertCodeReleaseResponse struct {
}
