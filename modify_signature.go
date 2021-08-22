package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/modifysignature.html

const ApiModifySignature = "https://api.weixin.qq.com/cgi-bin/account/modifysignature"

func (client *WeChatClient) ModifySignature(data *ModifySignatureRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiModifySignature).
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

func (client *WeChatClient) BuildModifySignatureRequest() *ModifySignatureRequest {
	return &ModifySignatureRequest{}
}

type ModifySignatureRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	Signature   string `position:"body" name:"signature" json:"signature"`
}

type ModifySignatureResponse struct {
}
