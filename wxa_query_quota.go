package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/query_quota.html

const ApiWxaQueryQuota = "https://api.weixin.qq.com/wxa/queryquota"



func (client *WeChatClient) WxaQueryQuota() {

}

type WxaQueryQuotaRequest struct {
	AccessToken string `position:"query" name:"access_token"`
}