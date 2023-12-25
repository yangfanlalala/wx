package wx

import "net/http"

// app document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/record/queryIcpSubjectTypes.html

const ApiICPQuerySubjectTypes = "https://api.weixin.qq.com/wxa/icp/query_icp_subject_types"

func (client *WeChatClient) ICPQuerySubjectTypes(data *ICPQuerySubjectTypesRequest) (*ICPQuerySubjectTypesResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiICPQuerySubjectTypes).
		WithMethod(http.MethodGet).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		ICPQuerySubjectTypesResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.ICPQuerySubjectTypesResponse, nil
}

func (client *WeChatClient) BuildICPQuerySubjectTypesRequest() *ICPQuerySubjectTypesRequest {
	return &ICPQuerySubjectTypesRequest{}
}

type ICPQuerySubjectTypesRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type ICPSubjectItem struct {
	Type   int32  `json:"type"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
}

type ICPQuerySubjectTypesResponse struct {
	Items []*ICPSubjectItem `json:"items"`
}
