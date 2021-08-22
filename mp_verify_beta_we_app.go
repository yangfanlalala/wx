package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/verify_beta_Mini_Programs/mpverifybetaweapp.html

const ApiMpVerifyBetaWeApp = "https://api.weixin.qq.com/wxa/mpverifybetaweapp" //复用公众号主体认证小程序

func (client *WeChatClient) MpVerifyBetaWeApp(data *MpVerifyBetaWeAppRequest) (*MpVerifyBetaWeAppResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiMpVerifyBetaWeApp).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		MpVerifyBetaWeAppResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.MpVerifyBetaWeAppResponse, nil
}

func (client *WeChatClient) BuildMpVerifyBetaWeAppRequest() *MpVerifyBetaWeAppRequest {
	return &MpVerifyBetaWeAppRequest{}
}

type MpVerifyBetaWeAppRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	MpAppID     string `position:"body" name:"mp_appid" json:"mp_appid"`
	Ticket      string `position:"body" name:"ticket" json:"ticket"`
}

type MpVerifyBetaWeAppResponse struct {
	AuthorizerURL string `json:"authorizer_url"`
}
