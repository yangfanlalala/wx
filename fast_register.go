package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/fast_registration_of_mini_program.html

const ApiFastRegister = "https://api.weixin.qq.com/cgi-bin/account/fastregister" //复用公众号主体注册小程序

func (client *WeChatClient) FastRegister(data *FastRegisterRequest) (*FastRegisterResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiFastRegister).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		FastRegisterResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.FastRegisterResponse, nil
}

func (client *WeChatClient) BuildFastRegisterRequest() *FastRegisterRequest {
	return &FastRegisterRequest{}
}

type FastRegisterRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	Ticket      string `position:"body" name:"ticket" json:"ticket"`
}

type FastRegisterResponse struct {
	AppID             string `json:"appid"`
	AuthorizationCode string `json:"authorization_code"`
	IsWxVerifySucc    string `json:"is_wx_verify_succ"`
	IsLinkSucc        string `json:"is_link_succ"`
}
