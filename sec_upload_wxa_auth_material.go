package wx

import "net/http"

// app document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/weapp-wxverify/secwxaapi_wxaauth.html

const ApiSecUploadWxaAuthMaterial = "https://api.weixin.qq.com/wxa/sec/uploadauthmaterial"

func (client *WeChatClient) SecUploadWxaAuthMaterial(data *SecUploadWxaAuthMaterialRequest) (*SecUploadWxaAuthMaterialResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiSecUploadWxaAuthMaterial).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		SecUploadWxaAuthMaterialResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.SecUploadWxaAuthMaterialResponse, nil
}

func (client *WeChatClient) BuildSecUploadWxaAuthMaterialRequest() *SecUploadWxaAuthMaterialRequest {
	return &SecUploadWxaAuthMaterialRequest{}
}

type SecUploadWxaAuthMaterialRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	Media       []byte `position:"body" name:"media" json:"media"`
}

type SecUploadWxaAuthMaterialResponse struct {
	Type    string `json:"type"`
	MediaID string `json:"media_id"`
}
