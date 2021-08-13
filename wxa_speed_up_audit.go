package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/speedup_audit.html

const ApiWxaSpeedupAudit = "https://api.weixin.qq.com/wxa/speedupaudit"

func (client *WeChatClient) WxaSpeedupAudit() {

}

type WxaSpeedupAuditRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	AuditID     int64  `position:"body" name:"auditid" json:"audit_id"`
}
