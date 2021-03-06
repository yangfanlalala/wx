	package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/modifyheadimage.html

const ApiModifyHeadImage = "https://api.weixin.qq.com/cgi-bin/account/modifyheadimage"

func (client *WeChatClient) ModifyHeadImage(data *ModifyHeadImageRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiModifyHeadImage).
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

func (client *WeChatClient) BuildModifyHeadImageRequest() *ModifyHeadImageRequest {
	return &ModifyHeadImageRequest{}
}

type ModifyHeadImageRequest struct {
	AccessToken      string `position:"query" name:"access_token" json:"-"`
	HeadImageMediaID string `position:"body" name:"head_img_media_id" json:"head_img_media_id"`
	X1               int64 `position:"body" name:"x1" json:"x1"`
	Y1               int64 `position:"body" name:"y1" json:"y1"`
	X2               int64 `position:"body" name:"x2" json:"x2"`
	Y2               int64 `position:"body" name:"y2" json:"y2"`
}

type ModifyHeadImageResponse struct {
}
