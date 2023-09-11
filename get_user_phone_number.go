package wx

import "net/http"

// api document https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/phone-number/getPhoneNumber.html

const ApiGetUserPhoneNumber = "https://api.weixin.qq.com/wxa/business/getuserphonenumber"

func (client *WeChatClient) GetUserPhoneNumber(data *GetUserPhoneNumberRequest) (*GetUserPhoneNumberResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiGetUserPhoneNumber).
		WithMethod(http.MethodGet).
		WithData(data)
	rsp := &struct {
		CommonResponse
		GetUserPhoneNumberResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.GetUserPhoneNumberResponse, nil
}
func (client WeChatClient) BuildGetUserPhoneNumberRequest() *GetUserPhoneNumberRequest {
	return &GetUserPhoneNumberRequest{}
}

type GetUserPhoneNumberRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	Code        string `json:"code"`
}

type GetUserPhoneNumberResponse struct {
	PhoneInfo *UserPhoneInfo `json:"phone_info"`
}
