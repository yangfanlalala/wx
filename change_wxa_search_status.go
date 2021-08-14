package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/changewxasearchstatus.html

const WxaSearchStatusEnable = "1"
const WxaSearchStatusDisabled = "0"
const ApiChangeWxaSearchStatus = "https://api.weixin.qq.com/wxa/changewxasearchstatus"

func (client *WeChatClient) ChangeWxaSearchStatus(data *ChangeWxaSearchStatusRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiChangeWxaSearchStatus).
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

type ChangeWxaSearchStatusRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	Status      string `position:"body" name:"status" json:"status"`
}

type ChangeWxaSearchStatusResponse struct {
}
