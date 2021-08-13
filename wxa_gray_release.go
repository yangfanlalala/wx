package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/grayrelease.html

const ApiWxaGrayRelease = "https://api.weixin.qq.com/wxa/grayrelease"

func (client *WeChatClient) WxaGrayRelease() {

}

type WxaGrayReleaseRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	GrayPercent int64  `position:"query" name:"gray_percent" json:"gray_percent"`
}

type WxaGrayReleaseResponse struct {
	ErrorCode    int64  `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}
