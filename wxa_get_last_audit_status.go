package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_latest_auditstatus.html

const ApiWxaGetLastAuditStatus = "https://api.weixin.qq.com/wxa/get_latest_auditstatus"

func (client *WeChatClient) WxaGetLastAuditStatus(data *WxaGetLastAuditStatusRequest) (*WxaGetLastAuditStatusResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiWxaGetLastAuditStatus).
		WithMethod(http.MethodGet).
		WithData(data)
	rsp := &struct {
		CommonResponse
		WxaGetLastAuditStatusResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.WxaGetLastAuditStatusResponse, nil
}

func (client *WeChatClient) BuildWxaGetLastAuditStatusRequest() *WxaGetLastAuditStatusRequest {
	return &WxaGetLastAuditStatusRequest{}
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
