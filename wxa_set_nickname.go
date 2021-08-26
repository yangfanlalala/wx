package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/setnickname.html

const ApiSetNickname = "https://api.weixin.qq.com/wxa/setnickname"

func (client *WeChatClient) WxaSetNickname(data *WxaSetNicknameRequest) (*WxaSetNicknameResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiSetNickname).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		WxaSetNicknameResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.WxaSetNicknameResponse, nil
}

func (client *WeChatClient) BuildWxaSetNicknameRequest() *WxaSetNicknameRequest {
	return &WxaSetNicknameRequest{}
}

type WxaSetNicknameRequest struct {
	AccessToken       string `position:"query" name:"access_token" json:"-"`
	Nickname          string `position:"body" name:"nick_name" json:"nick_name"`
	IDCard            string `position:"body" name:"id_card" json:"id_card"`
	License           string `position:"body" name:"license" json:"license"`
	NamingOtherStuff1 string `position:"body" name:"naming_other_staff_1" json:"naming_other_stuff_1"`
	NamingOtherStuff2 string `position:"body" name:"naming_other_stuff_2" json:"naming_other_stuff_2"`
	NamingOtherStuff3 string `position:"body" name:"naming_other_stuff_3" json:"naming_other_stuff_3"`
	NamingOtherStuff4 string `position:"body" name:"naming_other_stuff_4" json:"naming_other_stuff_4"`
	NamingOtherStuff5 string `position:"body" name:"naming_other_stuff_5" json:"naming_other_stuff_5"`
}

type WxaSetNicknameResponse struct {
	Wording string `json:"wording"`
	AuditID int64  `json:"audit_id"`
}
