package wx

import "net/http"

// API DOCUMENT https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/privacy_config/get_privacy_setting.html

const (
	APIGetPrivacySetting   = "https://api.weixin.qq.com/cgi-bin/component/getprivacysetting"
	PrivacyVerProduct      = 1
	PrivacyVerDevelopement = 2
)

func (client *WeChatClient) GetPrivacySetting(data *GetPrivacySettingRequest) (*GetPrivacySettingResponse, error) {
	req := &CommonRequest{}
	req.WithURL(APIGetPrivacySetting).
		WithContentType("application/json").
		WithMethod(http.MethodPost).
		WithData(data)
	rsp := &struct {
		CommonResponse
		GetPrivacySettingResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.GetPrivacySettingResponse, nil
}

func (client *WeChatClient) BuildGetPrivacySettingRequest() *GetPrivacySettingRequest {
	return &GetPrivacySettingRequest{}
}

type GetPrivacySettingRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	PrivacyVer  int64  `position:"body" name:"privacy_ver" json:"privacy_ver"`
}

type GetPrivacySettingResponse struct {
	CodeExist    int64                 `json:"code_exist"`
	PrivacyList  []string              `json:"privacy_list"`
	SettingList  []*PrivacySettingItem `json:"setting_list"`
	UpdateTime   int64                 `json:"update_time"`
	OwnerSetting PrivacyOwnerSetting   `json:"owner_stting"`
	PrivacyDesc  PrivacySettingDesc    `json:"privacy_desc"`
}
