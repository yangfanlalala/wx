package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_latest_auditstatus.html

const ApiWxaGetLastAuditStatus = "https://api.weixin.qq.com/wxa/get_latest_auditstatus"

func (client *WeChatClient) WxaGetLastAuditStatus() {

}

type WxaGetLastAuditStatusRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type WxaGetLastAuditStatusResponse struct {
	AuditID    int64  `json:"auditid"`
	Status     int64  `json:"status"`
	Reason     string `json:"reason"`
	ScreenShot string `json:"ScreenShot"`
}
