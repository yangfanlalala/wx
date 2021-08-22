package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/authorization_info.html

const ApiQueryAuth = "https://api.weixin.qq.com/cgi-bin/component/api_query_auth"

func (client *WeChatClient) QueryAuth(data *QueryAuthRequest) (*QueryAuthResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiQueryAuth).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		QueryAuthResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil,  err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.QueryAuthResponse, nil
}

func (client *WeChatClient) BuildQueryAuthRequest() *QueryAuthRequest {
	return &QueryAuthRequest{
		ComponentAppID:       client.options.AppID,
	}
}

type QueryAuthRequest struct {
	ComponentAccessToken string `position:"query" name:"component_access_token" json:"-"`
	ComponentAppID       string `position:"body" name:"component_appid" json:"component_appid"`
	AuthorizationCode    string `position:"body" name:"authorization_code" json:"authorization_code"`
}

type QueryAuthResponse struct {
	AuthorizationInfo AuthorizationInfo `json:"authorization_info"`
}
