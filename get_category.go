package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/getcategory.html

const ApiGetCategory = "https://api.weixin.qq.com/cgi-bin/wxopen/getcategory"

func (client *WeChatClient) GetCategory() {

}

type GetCategoryRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type GetCategoryResponse struct {
}
