package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/get_category.html

const ApiWxaGetCategory = "https://api.weixin.qq.com/wxa/get_category"

func (client *WeChatClient) WxaGetCategory() {

}

type WxaGetCategoryRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type WxaGetCategoryResponse struct {
	CommonResponse
	CategoriesList []WxaCategory `json:"categories_list"`
}
