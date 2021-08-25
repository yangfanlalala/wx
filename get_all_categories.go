package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/getallcategories.html

const ApiGetAllCategories = "https://api.weixin.qq.com/cgi-bin/wxopen/getallcategories"

func (client *WeChatClient) GetAllCategories(data *GetAppCategoriesRequest) (*GetAppCategoriesResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiGetAllCategories).
		WithMethod(http.MethodGet).
		WithData(data)
	rsp := &struct {
		CommonResponse
		GetAppCategoriesResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.GetAppCategoriesResponse, nil
}

func (client *WeChatClient) BuildGetAppCategoriesRequest() *GetAppCategoriesRequest {
	return &GetAppCategoriesRequest{}
}

type GetAppCategoriesRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type GetAppCategoriesResponse struct {
	CategoriesList struct {
		Categories []Category `json:"categories"`
	} `json:"categories_list"`
}
