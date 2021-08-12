package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/getallcategories.html

const ApiGetAllCategories = ""

func (client *WeChatClient) GetAllCategories() {

}

type GetAppCategoriesRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}