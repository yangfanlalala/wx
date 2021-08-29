package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/component_verify_ticket_service.html

const ApiStartPushTicket = "https://api.weixin.qq.com/cgi-bin/component/api_start_push_ticket"

func (client *WeChatClient) StartPushTicket(data *StartPushTicketRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiStartPushTicket).
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

type StartPushTicketRequest struct {
	ComponentAppID  string `position:"body" name:"component_appid" json:"component_appid"`
	ComponentSecret string `position:"body" name:"component_secret" json:"component_secret"`
}

type StartPushTicketResponse struct {
}
