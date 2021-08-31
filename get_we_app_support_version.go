package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/getweappsupportversion.html

const ApiGetWeAppSupportVersion = "https://api.weixin.qq.com/cgi-bin/wxopen/getweappsupportversion"

func (client *WeChatClient) GetWeAppSupportVersion(data *GetWeAppSupportVersionRequest) (*GetWeAppSupportVersionResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiGetWeAppSupportVersion).
		WithMethod(http.MethodGet).
		WithData(data)
	rsp := &struct {
		CommonResponse
		GetWeAppSupportVersionResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.GetWeAppSupportVersionResponse, nil
}

func (client WeChatClient) BuildGetWeAppSupportVersionRequest() *GetWeAppSupportVersionRequest {
	return &GetWeAppSupportVersionRequest{}
}

type GetWeAppSupportVersionRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type GetWeAppSupportVersionResponse struct {
	NowVersion string `json:"now_version"`
	UVInfo     struct {
		Items []struct {
			Percentage float64 `json:"percentage"`
			Version    string  `json:"version"`
		} `json:"items"`
	} `json:"uv_info"`
}
