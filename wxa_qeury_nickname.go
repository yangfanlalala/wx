package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/api_wxa_querynickname.html

const ApiWxaQueryNickname = "https://api.weixin.qq.com/wxa/api_wxa_querynickname"

func (client *WeChatClient) WxaQueryNickname(data *WxaQueryNicknameRequest) (*WxaQueryNicknameResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiWxaQueryNickname).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		WxaQueryNicknameResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil{
		return nil, err
	}
	return &rsp.WxaQueryNicknameResponse, nil
}

type WxaQueryNicknameRequest struct {
	AccessToken string `position:"access_token" name:"access_token" json:"-"`
	AuditID     string `position:"body" name:"audit_id" json:"audit_id"`
}

type WxaQueryNicknameResponse struct {
	Nickname    string `json:"nickname"`
	AuditStatus int64  `json:"audit_stat"`
	FailReason  string `json:"fail_reason"`
	CreateTime  int64  `json:"create_time"`
	AuditTime   int64  `json:"audit_time"`
}
