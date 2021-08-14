package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/speedup_audit.html

const ApiWxaSpeedupAudit = "https://api.weixin.qq.com/wxa/speedupaudit"

func (client *WeChatClient) WxaSpeedupAudit(data *WxaSpeedupAuditRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiWxaSpeedupAudit).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &CommonResponse{}
	if err := client.DoRequest(req, rsp); err != nil {
		return err
	}
	if err := rsp.Error(); err != nil {
		return err
	}
	return nil
}

type WxaSpeedupAuditRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	AuditID     int64  `position:"body" name:"auditid" json:"audit_id"`
}

type WxaSpeedupAuditResponse struct {
}
