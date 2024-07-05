package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_auditstatus.html

const ApiWxaGetAuditStatus = "https://api.weixin.qq.com/wxa/get_auditstatus"

var StatusMap = map[int64]struct {
	Status     string
	StatusText string
}{
	0: {
		Status:     "success",
		StatusText: "审核成功",
	},
	1: {
		Status:     "fail",
		StatusText: "审核被拒绝",
	},
	2: {
		Status:     "underway",
		StatusText: "审核中",
	},
	3: {
		Status:     "withdraw",
		StatusText: "已撤回",
	},
	4: {
		Status:     "delay",
		StatusText: "延期审核",
	},
}

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

func (client *WeChatClient) BuildWxaGetAuditStatusRequest() *WxaGetAuditStatusRequest {
	return &WxaGetAuditStatusRequest{}
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
