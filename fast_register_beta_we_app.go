package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/beta_Mini_Programs/fastregister.html

const ApiFastRegisterBetaWeApp = "https://api.weixin.qq.com/wxa/component/fastregisterbetaweapp?access_token=ACCESS_TOKEN" //创建试用小程序

func (client *WeChatClient) FastRegisterBetaWeApp() {

}

type FastRegisterBetaWeAppRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	Name        string `position:"body" name:"name" json:"name"`
	OpenID      string `position:"body" name:"openid" json:"openid"`
}

type FastRegisterBetaWeAppResponse struct {
	UniqueID     string `json:"unique_id"`
	AuthorizeURL string `json:"authorize_url"`
}
