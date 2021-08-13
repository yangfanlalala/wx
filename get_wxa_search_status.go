package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/getwxasearchstatus.html

const ApiGetWxaSearchStatus = "https://api.weixin.qq.com/wxa/getwxasearchstatus"

func (client *WeChatClient) GetWxaSearchStatus() {

}

type GetWxaSearchStatusRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type GetWxaSearchStatusResponse struct {
	ErrorCode    int64  `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
	Status       int64  `json:"status"`
}
