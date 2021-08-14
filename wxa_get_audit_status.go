package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_auditstatus.html

const ApiWxaGetAuditStatus = "https://api.weixin.qq.com/wxa/get_auditstatus"

func (client *WeChatClient) WxaGetAuditStatus(data *WxaGetAuditStatusRequest) (*WxaGetAuditStatusResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiWxaGetAuditStatus).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		WxaGetAuditStatusResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.WxaGetAuditStatusResponse, nil
}

type WxaGetAuditStatusRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	AuditID     int64  `position:"body" name:"auditid" json:"auditid"`
}

type WxaGetAuditStatusResponse struct {
	Status     int64  `json:"status"`
	Reason     string `json:"reason"`
	ScreenShot string `json:"screen_shot"`
}
