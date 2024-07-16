package wx

import "net/http"

// app document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/record/uploadIcpMedia.html

type ICPMediaType string

const (
	ICPMediaTypeImage = "image"
	ICPMediaTypeVideo = "video"
)

const ApiICPUploadMedia = "https://api.weixin.qq.com/wxa/icp/upload_icp_media"

func (client *WeChatClient) ICPUploadMedia(data *ICPUploadMediaRequest) (*ICPUploadMediaResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiICPUploadMedia).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		ICPUploadMediaResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.ICPUploadMediaResponse, nil
}

func (client *WeChatClient) BuildICPUploadMediaRequest() *ICPUploadMediaRequest {
	return &ICPUploadMediaRequest{}
}

type ICPUploadMediaRequest struct {
	AccessToken     string       `position:"query" name:"access_token" json:"-"`
	Type            ICPMediaType `position:"body" name:"type" json:"type"`
	CertificateType int32        `position:"body" name:"certificate_type" json:"certificate_type"`
	Media           string       `position:"body" name:"media" json:"media"`
	ICPOrderField   string       `position:"body" name:"icp_order_field" json:"icp_order_field"`
}

type ICPUploadMediaResponse struct {
	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}
