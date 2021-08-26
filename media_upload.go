package wx

//Api Document https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
)

const (
	ApiMediaUpload = "https://api.weixin.qq.com/cgi-bin/media/upload?"

	MediaTypeImage = "image"
	MediaTypeVoice = "voice"
	MediaTypeVideo = "video"
	MediaTypeThumb = "thumb"
)

func (client *WeChatClient) MediaUpload(data *MediaUploadRequest) (*MediaUploadResponse, error) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s?access_token=%s&type=%s", ApiMediaUpload, data.AccessToken, data.Type), data.Media)
	if err != nil {
		return nil, err
	}
	rsp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {_=rsp.Body.Close()}()
	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status not ok, code[%d]", rsp.StatusCode)
	}
	reply := &struct {
		CommonResponse
		MediaUploadResponse
	}{}
	err = jsoniter.NewDecoder(rsp.Body).Decode(reply)
	if err != nil {
		return nil, err
	}
	if err = reply.Error(); err != nil {
		return nil, err
	}
	return &reply.MediaUploadResponse, nil
}

func (client *WeChatClient) BuildMediaUploadRequest() *MediaUploadRequest {
	return &MediaUploadRequest{}
}

type MediaUploadRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	Type string `position:"query" name:"type" json:"-"`
	Media io.Reader `position:"body" name:"media" json:"media"`
}

type MediaUploadResponse struct {
	Type string `json:"type"`
	MediaID string `json:"media_id"`
	CreatedAt int64 `json:"created_at"`
}