package wx

import "net/http"

// app document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/record/getIcpEntranceInfo.html
var ICPEntranceStatus = map[int32]string{
	2:    "平台审核中",
	3:    "平台审核驳回",
	4:    "管局审核中",
	5:    "管局审核驳回",
	6:    "已备案",
	1024: "未备案",
	1025: "未备案 && 小程序信息未填",
	1026: "未备案 && 小程序类目未填",
	1027: "未备案 && 小程序信息未填 && 小程序类目未填",
	1028: "未备案 && 小程序未认证",
	1029: "未备案 && 小程序信息未填 && 小程序未认证",
	1030: "未备案 && 小程序类目未填 && 小程序未认证",
	1031: "未备案 && 小程序信息未填 && 小程序类目未填 && 小程序未认证",
}

const ApiICPGetEntranceInfo = "https://api.weixin.qq.com/wxa/icp/get_icp_entrance_info"

func (client *WeChatClient) ICPGetEntranceInfo(data *ICPGetEntranceInfoRequest) (*ICPGetEntranceInfoResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiICPGetEntranceInfo).
		WithMethod(http.MethodGet).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		ICPGetEntranceInfoResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.ICPGetEntranceInfoResponse, nil
}

func (client *WeChatClient) BuildICPGetEntranceInfoRequest() *ICPGetEntranceInfoRequest {
	return &ICPGetEntranceInfoRequest{}
}

type ICPGetEntranceInfoRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type ICPEntranceAuditData struct {
	KeyName string `json:"key_name"` // 审核不通过的字段中文名
	Error   string `json:"error"`    // 字段不通过的原因
	Suggest string `json:"suggest"`  // 修改建议
}

type ICPEntranceInfo struct {
	Status      int32                   `json:"status"`       // 备案状态，取值参考备案状态枚举，示例值：1024
	IsCanceling bool                    `json:"is_canceling"` // 是否正在注销备案
	Available   int32                   `json:"available"`    // 备案入口是否对该小程序开放，0：不开放，1：开放。特定情况下入口不会开放，如小程序昵称包含某些关键词时、管局系统不可用时，当备案入口开放时才能提交备案申请。
	AuditData   []*ICPEntranceAuditData `json:"audit_data"`   //驳回原因，备案不通过时返回
}

type ICPGetEntranceInfoResponse struct {
	Info *ICPEntranceInfo `json:"info"`
}
