package wx

import "net/http"

// app document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/record/cancelApplyIcpFiling.html

const ApiICPCancelApplyFiling = "https://api.weixin.qq.com/wxa/icp/cancel_apply_icp_filing"

type ICPCancelApplyType int32

func (client *WeChatClient) ICPCancelApplyFiling(data *ICPCancelApplyFilingRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiICPCancelApplyFiling).
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

func (client *WeChatClient) BuildICPCancelApplyFilingRequest() *ICPCancelApplyFilingRequest {
	return &ICPCancelApplyFilingRequest{}
}

type ICPCancelApplyFilingRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type ICPCancelApplyFilingResponse struct {
}
