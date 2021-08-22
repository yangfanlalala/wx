package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/getgrayreleaseplan.html

const ApiWxaGetGrayReleasePlan = "https://api.weixin.qq.com/wxa/getgrayreleaseplan"

func (client *WeChatClient) WxaGetGrayReleasePlan(data *WxaGetGrayReleasePlanRequest) (*WxaGetGrayReleasePlanResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiWxaGetGrayReleasePlan).
		WithMethod(http.MethodGet).
		WithData(data)
	rsp := &struct {
		CommonResponse
		WxaGetGrayReleasePlanResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.WxaGetGrayReleasePlanResponse, nil
}

func (client *WeChatClient) BuildWxaGetGrayReleasePlanRequest() *WxaGetGrayReleasePlanRequest {
	return &WxaGetGrayReleasePlanRequest{}
}

type WxaGetGrayReleasePlanRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type WxaGetGrayReleasePlanResponse struct {
	GrayReleasePlan struct {
		Status         int64 `json:"status"`
		CreateTime     int64 `json:"create_time"`
		GrayPercentage int64 `json:"gray_percentage"`
	} `json:"gray_release_plan"`
}
