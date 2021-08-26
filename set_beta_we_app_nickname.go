package wx

import (
	"net/http"
)

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/beta_Mini_Programs/fastmodify.html

const ApiSetBetaWeAppNickname = "https://api.weixin.qq.com/wxa/setbetaweappnickname" //设置微信小程序名称

func (client *WeChatClient) SetBetaWeAppNickname(data *SetBetaWeAppNicknameRequest) (*SetBetaWeAppNicknameResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiSetBetaWeAppNickname).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		SetBetaWeAppNicknameResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.SetBetaWeAppNicknameResponse, nil
}

func (client *WeChatClient) BuildSetBataWeAppNicknameRequest() *SetBetaWeAppNicknameRequest {
	return &SetBetaWeAppNicknameRequest{}
}

type SetBetaWeAppNicknameRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	Name        string `position:"Body" name:"name" json:"name"`
}

type SetBetaWeAppNicknameResponse struct {
	Wording string `json:"wording"`
	AuditID int64  `json:"audit_id"`
}
