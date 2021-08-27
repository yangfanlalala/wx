package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/undocodeaudit.html

const ApiUndoCodeAudit = "https://api.weixin.qq.com/wxa/undocodeaudit"

func (client *WeChatClient) UndoCodeAudit(data *UndoCodeAuditRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiUndoCodeAudit).
		WithMethod(http.MethodGet).
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

func (client *WeChatClient) BuildUndoCodeAuditRequest() *UndoCodeAuditRequest {
	return &UndoCodeAuditRequest{}
}

type UndoCodeAuditRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type UndoCodeAuditResponse struct {
}
