package authorization

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"
)

const (
	ApiStartPushTicket     = "https://api.weixin.qq.com/cgi-bin/component/api_start_push_ticket"     // 开启推送
	ApiComponentToken      = "https://api.weixin.qq.com/cgi-bin/component/api_component_token"       // 第三方平台access_token
	ApiCreatePreAuthCode   = "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode"    // 预授权码
	ApiQueryAuth           = "https://api.weixin.qq.com/cgi-bin/component/api_query_auth"            // 使用授权码获取授权信息
	ApiAuthorizerToken     = "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token"      // 获取/刷新接口调用令牌
	ApiGetAuthorizerList   = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_list"   // 拉取所有已授权的帐号信息
	ApiGetAuthorizerOption = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_option" // 获取授权方选项信息
	ApiSetAuthorizerOption = "https://api.weixin.qq.com/cgi-bin/component/api_set_authorizer_option" // 设置授权方选项信息
	ApiGetAuthorizerInfo   = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info"   // 获取授权方的帐号基本信息
)

type baseResponse struct {
	ErrorCode int64 `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}

type Authorizer struct {
	appid              string
	secret             string
	VerifyTicket       string
	VerifyTicketExpire int64
	AccessToken        string
	AccessTokenExpire  int64
	lck                sync.Mutex
	cli                *http.Client
}

func NewAuthorizer(appid, secret string, cli *http.Client) *Authorizer {
	if cli == nil {
		tr := &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 10 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          10,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}
		cli = &http.Client{
			Transport: tr,
			Timeout:   10 * time.Second,
		}
	}
	return &Authorizer{
		appid:  appid,
		secret: secret,
		cli:    cli,
	}
}

// StartPushTicket
// @Summery 开启推送票据
func (a *Authorizer) StartPushTicket() error {
	payload := map[string]interface{} {
		"component_appid": a.appid,
		"component_secret": a.secret,
	}
	data, _ := json.Marshal(payload)
	resp := &baseResponse{}
	err := a.exec(ApiStartPushTicket, http.MethodPost, data, resp)
	if err != nil {
		return err
	}
	if resp.ErrorCode != 0 {
		return fmt.Errorf("request wx service failed, code[%d], message[%s]", resp.ErrorCode, resp.ErrorMessage)
	}
	return nil
}

type Ticket struct {
	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn int64 `json:"expires_in"`
}
// GetAccessToken
// @Summery 从服务端获取accessToken
func (a *Authorizer) GetAccessToken(ticket *Ticket) error {
	payload := map[string]interface{} {
		"component_appid": a.appid,
		"component_appsecret": a.secret,
		"component_verify_ticket": ticket,
	}
	data, _ := json.Marshal(payload)
	resp := &struct {
		baseResponse
		Ticket
	}{}
	err := a.exec(a.buildURL(ApiComponentToken), http.MethodPost, data, resp)
	if err != nil {
		return err
	}
	if resp.ErrorCode != 0 {
		return fmt.Errorf("request wx service failed, code[%d], message[%s]", resp.ErrorCode, resp.ErrorMessage)
	}
	a.AccessToken = resp.ComponentAccessToken
	a.AccessTokenExpire = time.Now().Unix() + resp.ExpiresIn
	return nil
}


type PreAuthorization struct {
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn int64 `json:"expires_in"`
}
// CreatePreAuthCode
// @Summary
func (a *Authorizer) CreatePreAuthCode() (*PreAuthorization, error) {
	payload := map[string]interface{} {
		"component_appid": a.appid,
	}
	data, _ := json.Marshal(payload)
	resp := &struct{
		baseResponse
		PreAuthorization
	}{}
	err := a.exec(a.buildURL(ApiCreatePreAuthCode), http.MethodPost, data, resp)
	if err != nil {
		return nil, err
	}
	if resp.ErrorCode != 0 {
		return nil, fmt.Errorf("request wx service failed code[%d], message[%s]", resp.ErrorCode, resp.ErrorMessage)
	}
	return &resp.PreAuthorization, nil
}

type Authorization struct{
	AuthorizerAppID string `json:"authorizer_appid"`
	AuthorizerAccessToken string `json:"authorizer_access_token"`
	ExpiresIn int64 `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
	FuncInfo []struct{
		FuncScopeCategory struct{
			ID int64 `json:"id"`
		} `json:"funcscope_category"`
	} `json:"func_info"`
}
// QueryAuth
// @Summary
func (a *Authorizer) QueryAuth(ac string) (*Authorization, error) {
	payload := map[string]interface{} {
		"component_appid": a.appid,
		"authorization_code": ac,
	}
	data, _ := json.Marshal(payload)
	resp := &struct {
		baseResponse
		AuthorizationInfo Authorization  `json:"authorization_info"`
	}{}
	err := a.exec(a.buildURL(ApiQueryAuth), http.MethodPost, data, resp)
	if err != nil {
		return nil, err
	}
	if resp.ErrorCode != 0 {
		return nil, fmt.Errorf("request wx service failed, code[%d], message[%s]", resp.ErrorCode, resp.ErrorMessage)
	}
	return &resp.AuthorizationInfo, nil
}

type AuthorizerToken struct {
	AuthorizerAccessToken string `json:"authorizer_access_token"`
	ExpiresIn int64 `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}

func (a *Authorizer) AuthorizerToken(appid, token string) (*AuthorizerToken, error) {
	payload := map[string]interface{} {
		"component_appid": a.appid,
		"authorizer_appid": appid,
		"authorizer_refresh_token": token,
	}
	data, _ := json.Marshal(payload)
	resp := &struct {
		baseResponse
		AuthorizerToken
	}{}
	if err := a.exec(a.buildURL(ApiAuthorizerToken), http.MethodPost, data, resp); err != nil {
		return nil, err
	}
	return &resp.AuthorizerToken, nil
}


func (a *Authorizer) GetAuthorizerList(appid string) {
	payload := &map[string]interface{}{
		"component_appid": a.appid,
		"authorizer_appid": appid,
	}
    data, _ := json.Marshal(payload)
    a.exec(a.buildURL(ApiGetAuthorizerList), http.MethodPost, data, nil)
}

func (a *Authorizer) GetAuthorizerOption(appid string) {
	payload := &map[string]interface{}{
		"component_appid": a.appid,
		"authorizer_appid": appid,
		"option_name": "",
	}
	data, _ := json.Marshal(payload)
	a.exec("", http.MethodPost, data, nil)

}

func (a *Authorizer) SetAuthorizerOption() {

}

func (a *Authorizer) GetAuthorizerInfo() {

}

func (a *Authorizer) exec(url, method string, data []byte, rst interface{}) error {
	req, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	resp, err := a.cli.Do(req)
	if err != nil {
		return err
	}
	defer func() {resp.Body.Close()}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request wx service failed, status code[%d], payload[%s], body[%s]", resp.StatusCode, data, body)
	}
	err = json.Unmarshal(body, rst)
	if err != nil {
		return err
	}
	return nil
}

func (a *Authorizer) buildURL(url string) string {
	return url + "?component_access_token=" + a.AccessToken
}
