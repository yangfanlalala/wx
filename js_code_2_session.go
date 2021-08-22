package wx

import "net/http"

//API Document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/others/WeChat_login.html

const ApiJSCodeToSession = "https://api.weixin.qq.com/sns/component/jscode2session"

func (client *WeChatClient) JSCodeToSession(data *JSCodeToSessionRequest) (*JSCodeToSessionResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiJSCodeToSession).
		WithMethod(http.MethodGet).
		WithData(data)
	rsp := &struct {
		CommonResponse
		JSCodeToSessionResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	return &rsp.JSCodeToSessionResponse, nil
}


func (client *WeChatClient) BuildJSCodeToSessionRequest() *JSCodeToSessionRequest {
	return &JSCodeToSessionRequest{
		GrantType: "authorization_code",
	}
}

type JSCodeToSessionRequest struct {
	AppID string `position:"query" name:"appid" json:"-"`
	JSCode string `position:"query" name:"js_code" json:"-"`
	GrantType string `position:"query" name:"grant_type" json:"-"`
	ComponentAppID string `position:"query" name:"component_appid" json:"-"`
	ComponentAccessToken string `position:"query" name:"component_access_token" json:"-"`
}

type JSCodeToSessionResponse struct {
	OpenID string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID string `json:"unionid"`
}