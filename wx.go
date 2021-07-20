package wx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	MiniProgramStateDeveloper = "developer" // 开发版本
	MiniProgramStateTrial     = "trial"     // 体验版
	MiniProgramStateFormal    = "formal"    // 正式

	ErrorCodeSuccess = 0
)

type rsp struct {
	ErrorCode int64 `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}

type WeChatClient struct {
	httpClient *http.Client
}

func NewWeChatClient() *WeChatClient {
	return &WeChatClient{}
}

func (cli *WeChatClient) execute(url, method string, data []byte, reply interface{}) error {
	req, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	resp, err := cli.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close()}()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		return fmt.Errorf("request wx service failed, status code[%d], message[%s]", resp.StatusCode, body)
	}
	err = json.Unmarshal(body, reply)
	if err != nil {
		return err
	}
	return nil
}

func buildQueryURL(u string, ac string) string {
	return ""
}