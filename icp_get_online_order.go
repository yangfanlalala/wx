package wx

import "net/http"

// app document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/record/getOnlineIcpOrder.html

const ApiICPGetOnlineOrder = "https://api.weixin.qq.com/wxa/icp/get_online_icp_order"

func (client *WeChatClient) ICPGetOnlineOrder(data *ICPGetOnlineOrderRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiICPGetOnlineOrder).
		WithMethod(http.MethodGet).
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

func (client *WeChatClient) BuildICPGetOnlineOrderRequest() *ICPGetOnlineOrderRequest {
	return &ICPGetOnlineOrderRequest{}
}

type ICPGetOnlineOrderRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type ICPGetOnlineOrderResponse struct {
	ICPSubject *ICPSubject `json:"icp_subject"`
	ICPApplets *ICPApplets `json:"icp_applets"`
}
