package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/verify_beta_Mini_Programs/mpverifybetaweapp.html

const ApiMpVerifyBetaWeApp = "https://api.weixin.qq.com/wxa/mpverifybetaweapp?access_token=ACCESS_TOKEN" //复用公众号主体认证小程序

func (client *WeChatClient) MpVerifyBetaWeApp() {

}

type MpVerifyBetaWeAppRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	MpAppID     string `position:"body" name:"mp_appid" json:"mp_appid"`
	Ticket      string `position:"body" name:"ticket" json:"ticket"`
}

type MpVerifyBetaWeAppResponse struct {
	AuthorizerURL string `json:"authorizer_url"`
}
