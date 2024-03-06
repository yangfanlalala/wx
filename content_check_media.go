package wx

import "net/http"

// api Document https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/sec-center/sec-check/mediaCheckAsync.html

const ApiContentCheckMedia = "https://api.weixin.qq.com/wxa/media_check_async"

func (client *WeChatClient) ContentCheckMedia(data *ContentCheckTextRequest) (*ContentCheckMediaResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiContentCheckMedia).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		ContentCheckMediaResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.ContentCheckMediaResponse, nil
}

type ContentCheckMediaRequest struct {
	AccessToken string       `position:"query" name:"access_token" json:"-"`
	MediaURL    string       `position:"body" name:"media_url" json:"media_url"`
	MediaType   string       `position:"body" name:"media_type" json:"media_type"`
	Version     uint16       `position:"body" name:"version" json:"version"`
	Scene       ContentScene `position:"body" name:"scene" json:"scene"`
	OpenID      string       `position:"body" name:"openid" json:"openid"`
}

func (client *WeChatClient) BuildContentCheckMediaRequest() *ContentCheckMediaRequest {
	return &ContentCheckMediaRequest{
		Version: 2,
	}
}

type ContentCheckMediaResponse struct {
	TraceID string `json:"trace_id"`
}
