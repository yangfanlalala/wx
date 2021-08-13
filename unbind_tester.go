package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_AdminManagement/unbind_tester.html

const ApiUnbindTester = "https://api.weixin.qq.com/wxa/unbind_tester"

func (client *WeChatClient) UnbindTester() {

}

type UnbindTesterRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	WeChatID    string `position:"body" name:"wechatid" json:"wechatid"`
	UserString  string `position:"body" name:"userstr" json:"userstr"`
}

type UnbindTesterResponse struct {
	ErrorCode    int64  `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}
