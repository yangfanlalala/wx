package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/wxverify_checknickname.html

const ApiCheckWxVerifyNickname = "https://api.weixin.qq.com/cgi-bin/wxverify/checkwxverifynickname"

func (client *WeChatClient) CheckWxVerifyNickname(data *CheckWxVerifyNicknameRequest)(*CheckWxVerifyNicknameResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiCheckWxVerifyNickname).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		CheckWxVerifyNicknameResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.CheckWxVerifyNicknameResponse, nil
}

type CheckWxVerifyNicknameRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	Nickname    string `position:"body" name:"nick_name" json:"nickname"`
}

type CheckWxVerifyNicknameResponse struct {
	HitCondition bool   `json:"hit_condition"`
	Wording      string `json:"wording"`
}
