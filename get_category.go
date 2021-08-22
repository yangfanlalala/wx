package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/getcategory.html

const ApiGetCategory = "https://api.weixin.qq.com/cgi-bin/wxopen/getcategory"

func (client *WeChatClient) GetCategory(data *GetCategoryRequest) (*GetCategoryResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiGetCategory).
		WithMethod(http.MethodGet).
		WithData(data)
	rsp := &struct {
		CommonResponse
		GetCategoryResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.GetCategoryResponse, nil
}

func (client *WeChatClient) BuildGetCategoryRequest() *GetCategoryRequest {
	return &GetCategoryRequest{}
}

type GetCategoryRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type GetCategoryResponse struct {
	Categories []*WxaCategory `json:"categories"`
	Limit int64 `json:"limit"` //?access_token=TOKEN
}
