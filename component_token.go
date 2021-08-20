package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/component_access_token.html

const ApiComponentToken = "https://api.weixin.qq.com/cgi-bin/component/api_component_token"

func (client *WeChatClient) ComponentToken(data *ComponentTokenRequest) (*ComponentTokenResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiComponentToken).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		ComponentTokenResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.ComponentTokenResponse, nil
}

type ComponentTokenRequest struct {
	ComponentAppID        string `position:"body" name:"component_appid" json:"component_appid"`
	ComponentAppSecret    string `position:"body" name:"component_appsecret" json:"component_appsecret"`
	ComponentVerifyTicket string `position:"body" name:"component_verify_ticket" json:"component_verify_ticket"`
}

func (client *WeChatClient) BuildComponentTokenRequest() *ComponentTokenRequest {
	return &ComponentTokenRequest{
		ComponentAppID:        client.options.AppID,
		ComponentAppSecret:    client.options.AppSecret,
		ComponentVerifyTicket: "",
	}
}

type ComponentTokenResponse struct {
	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn            int64  `json:"expires_in"`
}
