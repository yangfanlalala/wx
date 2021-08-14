package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_history_version.html

const ApiWxaRevertCodeReleaseHistoryVersion = "https://api.weixin.qq.com/wxa/revertcoderelease"

func (client *WeChatClient) WxaRevertCodeReleaseHistoryVersion() {

}

type WxaRevertCodeReleaseHistoryVersionRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	Action      string `position:"query" name:"action" json:"action"`
}

type WxaRevertCodeReleaseHistoryVersionResponse struct {
	VersionList []AppVersion `json:"version_list"`
}
