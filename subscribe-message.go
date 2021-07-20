package wx

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ApiTmplGetCategory = "https://api.weixin.qq.com/wxaapi/newtmpl/getcategory" // 获取小程序账号的类目
	ApiTmplGetPubTemplateTitles = "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatetitles" // 获取模板标题列表
	ApiTmplGetPubTemplateKeyWords = "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatekeywords" // 获取模板标题下的关键词库
	ApiTmplAddTemplate = "https://api.weixin.qq.com/wxaapi/newtmpl/addtemplate" // 组合模板并添加至帐号下的个人模板库
	ApiTmplGetTemplate = "https://api.weixin.qq.com/wxaapi/newtmpl/gettemplate" // 获取帐号下的模板列表
	ApiTmplDeleteTemplate = "https://api.weixin.qq.com/wxaapi/newtmpl/deltemplate" // 删除帐号下的某个模板
	ApiTmplSend = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send" // 发送订阅消息

	ErrorCodeTmplOpenIDInvalid = 40003
	ErrorCodeTmplTemplateIDInvalid = 40037
	ErrorCodeTmplUserRefused = 43101
	ErrorCodeTmplParamsInvalid = 47003
	ErrorCodeTmplPagePathInvalid = 41030
)

// 获取小程序账号的类目
type TmplCategory struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
}
func (cli *WeChatClient) TmplGetCategory(token string) ([]*TmplCategory, error) {
	reply := &struct {
		rsp
		Data []*TmplCategory `json:"data"`
	}{}
	if err := cli.execute(ApiTmplGetCategory, http.MethodGet, nil, reply); err != nil {
		return nil, err
	}
	if reply.ErrorCode != ErrorCodeSuccess {
		return nil, fmt.Errorf("request wx service failed, code[%d], message[%s]", reply.ErrorCode, reply.ErrorMessage)
	}
	return reply.Data, nil
}

// 获取模板标题列表
type TmplTitle struct {
	Tid int64 `json:"tid"`
	Title string `json:"title"`
	Type int64 `json:"type"`
	CategoryID int64 `json:"categoryId"`
}
func (cli *WeChatClient) TmplGetPubTemplateTitles(token string, ids []string, start, size int) ([]*TmplTitle, error) {
	reply := &struct {
		rsp
		Count int64 `json:"count"`
		Data []*TmplTitle `json:"data"`
	}{}
	if err := cli.execute(ApiTmplGetPubTemplateTitles, http.MethodGet, nil, reply); err != nil {
		return nil, err
	}
	if reply.ErrorCode != ErrorCodeSuccess {
		return nil, fmt.Errorf("request wx service failed, code[%d], message[%s]", reply.ErrorCode, reply.ErrorMessage)
	}
	return reply.Data, nil
}

// 获取模板标题下的关键词库
type TmplKeyWord struct {
	Kid int64 `json:"kid"`
	Name string `json:"name"`
	Example string `json:"example"`
	Rule string `json:"rule"`
}
func (cli *WeChatClient) TmplGetPubTemplateKeyWords(token, tid string) ([]*TmplKeyWord, error) {
	reply := &struct {
		rsp
		Count int64 `json:"count"`
		Data []*TmplKeyWord `json:"data"`
	}{}
	if err := cli.execute(ApiTmplGetPubTemplateKeyWords, http.MethodGet, nil, reply); err != nil {
		return nil, err
	}
	if reply.ErrorCode != ErrorCodeSuccess {
		return nil, fmt.Errorf("request wx service failed, code[%d], message[%s]", reply.ErrorCode, reply.ErrorMessage)
	}
	return reply.Data, nil
}

// 组合模板并添加至帐号下的个人模板库
func (cli *WeChatClient) TmplAddTemplate(token, tid string, kids []int64, scene string) (string, error) {
	payload := map[string]interface{} {
		"tid": tid,
		"kidList": kids,
		"sceneDesc": scene,
	}
	data, _ := json.Marshal(payload)
	reply := &struct {
		rsp
		PriTmplID string `json:"priTmplId"`
	}{}
	if err := cli.execute(ApiTmplAddTemplate, http.MethodPost, data, reply); err != nil {
		return "", err
	}
	if reply.ErrorCode != ErrorCodeSuccess {
		return "", fmt.Errorf("request wx service failed, code[%d], message[%s]", reply.ErrorCode, reply.ErrorMessage)
	}
	return reply.PriTmplID, nil
}

// 获取帐号下的模板列表
type Template struct {
	PriTmplId string `json:"priTmplId"`
	Title string `json:"title"`
	Content string `json:"content"`
	Example string `json:"example"`
	Type int64 `json:"type"`
}
func (cli *WeChatClient) TmplGetTemplates (token string) ([]*Template, error) {
	reply := &struct {
		rsp
		Data []*Template `json:"data"`
	}{}
	if err := cli.execute(ApiTmplGetTemplate, http.MethodGet, nil, reply); err != nil {
		return nil, err
	}
	if reply.ErrorCode != ErrorCodeSuccess {
		return nil, fmt.Errorf("request wx service failed, code[%d], message[%s]", reply.ErrorCode, reply.ErrorMessage)
	}
	return reply.Data, nil
}
// 删除帐号下的某个模板
func (cli *WeChatClient) Tmpl (token, id string) error {
	payload := map[string] interface{}{"priTmplId": id}
	data, _ := json.Marshal(payload)
	reply := &rsp{}
	if err := cli.execute(ApiTmplDeleteTemplate, http.MethodPost, data, reply); err != nil {
		return err
	}
	if reply.ErrorCode != ErrorCodeSuccess {
		return fmt.Errorf("request wx service failed, code[%d], message[%s]", reply.ErrorCode, reply.ErrorMessage)
	}
	return nil
}

// 发送订阅消息
type TmplMessageData map[string]interface{}
type TmplMessage struct {
	ToUser string `json:"touser"`
	TemplateID string `json:"template_id"`
	Page string `json:"page"`
	Data TmplMessageData `json:"data"`
	MiniProgramState string `json:"miniprogram_state,omitempty"`
	Lang string `json:"lang,omitempty"`
}
func (cli *WeChatClient) TmplSendMessage(token string, tmpl *TmplMessage) error {
	data, _ := json.Marshal(tmpl)
	reply := &rsp{}
	if err := cli.execute(ApiTmplSend, http.MethodPost, data, reply); err != nil {
		return err
	}
	if reply.ErrorCode != ErrorCodeSuccess {
		return fmt.Errorf("request wx service failed, code[%d], message[%s]", reply.ErrorCode, reply.ErrorMessage)
	}
	return nil
}

