package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_AdminManagement/memberauth.html

const MemberAuthActionGetExperiencer = "get_experiencer"
const ApiMemberAuth = "https://api.weixin.qq.com/wxa/memberauth"

func (client *WeChatClient) MemberAuth() {

}

type MemberAuthRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	Action      string `position:"body" name:"action" json:"action"`
}

type MemberAuthResponse struct {
	ErrorCode    int64  `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
	Members      []struct {
		UserString string `json:"userstr"`
	} `json:"members"`
}
