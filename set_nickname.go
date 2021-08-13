package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/setnickname.html

const ApiSetNickname = "https://api.weixin.qq.com/wxa/setnickname"

func (client *WeChatClient) SetNickname() {

}

type SetNicknameRequest struct {
	AccessToken       string `position:"query" name:"access_token" json:"-"`
	Nickname          string `position:"body" name:"nick_name" json:"nickname"`
	IDCard            string `position:"body" name:"id_card" json:"id_card"`
	License           string `position:"body" name:"license" json:"license"`
	NamingOtherStuff1 string `position:"body" name:"naming_other_staff_1" json:"naming_other_stuff_1"`
	NamingOtherStuff2 string `position:"body" name:"naming_other_stuff_2" json:"naming_other_stuff_2"`
	NamingOtherStuff3 string `position:"body" name:"naming_other_stuff_3" json:"naming_other_stuff_3"`
	NamingOtherStuff4 string `position:"body" name:"naming_other_stuff_4" json:"naming_other_stuff_4"`
	NamingOtherStuff5 string `position:"body" name:"naming_other_stuff_5" json:"naming_other_stuff_5"`
}

type SetNicknameResponse struct {
	ErrorCode    int64  `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}
