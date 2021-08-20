package wx

import (
	"encoding/xml"
	"errors"
)

const NotifyTypeComponentVerifyTicket = "component_verify_ticket"
const NotifyTypeUnauthorized = "unauthorized"
const NotifyTypeUpdateAuthorized = "updateauthorized"
const NotifyAuthorized = "authorized"

type NotifySignature struct {
	Signature string
	Nonce string
	EncryptType string
	MessageSignature string
}

type Notification struct {
	AppID string `xml:"AppId"`
	InfoType string `xml:"InfoType"`
	CreateTime int64 `xml:"CreateTime"`
	ComponentVerifyTicket string `xml:"ComponentVerifyTicket"` //验证票据
	AuthorizerAppID string `xml:"AuthorizerAppid"`
	AuthorizationCode string `xml:"AuthorizationCode"`
	AuthorizationCodeExpiredTime int64 `xml:"AuthorizationCodeExpiredTime"`
	PreAuthCode string `xml:"PreAuthCode"`
}

type NotificationProto struct {
	AppID string `xml:"AppId"`
	Encrypt string `xml:"Encrypt"`
}

func ParseNotification(raw []byte) (*Notification, error) {
	proto := &NotificationProto{}
	if err := xml.Unmarshal(raw, proto); err != nil {
		return nil, err
	}
	if proto.Encrypt == "" {
		return nil, errors.New("encrypt content empty")
	}

}