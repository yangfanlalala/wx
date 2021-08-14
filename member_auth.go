package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_AdminManagement/memberauth.html

const MemberAuthActionGetExperiencer = "get_experiencer"
const ApiMemberAuth = "https://api.weixin.qq.com/wxa/memberauth"

func (client *WeChatClient) MemberAuth(data *MemberAuthRequest) (*MemberAuthResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiMemberAuth).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		MemberAuthResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.MemberAuthResponse, nil
}

type MemberAuthRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	Action      string `position:"body" name:"action" json:"action"`
}

type MemberAuthResponse struct {
	Members []struct {
		UserString string `json:"userstr"`
	} `json:"members"`
}
