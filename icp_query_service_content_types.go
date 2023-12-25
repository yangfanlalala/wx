package wx

import "net/http"

// app document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/record/queryIcpServiceContentTypes.html

const ApiICPQueryServiceContentTypes = "https://api.weixin.qq.com/wxa/icp/query_icp_service_content_types"

func (client *WeChatClient) ICPQueryServiceContentTypes(data *ICPQueryServiceContentTypesRequest) (*ICPQueryServiceContentTypesResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiICPQueryServiceContentTypes).
		WithMethod(http.MethodGet).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		ICPQueryServiceContentTypesResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.ICPQueryServiceContentTypesResponse, nil
}

func (client *WeChatClient) BuildICPQueryServiceContentTypesRequest() *ICPQueryServiceContentTypesRequest {
	return &ICPQueryServiceContentTypesRequest{}
}

type ICPQueryServiceContentTypesRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type ICPServiceContentItem struct {
	Type       int32  `json:"type"`
	ParentType int32  `json:"parent_type"`
	Name       string `json:"name"`
	Remark     string `json:"remark"`
}

type ICPQueryServiceContentTypesResponse struct {
	Items []*ICPServiceContentItem `json:"items"`
}
