package wx

var PrivacyOptions = []*PrivacySettingItem{
	{
		PrivacyKey: "UserInfo",
		PrivacyDesc: "用户信息（微信昵称、头像）",
	},
	{
		PrivacyKey: "Location",
		PrivacyDesc: "位置信息",
	},
	{
		PrivacyKey: "Address",
		PrivacyDesc: "地址",
	},
	{
		PrivacyKey: "Invoice",
		PrivacyDesc: "发票信息",
	},
	{
		PrivacyKey: "",
		PrivacyDesc: "",
	},
	{
		PrivacyKey: "",
		PrivacyDesc: "",
	},
	{
		PrivacyKey: "",
		PrivacyDesc: "",
	},
	{
		PrivacyKey: "",
		PrivacyDesc: "",
	},
	{
		PrivacyKey: "",
		PrivacyDesc: "",
	},
	{
		PrivacyKey: "",
		PrivacyDesc: "",
	},
	{
		PrivacyKey: "",
		PrivacyDesc: "",
	},
	{
		PrivacyKey: "",
		PrivacyDesc: "",
	},
	{
		PrivacyKey: "",
		PrivacyDesc: "",
	},
	{
		PrivacyKey: "",
		PrivacyDesc: "",
	},
	{
		PrivacyKey: "",
		PrivacyDesc: "",
	},
	{
		PrivacyKey: "",
		PrivacyDesc: "",
	},
	{
		PrivacyKey: "",
		PrivacyDesc: "",
	},
	{
		PrivacyKey: "",
		PrivacyDesc: "",
	},
	{
		PrivacyKey: "",
		PrivacyDesc: "",
	},
	{
		PrivacyKey: "",
		PrivacyDesc: "",
	},
	{
		PrivacyKey: "",
		PrivacyDesc: "",
	},
	{
		PrivacyKey: "",
		PrivacyDesc: "",
	},
}

func SetPrivacySetting() {

}

type PrivacyOwnerSetting struct {
	ContactEmail string `json:"contact_email"`
	ContactPhone string `json:"contact_phone"`
	ContactQQ string `json:"contact_qq"`
	ContactWeixin string `json:"contact_weixin"`
	ExtFileMediaID string `json:"ext_file_media_id"`
	NoticeMethod string `json:"notice_method"`
	StoreExpireTimestamp string `json:"store_expire_timestamp"`
}

type PrivacySettingItem struct {
	PrivacyKey string `json:"privacy_key"`
	PrivacyDesc string `json:"privacy_desc,omitempty"`
	PrivacyText string `json:"privacy_text,omitempty"`
}


type SetPrivacySettingRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	OwnerSetting PrivacyOwnerSetting `json:"owner_setting"`
	SettingList []*PrivacySettingItem `json:"setting_list"`
}

type SetPrivacySettingResponse struct {
}