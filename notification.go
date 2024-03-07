package wx

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"encoding/xml"
	"errors"

	"github.com/yangfanlalala/wx/crypt/aes"
)

const NotifyTypeComponentVerifyTicket = "component_verify_ticket"         // 票据推送
const NotifyTypeUnauthorized = "unauthorized"                             // 取消授权
const NotifyTypeUpdateAuthorized = "updateauthorized"                     // 更新授权
const NotifyAuthorized = "authorized"                                     // 授权
const NotifyThirdFastRegisterBetaApp = "notify_third_fastregisterbetaapp" // 快速注册试用小程序
const NotifyThirdFastRegister = "notify_third_fasteregister"              // 快速注册
const NotifyThirdWxaAuth = "notify_3rd_wxa_auth"                          // 认证消息
const NotifyThirdApplyICPFillingResult = "notify_apply_icpfiling_result"  // 备案结果
const NotifyThirdICPFilingVerifyResult = "notify_icpfiling_verify_result" // 人脸核身结果

type NotifySignature struct {
	Signature        string
	Nonce            string
	EncryptType      string
	MessageSignature string
}

type Notification struct {
	AppID                        string `xml:"AppId"`                        // 票据、
	APPID                        string `xml:"appid"`                        // 快速注册、认证
	InfoType                     string `xml:"InfoType"`                     // All
	CreateTime                   int64  `xml:"CreateTime"`                   // All
	ComponentVerifyTicket        string `xml:"ComponentVerifyTicket"`        // 验证票据
	AuthorizerAppID              string `xml:"AuthorizerAppid"`              // 授权
	AuthorizationCode            string `xml:"AuthorizationCode"`            // 授权
	AuthorizationCodeExpiredTime int64  `xml:"AuthorizationCodeExpiredTime"` // 授权
	PreAuthCode                  string `xml:"PreAuthCode"`                  // 授权
	Status                       string `xml:"status"`                       // 快速注册
	FastRegister
	NotificationICPApply
	NotificationICPVerify
	NotificationVerification
}

type FastRegister struct {
	AuthCode string `xml:"auth_code"`
	Msg      string `xml:"msg"`
	Info     struct {
		Name               string `xml:"name"`
		Code               string `xml:"code"`
		CodeType           int64  `xml:"code_type"`
		LegalPersonaWechat string `xml:"legal_persona_wechat"`
		LegalPersonaName   string `xml:"legal_persona_name"`
		ComponentPhone     string `xml:"component_phone"`
		WxUser             string `xml:"wxuser"`
		IDName             string `xml:"idname"`
	} `xml:"info"`
}

type NotificationICPApply struct {
	AuthorizerAPPID string `xml:"authorizer_appid"`
	BeianStatus     int32  `xml:"beian_status"`
}
type NotificationICPVerify struct {
	TaskID      string `xml:"task_id"`
	VerifyAPPID string `xml:"verify_appid"`
	Result      int32  `xml:"result"`
}

type NotificationVerification struct {
	TaskID      string                           `xml:"taskid"`
	TaskStatus  int32                            `xml:"task_status"`
	ApplyStatus int32                            `xml:"apply_status"`
	Dispatch    NotificationVerificationDispatch `xml:"dispatch_info"`
	Message     string                           `xml:"message"`
}
type NotificationVerificationDispatch struct {
	Provider     string `xml:"provider"`
	Contact      string `xml:"contact"`
	DispatchTime int64  `xml:"dispatch_time"`
}

type NotificationProto struct {
	AppID   string `xml:"AppId"`
	Encrypt string `xml:"Encrypt"`
}

type NotificationMessage struct {
	ToUserName   string `json:"ToUserName" xml:"ToUserName"`
	FromUserName string `json:"FromUserName" xml:"FromUserName"`
	CreateTime   int64  `json:"CreateTime" xml:"CreateTime"`
	MessageType  string `json:"MsgType" xml:"MsgType"`
	Content      string `json:"Content" xml:"Content"`
	MessageID    string `json:"MsgId" xml:"MsgId"`
	Event        string `json:"Even" xml:"Event"`
	ContentCheckMessage
}

type ContentCheckMessage struct {
	Appid   string                `json:"appid" xml:"appid"`
	TraceID string                `json:"trace_id" xml:"trace_id"`
	Version uint16                `json:"version" xml:"version"`
	Result  *ContentCheckResult   `json:"result" xml:"result"`
	Detail  []*ContentCheckDetail `json:"detail" xml:"detail"`
}

func (client *WeChatClient) ParseNotification(raw []byte) (*Notification, error) {
	proto := &NotificationProto{}
	if err := xml.Unmarshal(raw, proto); err != nil {
		return nil, err
	}
	if proto.Encrypt == "" {
		return nil, errors.New("encrypt content empty")
	}
	plainTxt, err := aes.WxDecrypt([]byte(proto.Encrypt), []byte(client.options.EncryptKey))
	if err != nil {
		return nil, err
	}
	buff := bytes.NewReader(plainTxt[16:20])
	var length uint32
	if err = binary.Read(buff, binary.BigEndian, &length); err != nil {
		return nil, err
	}
	plainTxt = plainTxt[20 : length+20]
	notify := &Notification{}
	if err = xml.Unmarshal(plainTxt, notify); err != nil {
		return nil, err
	}
	return notify, nil
}

func (client *WeChatClient) ParseMessage(raw []byte, data any) error {
	proto := &NotificationProto{}
	if err := xml.Unmarshal(raw, proto); err != nil {
		return err
	}
	plain, err := aes.WxDecrypt([]byte(proto.Encrypt), []byte(client.options.EncryptKey))
	if err != nil {
		return err
	}
	buff := bytes.NewReader(plain[16:20])
	var leng int32
	if err = binary.Read(buff, binary.BigEndian, &leng); err != nil {
		return err
	}
	plain = plain[20 : leng+20]
	if err = xml.Unmarshal(plain, data); err != nil {
		return err
	}
	return nil
}

func (client *WeChatClient) ParseJSONMessage(raw []byte, data any) error {
	proto := &NotificationProto{}
	if err := json.Unmarshal(raw, proto); err != nil {
		return err
	}
	plain, err := aes.WxDecrypt([]byte(proto.Encrypt), []byte(client.options.EncryptKey))
	if err != nil {
		return err
	}
	buff := bytes.NewReader(plain[16:20])
	var leng int32
	if err = binary.Read(buff, binary.BigEndian, &leng); err != nil {
		return err
	}
	plain = plain[20 : leng+20]
	if err = json.Unmarshal(plain, data); err != nil {
		return err
	}
	return nil
}
