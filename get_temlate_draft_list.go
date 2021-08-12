package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/gettemplatedraftlist.html

const ApiGetTemplateDraftList = "https://api.weixin.qq.com/wxa/gettemplatedraftlist?access_token=ACCESS_TOKEN"

func (client *WeChatClient) GetTemplateDraftList() {

}

type GetTemplateDraftListRequest struct {
	AccessToken string `position:"query" name:"access_token"`
}
