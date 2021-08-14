package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/api_get_authorizer_info.html

const ApiGetAuthorizerInfo = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info"

func (client *WeChatClient) GetAuthorizerInfo(data *GetAuthorizerInfoRequest) (*GetAccountBasicInfoResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiGetAuthorizerInfo).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		GetAccountBasicInfoResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.GetAccountBasicInfoResponse, nil
}

type GetAuthorizerInfoRequest struct {
	ComponentAccessToken string `position:"query" name:"component_access_token" json:"-"`
	ComponentAppID       string `position:"body" name:"component_appid" json:"component_appid"`
	AuthorizerAppID      string `position:"body" name:"authorizer_appid" json:"authorizer_appid"`
}

type GetAuthorizerInfoResponse struct {
	AuthorizerInfo    AuthorizerInfo    `json:"authorizer_info"`
	AuthorizationInfo AuthorizationInfo `json:"authorization_info"`
}
