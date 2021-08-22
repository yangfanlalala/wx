package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_history_version.html

const ApiWxaRevertCodeReleaseHistoryVersion = "https://api.weixin.qq.com/wxa/revertcoderelease"

func (client *WeChatClient) WxaRevertCodeReleaseHistoryVersion(data *WxaRevertCodeReleaseHistoryVersionRequest) (*WxaRevertCodeReleaseHistoryVersionResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiWxaRevertCodeReleaseHistoryVersion).
		WithMethod(http.MethodGet).
		WithData(data)
	rsp := &struct {
		CommonResponse
		WxaRevertCodeReleaseHistoryVersionResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.WxaRevertCodeReleaseHistoryVersionResponse, nil
}

func (client *WeChatClient) BuildWxaRevertCodeReleaseHistoryVersionRequest() *WxaRevertCodeReleaseHistoryVersionRequest{
	return &WxaRevertCodeReleaseHistoryVersionRequest{
		Action: "get_history_version",
	}
}

type WxaRevertCodeReleaseHistoryVersionRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	Action      string `position:"query" name:"action" json:"action"`
}

type WxaRevertCodeReleaseHistoryVersionResponse struct {
	VersionList []AppVersion `json:"version_list"`
}
