package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/query_quota.html

const ApiWxaQueryQuota = "https://api.weixin.qq.com/wxa/queryquota"

func (client *WeChatClient) WxaQueryQuota(data *WxaQueryQuotaRequest) (*WxaQueryQuotaResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiWxaQueryQuota).
		WithMethod(http.MethodGet).
		WithData(data)
	rsp := &struct {
		CommonResponse
		WxaQueryQuotaResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.WxaQueryQuotaResponse, nil
}

type WxaQueryQuotaRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type WxaQueryQuotaResponse struct {
	Rest         int64 `json:"rest"`
	Limit        int64 `json:"limit"`
	SpeedupRest  int64 `json:"speedup_rest"`
	SpeedupLimit int64 `json:"speedup_limit"`
}
