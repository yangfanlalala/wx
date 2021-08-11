package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/deletecategory.html

const ApiDeleteCategory = "https://api.weixin.qq.com/cgi-bin/wxopen/deletecategory"

func (client WeChatClient) DeleteCategory() {

}

type DeleteCategoryRequest struct {
	AccessToken string `position:"query" name:"access_token"`
	First int64 `position:"body" name:"first"`
	Second int64 `position:"body" name:"second"`
}
