package wx

import "net/http"

// API DOCUMENT https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/privacy_config/set_privacy_setting.html

const ApiSetPrivacySetting = "https://api.weixin.qq.com/cgi-bin/component/setprivacysetting"

var PrivacyOptions = []*PrivacySettingDesc{
	{
		PrivacyKey:  "UserInfo",
		PrivacyDesc: "用户信息（微信昵称、头像）",
	},
	{
		PrivacyKey:  "Location",
		PrivacyDesc: "位置信息",
	},
	{
		PrivacyKey:  "Address",
		PrivacyDesc: "地址",
	},
	{
		PrivacyKey:  "Invoice",
		PrivacyDesc: "发票信息",
	},
	{
		PrivacyKey:  "RunData",
		PrivacyDesc: "微信运动数据",
	},
	{
		PrivacyKey:  "Record",
		PrivacyDesc: "麦克风",
	},
	{
		PrivacyKey:  "Album",
		PrivacyDesc: "选中的照片或视频信息",
	},
	{
		PrivacyKey:  "Camera",
		PrivacyDesc: "摄像头",
	},
	{
		PrivacyKey:  "PhoneNumber",
		PrivacyDesc: "手机号码",
	},
	{
		PrivacyKey:  "Contact",
		PrivacyDesc: "通讯录（仅写入）权限",
	},
	{
		PrivacyKey:  "DeviceInfo",
		PrivacyDesc: "设备信息",
	},
	{
		PrivacyKey:  "EXIDNumber",
		PrivacyDesc: "身份证号码",
	},
	{
		PrivacyKey:  "EXOrderInfo",
		PrivacyDesc: "订单信息",
	},
	{
		PrivacyKey:  "EXUserPublishContent",
		PrivacyDesc: "发布内容",
	},
	{
		PrivacyKey:  "EXUserFollowAcct",
		PrivacyDesc: "所关注账号",
	},
	{
		PrivacyKey:  "EXUserOpLog",
		PrivacyDesc: "操作日志",
	},
	{
		PrivacyKey:  "AlbumWriteOnly",
		PrivacyDesc: "相册（仅写入）权限",
	},
	{
		PrivacyKey:  "LicensePlate",
		PrivacyDesc: "车牌号",
	},
	{
		PrivacyKey:  "BlueTooth",
		PrivacyDesc: "蓝牙",
	},
	{
		PrivacyKey:  "CalendarWriteOnly",
		PrivacyDesc: "日历（仅写入）权限",
	},
	{
		PrivacyKey:  "Email",
		PrivacyDesc: "邮箱",
	},
	{
		PrivacyKey:  "MessageFile",
		PrivacyDesc: "选中的文件",
	},
}

func (client *WeChatClient) SetPrivacySetting(data *SetPrivacySettingRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiSetPrivacySetting).
		WithContentType("application/json").
		WithMethod(http.MethodPost).
		WithData(data)
	rsp := &CommonResponse{}
	if err := client.DoRequest(req, rsp); err != nil {
		return err
	}
	if err := rsp.Error(); err != nil {
		return err
	}
	return nil
}

func (client *WeChatClient) BuildSetPrivacySettingRequest() *SetPrivacySettingRequest {
	return &SetPrivacySettingRequest{}
}

type SetPrivacySettingRequest struct {
	AccessToken  string                `position:"query" name:"access_token" json:"-"`
	OwnerSetting PrivacyOwnerSetting   `json:"owner_setting"`
	SettingList  []*PrivacySettingItem `json:"setting_list"`
}

type PrivacyOwnerSetting struct {
	ContactEmail         string `json:"contact_email"`
	ContactPhone         string `json:"contact_phone"`
	ContactQQ            string `json:"contact_qq"`
	ContactWeixin        string `json:"contact_weixin"`
	ExtFileMediaID       string `json:"ext_file_media_id"`
	NoticeMethod         string `json:"notice_method"`
	StoreExpireTimestamp string `json:"store_expire_timestamp"`
}

type PrivacySettingItem struct {
	PrivacyKey   string `json:"privacy_key"`
	PrivacyText  string `json:"privacy_text,omitempty"`
	PrivacyLabel string `json:"privacy_label,omitempty"`
}

type PrivacySettingDesc struct {
	PrivacyKey  string `json:"privacy_key"`
	PrivacyDesc string `json:"privacy_desc"`
}

type SetPrivacySettingResponse struct {
}
