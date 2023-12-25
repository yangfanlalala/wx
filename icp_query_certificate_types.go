package wx

import "net/http"

// app document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/record/queryIcpCertificateTypes.html

const ApiICPQueryCertificateTypes = "https://api.weixin.qq.com/wxa/icp/query_icp_certificate_types"

func (client *WeChatClient) ICPQueryCertificateTypes(data *ICPQueryCertificateTypesRequest) (*ICPQueryCertificateTypesResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiICPQueryCertificateTypes).
		WithMethod(http.MethodGet).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		ICPQueryCertificateTypesResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.ICPQueryCertificateTypesResponse, nil
}

func (client *WeChatClient) BuildICPQueryCertificateTypesRequest() *ICPQueryCertificateTypesRequest {
	return &ICPQueryCertificateTypesRequest{}
}

type ICPQueryCertificateTypesRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type ICPCertificateItem struct {
	Type        int32  `json:"type"`
	SubjectType int32  `json:"subject_type"`
	Name        string `json:"name"`
	Remark      string `json:"remark"`
}

type ICPQueryCertificateTypesResponse struct {
	Items []*ICPCertificateItem `json:"items"`
}
