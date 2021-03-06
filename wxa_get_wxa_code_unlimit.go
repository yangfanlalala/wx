package wx

import (
	"io"
	"net/http"
)

//API Document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/qrcode/getwxacodeunlimit.html

const ApiGetWxaCodeUnlimit = "https://api.weixin.qq.com/wxa/getwxacodeunlimit"

func (client *WeChatClient) WxaGetWxaCodeUnlimit(data *WxaGetWxaCodeUnlimitRequest) (io.ReadCloser, error) {
	req := &CommonRequest{}
	req.WithURL(ApiGetWxaCodeUnlimit).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	return client.DoStream(req)
}

func (client *WeChatClient) BuildWxaGetWxaCodeUnlimitRequest() *WxaGetWxaCodeUnlimitRequest {
	return &WxaGetWxaCodeUnlimitRequest{}
}

type WxaGetWxaCodeUnlimitRequest struct {
	AccessToken string   `position:"query" name:"access_token" json:"-"`
	Scene       string   `position:"body" name:"scene"  json:"scene"`
	Page        string   `position:"body" name:"page"  json:"page"`
	Width       int64    `position:"body" name:"width"  json:"width"`
	AutoColor   bool     `position:"body" name:"auto_color"  json:"auto_color,omitempty"`
	LineColor   *RGBColor `position:"body" name:"line_color"  json:"line_color,omitempty"`
	IsHyaline   bool     `position:"body" name:"is_hyaline"  json:"is_hyaline,omitempty"`
}

type RGBColor struct {
	R int64 `json:"r"`
	G int64 `json:"g"`
	B int64 `json:"b"`
}

type WxaGetWxaCodeUnlimitResponse io.Reader
