package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/pre_auth_code.html

const ApiCreatePreAuthCode = "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode"

func (client *WeChatClient) CreatePreAuthCode(data *CreatePreAuthCodeRequest) (*CreatePreAuthCodeResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiCreatePreAuthCode).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		CreatePreAuthCodeResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.CreatePreAuthCodeResponse, nil
}

type CreatePreAuthCodeRequest struct {
	ComponentAccessToken string `position:"query" name:"component_access_token" json:"-"`
	ComponentAppID       string `position:"body" name:"component_appid" json:"component_appid"`
}

func (client *WeChatClient) BuildCreatePreAuthCodeRequest() *CreatePreAuthCodeRequest {
	return &CreatePreAuthCodeRequest{
		ComponentAccessToken: "",
		ComponentAppID:       client.options.AppID,
	}
}

type CreatePreAuthCodeResponse struct {
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int64  `json:"expires_in"` //有效期，单位:秒
}
