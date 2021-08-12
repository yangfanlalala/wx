package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_qrcode.html

const ApiWxaGetQRCode = "https://api.weixin.qq.com/wxa/get_qrcode"

func (client *WeChatClient) WxaGetQRCode() {

}

type WxaGetQRCodeRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	Path string `position:"query" name:"path" json:"path"`
}
