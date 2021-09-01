package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/commit.html

const ApiWxaCommit = "https://api.weixin.qq.com/wxa/commit"

func (client *WeChatClient) WxaCommit(data *WxaCommitRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiWxaCommit).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &CommonResponse{}
	if err := client.DoRequest(req, rsp); err != nil {
		return err
	}
	if err := rsp.Error(); err != nil {
		return err
	}
	return nil
}

func (client *WeChatClient) BuildWxaCommitRequest() *WxaCommitRequest {
	return &WxaCommitRequest{}
}

type WxaCommitRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	TemplateID  string `position:"body" name:"template_id" json:"template_id"`
	ExtJson     string `position:"body" name:"ext_json" json:"ext_json"`
	UserVersion string `position:"body" name:"user_version" json:"user_version"`
	UserDesc    string `position:"body" name:"user_desc" json:"user_desc"`
}

type WxaCommitResponse struct {
}
