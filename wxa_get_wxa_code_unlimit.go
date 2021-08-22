package wx

import (
	"io"
	"net/http"
)

//API Document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/qrcode/getwxacodeunlimit.html

const ApiGetWxaCodeUnlimit = "https://api.weixin.qq.com/wxa/getwxacodeunlimit"

func (client *WeChatClient) WxaGetWxaCodeUnlimit(data *WxaGetWxaCodeUnlimitRequest) (*io.Reader, error) {
	req := &CommonRequest{}
	req.WithURL(ApiGetWxaCodeUnlimit).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	return nil, nil
}

func (client *WeChatClient) BuildWxaGetWxaCodeUnlimitRequest() *WxaGetWxaCodeUnlimitRequest {
	return &WxaGetWxaCodeUnlimitRequest{}
}

type WxaGetWxaCodeUnlimitRequest struct {
	AccessToken string `position:"query" name:"access_token"  json:"access_token"`
	Scene string `position:"body" name:"scene"  json:"scene"`
	Page string `position:"body" name:"page"  json:"page"`
	Width int64 `position:"body" name:"width"  json:"width"`
	AutoColor bool `position:"body" name:"auto_color"  json:"auto_color"`
	LineColor RGBColor `position:"body" name:"line_color"  json:"line_color"`
	IsHyaline bool `position:"body" name:"is_hyaline"  json:"is_hyaline"`
}

type RGBColor struct {
	R int64 `json:"r"`
	G int64 `json:"g"`
	B int64 `json:"b"`
}

type WxaGetWxaCodeUnlimitResponse io.Reader