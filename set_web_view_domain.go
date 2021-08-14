package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/setwebviewdomain.html

const ApiSetWebViewDomain = "https://api.weixin.qq.com/wxa/setwebviewdomain"
const SetWebViewDomainActionAdd = "add"
const SetWebViewDomainActionDelete = "delete"
const SetWebViewDomainActionSet = "set"
const SetWebViewDomainActionGet = "get"

func (client *WeChatClient) SetWebViewDomain(data *SetWebViewDomainRequest) error {
	req := &CommonRequest{}
 	req.WithURL(ApiSetWebViewDomain).
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

type SetWebViewDomainRequest struct {
	AccessToken   string `position:"query" name:"access_token" json:"-"`
	action        string `position:"body" name:"action" json:"action"`
	WebViewDomain string `position:"body" name:"webviewdomain" json:"webviewdomain"`
}

type SetWebViewDomainResponse struct {
}
