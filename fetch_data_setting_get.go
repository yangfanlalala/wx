package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/get_fetchdatasetting.html

func (client *WeChatClient) FetchDataSettingGet() {

}

type FetchDataSettingGetRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	Action string `position:"body" name:"action" json:"action"`
}