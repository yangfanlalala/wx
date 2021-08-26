package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/getcategorybytype.html

const ApiGetCategoriesByType = ""

// GetCategoriesByType 获取不同主体类型的类目
func (client *WeChatClient) GetCategoriesByType(data *GetCategoriesByTypeRequest) (*GetCategoriesByTypeResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiGetCategoriesByType).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		GetCategoriesByTypeResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.GetCategoriesByTypeResponse, nil
}

func (client *WeChatClient) BuildGetCategoriesByTypeRequest() *GetCategoriesByTypeRequest {
	return &GetCategoriesByTypeRequest{}
}

type GetCategoriesByTypeRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	VerifyType  string `position:"body" name:"verify_type" json:"verify_type"`
}

type GetCategoriesByTypeResponse struct {
	CategoriesList struct {
		Categories []WxaCategory `json:"categories"`
	} `json:"categories_list"`
}
