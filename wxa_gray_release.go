package wx

import (
	"net/http"
)

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/grayrelease.html

const ApiWxaGrayRelease = "https://api.weixin.qq.com/wxa/grayrelease"

func (client *WeChatClient) WxaGrayRelease(data *WxaGrayReleaseRequest) error {
 	req := &CommonRequest{}
 	req.WithURL(ApiWxaGrayRelease).
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

func (client *WeChatClient) BuildWxaGrayReleaseRequest() *WxaGrayReleaseRequest {
	return &WxaGrayReleaseRequest{}
}

type WxaGrayReleaseRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	GrayPercent int64  `position:"body" name:"gray_percent" json:"gray_percent"`
}

type WxaGrayReleaseResponse struct {
}
