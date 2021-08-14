package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/addcategory.html

const ApiAddCategory = "https://api.weixin.qq.com/cgi-bin/wxopen/addcategory"

func (client *WeChatClient) AddCategory() {

}

type AddCategoryRequest struct {
	AccessToken string        `position:"query" name:"access_token"`
	Categories  []AddCategory `position:"body" name:"categories" json:"categories"`
}

type AddCategory struct {
	First      int64                  `name:"first"`
	Second     int64                  `name:"second" json:"second"`
	Certicates []AddCategoryCerticate `name:"certicates" json:"certicates"`
}

type AddCategoryCerticate struct {
	Key   string `name:"key" json:"key"`
	Value string `name:"value" json:"value"`
}

type AddCategoryResponse struct {
}
