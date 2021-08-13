package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_auditstatus.html

const ApiWxaGetAuditStatus = "https://api.weixin.qq.com/wxa/get_auditstatus"

func (client *WeChatClient) WxaGetAuditStatus() {

}

type WxaGetAuditStatusRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	AuditID     int64  `position:"body" name:"auditid" json:"auditid"`
}

type WxaGetAuditStatusResponse struct {
	ErrorCode    int64  `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
	Status       int64  `json:"status"`
	Reason       string `json:"reason"`
	ScreenShot   string `json:"screen_shot"`
}
