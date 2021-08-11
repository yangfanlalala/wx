package wx

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ApiSecurityImage = "https://api.weixin.qq.com/wxa/img_sec_check"
	ApiSecurityText  = "https://api.weixin.qq.com/wxa/msg_sec_check"
	ApiSecurityMedia = "https://api.weixin.qq.com/wxa/media_check_async"

	ErrorCodeRisky = 87014
)

func (client *WeChatClient) SecurityImage() {

}

func (client *WeChatClient) SecurityText(token, content string) (bool, error) {
	payload := map[string]interface{}{
		"content": content,
	}
	data, _ := json.Marshal(payload)
	reply := &rsp{}
	if err := client.execute(ApiSecurityText, http.MethodPost, data, reply); err != nil {
		return false, err
	}
	if reply.ErrorCode == ErrorCodeRisky {
		return false, nil
	}
	if reply.ErrorCode != ErrorCodeSuccess {
		return false, fmt.Errorf("request wx service failed, code[%d], message[%s]", reply.ErrorCode, reply.ErrorMessage)
	}
	return true, nil
}

func (client *WeChatClient) SecurityMedia() {

}
