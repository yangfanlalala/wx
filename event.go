package wx

import (
	"bytes"
	"encoding/binary"
	"encoding/xml"
	"errors"
	"github.com/yangfanlalala/wx/crypt/aes"
)

const (
	EventWxaCategoryAudit = "wxa_category_audit"
	EventWxaNicknameAudit = "wxa_nickname_audit"
)

type EventMessage struct {
	AppID        string `xml:"appid"`
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	Ret          string `xml:"ret"`
	Nickname     string `xml:"nickname"`
	Reason       string `xml:"reason"`
	First        int64  `xml:"first"`
	Second       int64  `xml:"second"`
}

type EventMessageProto struct {
	AppID   string `xml:"AppId"`
	ToUserName string `xml:"ToUserName"`
	Encrypt string `xml:"Encrypt"`
}

func (client *WeChatClient) ParseEvent(raw []byte) (*EventMessage, error) {
	proto := &EventMessageProto{}
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
	event := &EventMessage{}
	if err = xml.Unmarshal(plainTxt, event); err != nil {
		return nil, err
	}
	return event, nil
}