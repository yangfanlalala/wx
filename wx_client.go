package wx

import (
	"fmt"
	"github.com/json-iterator/go"
	"io"
	"net/http"
	"strings"
)

const (
	MiniProgramStateDeveloper = "developer" // 开发版本
	MiniProgramStateTrial     = "trial"     // 体验版
	MiniProgramStateFormal    = "formal"    // 正式

	ErrorCodeSuccess = 0
	MineJson         = "application/json;charset=utf-8"
)

type WeChatOption struct {
	AppID      string
	AppSecret  string
	VerifyKey  string
	EncryptKey string
}

type WeChatClient struct {
	options    *WeChatOption
	httpProxy  string
	httpsProxy string
	httpClient *http.Client
}

func (client *WeChatClient) SetOptions(options *WeChatOption) {
	client.options = options
}

func (client *WeChatClient) GetOptions() *WeChatOption {
	return client.options
}

func (client *WeChatClient) SetHttpProxy(proxy string) {
	client.httpProxy = proxy
}

func (client *WeChatClient) GetHttpProxy() string {
	return client.httpProxy
}

func (client *WeChatClient) SetHttpsProxy(proxy string) {
	client.httpsProxy = proxy
}

func (client *WeChatClient) GetHttpsProxy() string {
	return client.httpsProxy
}

func (client *WeChatClient) SetHttpClient(httpClient *http.Client) {
	client.httpClient = httpClient
}

func (client *WeChatClient) DoRequest(input Request, output Response) error {
	req, err := input.BuildRequest()
	if err != nil {
		return err
	}
	rsp, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf("status not ok, code [%d]", rsp.StatusCode)
	}
	defer func() { _ = rsp.Body.Close() }()
	err = jsoniter.NewDecoder(rsp.Body).Decode(output)
	if err != nil {
		return err
	}
	return nil
}

func (client *WeChatClient) DoStream(input Request) (io.ReadCloser, error) {
	req, err := input.BuildRequest()
	if err != nil {
		return nil, err
	}
	rsp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status not ok, code[%d]", rsp.StatusCode)
	}
	contentType := rsp.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		reply := &CommonResponse{}
		err = jsoniter.NewDecoder(rsp.Body).Decode(reply)
		_ = rsp.Body.Close()
		if err != nil {
			return nil, err
		}
		return nil, reply.Error()
	}
	return rsp.Body, nil
}

func NewWeChatClient() *WeChatClient {
	return &WeChatClient{}
}
