package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/Mini_Program_Information_Settings.html

const ApiGetAccountBasicInfo = "https://api.weixin.qq.com/cgi-bin/account/getaccountbasicinfo"

func (client *WeChatClient) GetAccountBasicInfo() {

}

type GetAccountBasicInfoRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}