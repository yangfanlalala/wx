package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/revertgrayrelease.html

const ApiRevertGrayRelease = "https://api.weixin.qq.com/wxa/revertgrayrelease"

func (client *WeChatClient) RevertGrayRelease(data *RevertGrayReleaseRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiRevertGrayRelease).
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

func (client *WeChatClient) BuildRevertGrayReleaseRequest() *RevertGrayReleaseRequest {
	return &RevertGrayReleaseRequest{}
}

type RevertGrayReleaseRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type RevertGrayReleaseResponse struct {
}
