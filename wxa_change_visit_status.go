package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/change_visitstatus.html

const WxaVisitStatusOpen = "open"
const WxaVisitStatusClose = "close"
const ApiWxaChangeVisitStatus = "https://api.weixin.qq.com/wxa/change_visitstatus"

func (client *WeChatClient) WxaChangeVisitStatus() {

}

type WxaChangeVisitStatusRequest struct {
	AccessToken string `position:"query" name:"access_token"`
	Action string `position:"body" name:"action"`
}