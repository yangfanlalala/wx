package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/deletecategory.html

const ApiDeleteCategory = "https://api.weixin.qq.com/cgi-bin/wxopen/deletecategory"

func (client *WeChatClient) DeleteCategory(data *DeleteCategoryRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiDeleteCategory).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := CommonResponse{}
	if err := client.DoRequest(req, rsp); err != nil {
		return err
	}
	if err := rsp.Error(); err != nil {
		return err
	}
	return nil
}

func (client *WeChatClient) BuildDeleteCategoryRequest() *DeleteCategoryRequest {
	return &DeleteCategoryRequest{}
}

type DeleteCategoryRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	First       int64  `position:"body" name:"first" json:"first"`
	Second      int64  `position:"body" name:"second" json:"second"`
}

type DeleteCategoryResponse struct {
}
