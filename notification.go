package wx

import (
	"bytes"
	"encoding/binary"
	"encoding/xml"
	"errors"
	"github.com/yangfanlalala/wx/crypt/aes"
)

const NotifyTypeComponentVerifyTicket = "component_verify_ticket"
const NotifyTypeUnauthorized = "unauthorized"
const NotifyTypeUpdateAuthorized = "updateauthorized"
const NotifyAuthorized = "authorized"
const NotifyThirdFastRegisterBetaApp = "notify_third_fastregisterbetaapp"
const NotifyThirdFastRegister = "notify_third_fasteregister"

type NotifySignature struct {
	Signature        string
	Nonce            string
	EncryptType      string
	MessageSignature string
}

type Notification struct {
	AppID                        string `xml:"AppId"`
	InfoType                     string `xml:"InfoType"`
	CreateTime                   int64  `xml:"CreateTime"`
	ComponentVerifyTicket        string `xml:"ComponentVerifyTicket"` //验证票据
	AuthorizerAppID              string `xml:"AuthorizerAppid"`
	AuthorizationCode            string `xml:"AuthorizationCode"`
	AuthorizationCodeExpiredTime int64  `xml:"AuthorizationCodeExpiredTime"`
	PreAuthCode                  string `xml:"PreAuthCode"`
	FastRegister
}

type FastRegister struct {
	Appid    string `xml:"appid"`
	Status   int64  `xml:"status"`
	AuthCode string `xml:"auth_code"`
	Msg      string `xml:"msg"`
	Info     struct {
		Name               string `xml:"name"`
		Code               string `xml:"code"`
		CodeType           int64  `xml:"code_type"`
		LegalPersonaWechat string `xml:"legal_personal_wechat"`
		LegalPersonaName   string `xml:"legal_persona_name"`
		ComponentPhone     string `xml:"component_phone"`
	} `xml:"info"`
}

type NotificationProto struct {
	AppID   string `xml:"AppId"`
	Encrypt string `xml:"Encrypt"`
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
