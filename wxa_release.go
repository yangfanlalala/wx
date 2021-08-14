package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/release.html

const ApiWxaRelease = "https://api.weixin.qq.com/wxa/release"

func (client *WeChatClient) WxaRelease(data *WxaReleaseRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiWxaRelease).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
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

type WxaReleaseRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type WxaReleaseResponse struct {
}
