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
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type WxaVersionExp struct {
	ExpTime    int64  `json:"exp_time"`
	ExpVersion string `json:"exp_version"`
	ExpDesc    string `json:"exp_desc"`
}

type WxaVersionRelease struct {
	ReleaseTime    int64  `json:"release_time"`
	ReleaseVersion string `json:"release_version"`
	ReleaseDesc    string `json:"release_desc"`
}
type WxaGetVersionInfoResponse struct {
	ExpInfo     *WxaVersionExp     `json:"exp_info"`
	ReleaseInfo *WxaVersionRelease `json:"release_info"`
}
