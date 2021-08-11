package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/release.html

const ApiWxaRelease = "https://api.weixin.qq.com/wxa/release"

func (client WeChatClient) WxaRelease() {

}

type WxaReleaseRequest struct {
	AccessToken string `position:"query" name:"access_token"`
}
