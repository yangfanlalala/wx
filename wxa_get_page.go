package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_page.html

const ApiWxaGetPage = "https://api.weixin.qq.com/wxa/get_page"

func (client *WeChatClient) WxaGetPage(data *WxaGetPageRequest) (*WxaGetPageResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiWxaGetPage).
		WithMethod(http.MethodGet).
		WithData(data)
	rsp := &struct {
		CommonResponse
		WxaGetPageResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.WxaGetPageResponse, nil
}

func (client *WeChatClient) BuildWxaGetPageRequest() *WxaGetPageRequest {
	return &WxaGetPageRequest{}
}

type WxaGetPageRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type WxaGetPageResponse struct {
	PageList []string `json:"page_list"`
}
