package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/verify_beta_Mini_Programs/getmpadminauth.html

const ApiGetMpAdminAuth = "https://api.weixin.qq.com/wxa/getmpadminauth" //获取公众号管理员授权

func (client *WeChatClient) GetMpAdminAuth(data *GetMpAdminAuthRequest) (*GetMpAdminAuthResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiGetMpAdminAuth).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		GetMpAdminAuthResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.GetMpAdminAuthResponse, nil
}

func (client *WeChatClient) BuildGetMpAdminAuthRequest() *GetMpAdminAuthRequest {
	return &GetMpAdminAuthRequest{}
}

type GetMpAdminAuthRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	MpAppID     string `position:"body" name:"mp_appid" json:"mp_appid"`
	SameAdmin   int64  `position:"body" name:"same_admin" json:"same_admin"`
}

type GetMpAdminAuthResponse struct {
	AuthorizerURL string `json:"authorizer_url"`
}
