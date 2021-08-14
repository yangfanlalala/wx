package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/beta_Mini_Programs/fastregister.html

const ApiFastRegisterBetaWeApp = "https://api.weixin.qq.com/wxa/component/fastregisterbetaweapp" //创建试用小程序

func (client *WeChatClient) FastRegisterBetaWeApp(data *FastRegisterBetaWeAppRequest) (*FastRegisterBetaWeAppResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiFastRegisterBetaWeApp).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		FastRegisterBetaWeAppResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.FastRegisterBetaWeAppResponse, nil
}

type FastRegisterBetaWeAppRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	Name        string `position:"body" name:"name" json:"name"`
	OpenID      string `position:"body" name:"openid" json:"openid"`
}

type FastRegisterBetaWeAppResponse struct {
	UniqueID     string `json:"unique_id"`
	AuthorizeURL string `json:"authorize_url"`
}
