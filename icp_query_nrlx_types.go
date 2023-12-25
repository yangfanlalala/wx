package wx

import "net/http"

// app document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/record/queryIcpNrlxTypes.html

const ApiICPQueryNRLXTypes = "https://api.weixin.qq.com/wxa/icp/query_icp_nrlx_types"

func (client *WeChatClient) ICPQueryNRLXTypes(data *ICPQueryNRLXTypesRequest) (*ICPQueryNRLXTypesResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiICPQueryNRLXTypes).
		WithMethod(http.MethodGet).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		ICPQueryNRLXTypesResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.ICPQueryNRLXTypesResponse, nil
}

func (client *WeChatClient) BuildICPQueryNRLXTypesRequest() *ICPQueryNRLXTypesRequest {
	return &ICPQueryNRLXTypesRequest{}
}

type ICPQueryNRLXTypesRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type ICPNRLXItem struct {
	Type int32  `json:"type"`
	Name string `json:"name"`
}

type ICPQueryNRLXTypesResponse struct {
	Items []*ICPNRLXItem `json:"items"`
}
