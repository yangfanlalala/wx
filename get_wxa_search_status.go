package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/getwxasearchstatus.html

const ApiGetWxaSearchStatus = "https://api.weixin.qq.com/wxa/getwxasearchstatus"

func (client *WeChatClient) GetWxaSearchStatus(data *GetWxaSearchStatusRequest) (*GetWxaSearchStatusResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiGetWxaSearchStatus).
		WithMethod(http.MethodGet).
		WithData(data)
	rsp := &struct {
		CommonResponse
		GetWxaSearchStatusResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.GetWxaSearchStatusResponse, nil
}

func (client *WeChatClient) BuildGetWxaSearchStatusRequest() *GetWxaSearchStatusRequest {
	return &GetWxaSearchStatusRequest{}
}

type GetWxaSearchStatusRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type GetWxaSearchStatusResponse struct {
	Status int64 `json:"status"`
}
