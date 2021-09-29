package wx

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

func (client *WeChatClient) Plugin() {

}

func (client *WeChatClient) BuildPluginRequest() *PluginRequest {
	return &PluginRequest{}
}

type PluginRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"access_token"`
	Action string `json:"action"`
	PluginID string `json:"plugin_id,omitempty"`
	Reason string `json:"reason,omitempty"`
}

type PluginResponseItem struct {
	Appid string `json:"appid"`
	Status int64 `json:"status"`
	Nickname string `json:"nickname"`
	HeadImageURL string `json:"headimgurl"`
}