package wx

import "net/http"

// app document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/setweappsupportversion.html

const ApiSetWeAppSupportVersion = "https://api.weixin.qq.com/cgi-bin/wxopen/setweappsupportversion"

func (client *WeChatClient) SetWeAppSupportVersion(data *SetWeAppSupportVersionRequest)  error {
	req := &CommonRequest{}
	req.WithURL(ApiSetWeAppSupportVersion).
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

type SetWeAppSupportVersionRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	Version     string `position:"body" name:"version" json:"version"`
}

type SetWeAppSupportVersionResponse struct {
}
