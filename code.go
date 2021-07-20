package wx

const (
	ApiQRCodeJumpGet = "https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumpget" // 获取已设置的二维码规则
	ApiQRCodeJumpDownload = "https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumpdownload" // 获取校验文件名称及内容
	ApiQRCodeJumpAdd = "https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumpadd" // 增加或修改二维码规则
	ApiQRCodeJumpPublish = "https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumppublish" // 发布已设置的二维码规则
	ApiQRCodeJumpDelete = "https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumpdelete" // 删除已设置的二维码规则
	ApiShortURL = "https://api.weixin.qq.com/cgi-bin/shorturl" // 将二维码长链接转成短链接
	ApiGetWxaCodeUnlimit = "https://api.weixin.qq.com/wxa/getwxacodeunlimit" // 获取unlimited小程序码
	ApiGetWxaCode = "https://api.weixin.qq.com/wxa/getwxacode" // 获取小程序码
	ApiCreateWxaQRCode = "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode" // 获取小程序二维码
)

func (cli *WeChatClient) ShortURL() {

}

func (cli *WeChatClient) GetWxaCodeUnlimit() {

}

func (cli *WeChatClient) CreateWxaQRCode() {

}