package wx

import "net/http"

// API Document https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.applyPlugin.html

const (
	ApiPlugin = "https://api.weixin.qq.com/wxa/plugin"
	PluginActionApply = "apply"
	PluginActionList = "list"
	PluginActionUnbind = "unbind"

	PluginStatusApplying = 1
	PluginStatusSucceeded = 2
	PluginStatusRejected = 3
	PluginStatusTimeout = 4
)

func (client *WeChatClient) Plugin(data *PluginRequest) (*PluginResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiPlugin).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		PluginResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.PluginResponse, nil
}

func (client *WeChatClient) BuildPluginRequest() *PluginRequest {
	return &PluginRequest{}
}

type PluginRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	Action string `json:"action"`
	PluginAppid string `json:"plugin_appid,omitempty"`
	Reason string `json:"reason,omitempty"`
}

type PluginResponseItem struct {
	Appid string `json:"appid"`
	Status int64 `json:"status"`
	Nickname string `json:"nickname"`
	HeadImageURL string `json:"headimgurl"`
}

type PluginResponse struct {
	PluginList []*PluginResponseItem `json:"plugin_list"`
}