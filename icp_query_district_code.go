package wx

import "net/http"

// app document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/record/queryIcpDistrictCode.html

type ICPDistrictType int32

const (
	ICPDistrictTypeProvince ICPDistrictType = 1
	ICPDistrictTypeCity     ICPDistrictType = 2
	ICPDistrictTypeDistrict ICPDistrictType = 3
)

const ApiICPQueryDistrictCode = "https://api.weixin.qq.com/wxa/icp/query_icp_district_code"

func (client *WeChatClient) ICPQueryDistrictCode(data *ICPQueryDistrictCodeRequest) (*ICPQueryDistrictCodeResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiICPQueryDistrictCode).
		WithMethod(http.MethodGet).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		ICPQueryDistrictCodeResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.ICPQueryDistrictCodeResponse, nil
}

func (client *WeChatClient) BuildICPQueryDistrictCodeRequest() *ICPQueryDistrictCodeRequest {
	return &ICPQueryDistrictCodeRequest{}
}

type ICPQueryDistrictCodeRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type ICPDistrictCodeItem struct {
	Type    ICPDistrictType        `json:"type"`               // 区域类型
	Code    string                 `json:"code"`               // 区域代码
	Name    string                 `json:"name"`               // 区域名称
	SubList []*ICPDistrictCodeItem `json:"sub_list,omitempty"` // 下级区域信息列表
}

type ICPQueryDistrictCodeResponse struct {
	Items []*ICPDistrictCodeItem `json:"items"`
}
