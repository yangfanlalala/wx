package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/change_visitstatus.html

const WxaVisitStatusOpen = "open"
const WxaVisitStatusClose = "close"
const ApiWxaChangeVisitStatus = "https://api.weixin.qq.com/wxa/change_visitstatus"

func (client *WeChatClient) WxaChangeVisitStatus(data *WxaChangeVisitStatusRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiWxaChangeVisitStatus).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &CommonResponse{}
	if err := client.DoRequest(req, rsp); err != nil{
		return err
	}
	if err := rsp.Error(); err != nil {
		return err
	}
	return nil
}

type WxaChangeVisitStatusRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	Action      string `position:"body" name:"action" json:"action"`
}

type WxaChangeVisitStatusResponse struct {
}
