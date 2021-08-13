package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_page.html

const ApiWxaGetPage = "https://api.weixin.qq.com/wxa/get_page"

func (client *WeChatClient) WxaGetPage() {

}

type WxaGetPageRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type WxaGetPageResponse struct {
	ErrorCode    int64    `json:"errcode"`
	ErrorMessage string   `json:"errmsg"`
	PageList     []string `json:"page_list"`
}
