package wx

import "net/http"

// app document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/record/cancelIcpfiling.html

const ApiICPCancelFiling = "https://api.weixin.qq.com/wxa/icp/apply_icp_filing"

type ICPCancelType int32

const (
	ICPCancelTypePrincipal ICPCancelType = 1 // 注销主体,
	ICPCancelTypeApp       ICPCancelType = 2 // 注销小程序
	ICPCancelTypeWxa       ICPCancelType = 3 // 注销微信小程序
)

func (client *WeChatClient) ICPCancelFiling(data *ICPCancelFilingRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiICPCancelFiling).
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

func (client *WeChatClient) BuildICPCancelFilingRequest() *ICPCancelFilingRequest {
	return &ICPCancelFilingRequest{}
}

type ICPCancelFilingRequest struct {
	AccessToken string        `position:"query" name:"access_token" json:"-"`
	CancelType  ICPCancelType `position:"cancel_type" name:"cancel_type" json:"cancel_type"`
}

type ICPCancelFilingResponse struct {
}
