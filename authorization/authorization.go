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

type atResponse struct {
	baseResponse
	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn int64 `json:"expires_in"`
}
// GetAccessToken
// @Summery 从服务端获取accessToken
func (a *Authorizer) GetAccessToken(ticket string) error {
	payload := map[string]interface{} {
		"component_appid": a.appid,
		"component_appsecret": a.secret,
		"component_verify_ticket": ticket,
	}
	data, _ := json.Marshal(payload)
	resp := &atResponse{}
	err := a.exec(ApiQueryAuth, http.MethodPost, data, resp)
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


type pacResponse struct {
	baseResponse
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn int64 `json:"expires_in"`
}
// CreatePreAuthCode
// @Summary
func (a *Authorizer) CreatePreAuthCode() (string, error) {
	payload := map[string]interface{} {
		"component_access_token": a.AccessToken,
		"component_appid": a.appid,
	}
	data, _ := json.Marshal(payload)
	resp := &pacResponse{}
	err := a.exec(ApiCreatePreAuthCode, http.MethodPost, data, resp)
	if err != nil {
		return "", err
	}
	if resp.ErrorCode != 0 {
		return "", fmt.Errorf("request wx service failed code[%d], message[%s]", resp.ErrorCode, resp.ErrorMessage)
	}
	return resp.PreAuthCode, nil
}

type qaResponse struct {
	baseResponse
	AuthorizationInfo struct{
		AuthorizerAppID string `json:"authorizer_appid"`
		AuthorizerAccessToken string `json:"authorizer_access_token"`
		ExpiresIn int64 `json:"expires_in"`
		AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
		FuncInfo []struct{
			FuncScopeCategory struct{
				ID int64 `json:"id"`
			} `json:"funcscope_category"`
		} `json:"func_info"`
	} `json:"authorization_info"`
}
// QueryAuth
// @Summary
func (a *Authorizer) QueryAuth(ac string) (string, error) {
	payload := map[string]interface{} {
		"component_access_token": "",
		"component_appid": "",
		"authorization_code": "",
	}
	data, _ := json.Marshal(payload)
	resp := &qaResponse{}
	err := a.exec(ApiQueryAuth, http.MethodPost, data, resp)
	if err != nil {
		return "", nil
	}
	if resp.ErrorCode != 0 {
		return "", fmt.Errorf("request wx service failed, code[%d], message[%s]", resp.ErrorCode, resp.ErrorMessage)
	}
	return "", nil
}


func (a *Authorizer) AuthorizerToken() (string, error) {
	return "", nil
}

func (a *Authorizer) GetAuthorizerList() {

}

func (a *Authorizer) GetAuthorizerOption() {

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
