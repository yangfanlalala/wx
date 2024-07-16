package wx

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

// app document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/record/uploadIcpMedia.html

type ICPMediaType string

const (
	ICPMediaTypeImage = "image"
	ICPMediaTypeVideo = "video"
)

const ApiICPUploadMedia = "https://api.weixin.qq.com/wxa/icp/upload_icp_media"

func (client *WeChatClient) ICPUploadMedia(data *ICPUploadMediaRequest) (*ICPUploadMediaResponse, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	fileWriter, err := writer.CreateFormFile("media", "media.png")
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fileWriter, data.Media)
	if err != nil {
		return nil, err
	}
	writer.WriteField("type", string(data.Type))
	writer.WriteField("certificate_type", strconv.FormatInt(int64(data.CertificateType), 10))
	writer.WriteField("icp_order_field", data.ICPOrderField)
	_ = writer.Close()
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s?access_token=%s", ApiICPUploadMedia, data.AccessToken), payload)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rsp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rsp.Body.Close() }()
	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status not ok, code[%d]", rsp.StatusCode)
	}
	reply := &struct {
		CommonResponse
		ICPUploadMediaResponse
	}{}
	err = jsoniter.NewDecoder(rsp.Body).Decode(reply)
	if err != nil {
		return nil, err
	}
	if err = reply.Error(); err != nil {
		return nil, err
	}
	return &reply.ICPUploadMediaResponse, nil
}

func (client *WeChatClient) BuildICPUploadMediaRequest() *ICPUploadMediaRequest {
	return &ICPUploadMediaRequest{}
}

type ICPUploadMediaRequest struct {
	AccessToken     string       `position:"query" name:"access_token" json:"-"`
	Type            ICPMediaType `position:"body" name:"type" json:"type"`
	CertificateType int32        `position:"body" name:"certificate_type" json:"certificate_type"`
	Media           io.Reader    `position:"body" name:"media" json:"media"`
	ICPOrderField   string       `position:"body" name:"icp_order_field" json:"icp_order_field"`
}

type ICPUploadMediaResponse struct {
	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}
