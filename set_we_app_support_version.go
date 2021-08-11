package wx

// app document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/setweappsupportversion.html

const ApiSetWeAppSupportVersion = "https://api.weixin.qq.com/cgi-bin/wxopen/setweappsupportversion"

func (client WeChatClient) SetWeAppSupportVersion() {

}

type SetWeAppSupportVersionRequest struct {
	AccessToken string `position:"query" name:"access_token"`
	Version string `position:"body" name:"version"`
}