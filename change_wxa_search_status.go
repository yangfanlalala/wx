package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/changewxasearchstatus.html

const WxaSearchStatusEnable = "1"
const WxaSearchStatusDisabled = "0"
const ApiChangeWxaSearchStatus = "https://api.weixin.qq.com/wxa/changewxasearchstatus"

func (client WeChatClient) ChangeWxaSearchStatus() {

}

type ChangeWxaSearchStatusRequest struct {
	AccessToken string `position:"query" name:"access_token"`
	Status string `position:"body" name:"status"`
}