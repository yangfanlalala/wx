package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/verify_beta_Mini_Programs/getmpadminauth.html

const ApiGetMpAdminAuth = "https://api.weixin.qq.com/wxa/getmpadminauth?access_token=ACCESS_TOKEN" //获取公众号管理员授权

func (client *WeChatClient) GetMpAdminAuth() {

}

type GetMpAdminAuthRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	MpAppID     string `position:"mp_appid" name:"mp_appid" json:"mp_appid"`
	SameAdmin   int64  `position:"body" name:"same_admin" json:"same_admin"`
}

type GetMpAdminAuthResponse struct {
	ErrorCode     int64  `json:"errcode"`
	ErrorMessage  string `json:"errmsg"`
	AuthorizerURL string `json:"authorizer_url"`
}
