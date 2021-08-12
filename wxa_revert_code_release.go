package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/revertcoderelease.html

const ApiWxaRevertCodeRelease = "https://api.weixin.qq.com/wxa/revertcoderelease"

func (client *WeChatClient) WxaRevertCodeRelease() {

}

type WxaRevertCodeReleaseRequest struct {
	AccessToken string `position:"query" name:"access_token"`
	AppVersion string `position:"query" name:"app_version"`
}