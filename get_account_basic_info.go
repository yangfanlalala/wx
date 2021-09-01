package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/Mini_Program_Information_Settings.html

const ApiGetAccountBasicInfo = "https://api.weixin.qq.com/cgi-bin/account/getaccountbasicinfo"

// GetAccountBasicInfo 获取基本信息
func (client *WeChatClient) GetAccountBasicInfo(data *GetAccountBasicInfoRequest) (*GetAccountBasicInfoResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiGetAccountBasicInfo).
		WithMethod(http.MethodGet).
		WithData(data)
	rsp := &struct {
		CommonResponse
		GetAccountBasicInfoResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.GetAccountBasicInfoResponse, nil
}

func (client *WeChatClient) BuildGetAccountBasicInfoRequest() *GetAccountBasicInfoRequest {
	return &GetAccountBasicInfoRequest{}
}

type GetAccountBasicInfoRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type GetAccountBasicInfoResponse struct {
	AppID             string        `json:"appid"`
	AccountType       int64        	`json:"account_type"`
	PrincipalType     int64         `json:"principal_type"`
	PrincipalName     string        `json:"principal_name"`
	Credential        string        `json:"credential"`
	RealNameStatus    int64         `json:"realname_status"` //实名验证状态
	WxVerifyInfo      VerifyInfo    `json:"wx_verify_info"`
	SignatureInfo     SignatureInfo `json:"signature_info"`
	HeadImageInfo     HeadImageInfo `json:"head_image_info"`
	NicknameInfo      NicknameInfo  `json:"nickname_info"`
	RegisteredCountry int64         `json:"registered_country"`
}
