package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/setwebviewdomain.html

const ApiSetWebViewDomain = "https://api.weixin.qq.com/wxa/setwebviewdomain"
const SetWebViewDomainActionAdd = "add"
const SetWebViewDomainActionDelete = "delete"
const SetWebViewDomainActionSet = "set"
const SetWebViewDomainActionGet = "get"

func (client *WeChatClient) SetWebViewDomain() {

}

type SetWebViewDomainRequest struct {
	AccessToken   string `position:"query" name:"access_token" json:"-"`
	action        string `position:"body" name:"action" json:"action"`
	WebViewDomain string `position:"body" name:"webviewdomain" json:"webviewdomain"`
}

type SetWebViewDomainResponse struct {
CommonResponse
}
