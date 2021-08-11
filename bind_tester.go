package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_AdminManagement/Admin.html

const ApiBindTester = "https://api.weixin.qq.com/wxa/bind_tester"

func (client WeChatClient) BindTester() {

}

type BindTesterRequest struct {
	AccessToken string `name:"access_token"`
	WeChatID string `name:"wechatid"`
}
