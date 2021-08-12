package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/getgrayreleaseplan.html

const ApiWxaGetGrayReleasePlan = "https://api.weixin.qq.com/wxa/getgrayreleaseplan"

func (client *WeChatClient) WxaGetGrayReleasePlan() {

}

type WxaGetGrayReleasePlanRequest struct {
	AccessToken string `position:"query" name:"access_token"`
}
