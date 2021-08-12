package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/speedup_audit.html

const ApiWxaSpeedUpAudit = "https://api.weixin.qq.com/wxa/speedupaudit"

func (client *WeChatClient) WxaSpeedUpAudit() {

}

type WxaSpeedUpAuditRequest struct {
	AccessToken string `position:"query" name:"access_token"`
	AuditID int64 `position:"body" name:"auditid"`
}
