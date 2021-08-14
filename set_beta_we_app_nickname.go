package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/beta_Mini_Programs/fastmodify.html

const ApiSetBetaWeAppNickname = "https://api.weixin.qq.com/wxa/setbetaweappnickname?access_token=ACCESS_TOKEN" //设置微信小程序名称

func (client *WeChatClient) SetBetaWeAppNickname() {

}

type SetBetaWeAppNicknameRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	Name        string `position:"Body" name:"name" json:"name"`
}

type SetBetaWeAppNicknameResponse struct {
	Wording string `json:"wording"`
	AuditID int64  `json:"audit_id"`
}
