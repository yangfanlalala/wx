 package wx

 import (
	 "bytes"
	 "encoding/json"
	 "net/http"
	 "github.com/json-iterator/go"
 )

const (
	MiniProgramStateDeveloper = "developer" // 开发版本
	MiniProgramStateTrial     = "trial"     // 体验版
	MiniProgramStateFormal    = "formal"    // 正式

	ErrorCodeSuccess = 0
)

type Reply struct {
	ErrorCode    int64  `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}

type WeChatClient struct {
	httpProxy string
	httpsProxy string
	httpClient *http.Client
}

//func (client *WeChatClient) SetHTTPSInsecure(isInsecure bool) {
//	client.isInsecure = isInsecure
//}
//
//func (client *WeChatClient) GetHTTPSInsecure() bool {
//	return client.isInsecure
//}

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

func (client *WeChatClient) DoRequest(url, method string, input interface{}, output interface{}) error {
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(input); err != nil {
		return err
	}
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return err
	}
	rsp, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {_ = rsp.Body.Close()}()
	jsoniter.NewDecoder(rsp.Body).Decode(output)
	return nil
}

func NewWeChatClient() *WeChatClient {
	return &WeChatClient{}
}

