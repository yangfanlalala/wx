package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/wxverify_checknickname.html

const ApiCheckWxVerifyNickname = "https://api.weixin.qq.com/cgi-bin/wxverify/checkwxverifynickname"

func (client *WeChatClient) CheckWxVerifyNickname() {

}

type CheckWxVerifyNicknameRequest struct {
	AccessToken string `position:"query" name:"access_token"`
	Nickname string `position:"body" name:"nick_name"`
}