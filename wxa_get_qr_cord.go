package wx

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
	"strings"
)

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_qrcode.html

const ApiWxaGetQRCode = "https://api.weixin.qq.com/wxa/get_qrcode"

func (client *WeChatClient) WxaGetQRCode(data *WxaGetQRCodeRequest) (io.ReadCloser, error) {
	req := &CommonRequest{}
	req.WithURL(ApiWxaGetQRCode).
		WithMethod(http.MethodGet).
		WithContentType(MineJson).
		WithData(data)
	request, err := req.BuildRequest()
	if err != nil {
		return nil, err
	}
	rsp, err := client.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status not ok, code[%d]", rsp.StatusCode)
	}
	contentType := rsp.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		reply := &CommonResponse{}
		err = jsoniter.NewDecoder(rsp.Body).Decode(reply)
		_ = rsp.Body.Close()
		if err != nil {
			return nil, err
		}
		return nil, reply.Error()
	}
	if strings.HasPrefix(contentType, "image"); err != nil {
		return rsp.Body, nil
	}
	return nil, fmt.Errorf("unknown content type")
}

func (client *WeChatClient) BuildWxaGetQRCodeRequest() *WxaGetQRCodeRequest {
	return &WxaGetQRCodeRequest{}
}

type WxaGetQRCodeRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	Path        string `position:"query" name:"path" json:"path"`
}
