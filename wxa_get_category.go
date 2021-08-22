package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/get_category.html

const ApiWxaGetCategory = "https://api.weixin.qq.com/wxa/get_category"

func (client *WeChatClient) WxaGetCategory(data *WxaGetCategoryRequest) (*WxaGetCategoryResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiWxaGetCategory).
		WithMethod(http.MethodGet).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		WxaGetCategoryResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.WxaGetCategoryResponse, nil
}

func (client *WeChatClient) BuildWxaGetCategoryRequest() *WxaGetCategoryRequest {
	return &WxaGetCategoryRequest{}
}

type WxaGetCategoryRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type WxaGetCategoryResponse struct {
	CategoriesList []WxaCategory `json:"categories_list"`
}
