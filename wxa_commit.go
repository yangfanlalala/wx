package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/commit.html

const ApiWxaCommit = "https://api.weixin.qq.com/wxa/commit"

func (client *WeChatClient) WxaCommit() {

}

type WxaCommitRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	TemplateID  string `name:"template_id" json:"template_id"`
	ExtJson     string `name:"ext_json" json:"ext_json"`
	UserVersion string `name:"user_version" json:"user_version"`
	UserDesc    string `name:"user_desc" json:"user_desc"`
}

type WxaCommitResponse struct {
	ErrorCode    int64  `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}
