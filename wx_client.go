package wx

import (
	"github.com/json-iterator/go"
	"net/http"
)

const (
	MiniProgramStateDeveloper = "developer" // 开发版本
	MiniProgramStateTrial     = "trial"     // 体验版
	MiniProgramStateFormal    = "formal"    // 正式

	ErrorCodeSuccess = 0
)

type Reply struct {
CommonResponse
}

type WeChatClient struct {
	httpProxy  string
	httpsProxy string
	httpClient *http.Client
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

func (client *WeChatClient) getHttpProxy() {

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
	defer func() { _ = rsp.Body.Close() }()
	err = jsoniter.NewDecoder(rsp.Body).Decode(output)
	if err != nil {
		return err
	}
	return nil
}

func NewWeChatClient() *WeChatClient {
	return &WeChatClient{}
}
