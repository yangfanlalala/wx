package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/modifycategory.html

const ApiModifyCategory = "https://api.weixin.qq.com/cgi-bin/wxopen/modifycategory"

func (client *WeChatClient) ModifyCategory() {

}

type ModifyCategoryRequest struct {
	AccessToken string                    `position:"query" name:"access_token" json:"-"`
	First       string                    `position:"body" name:"first" json:"first"`
	Second      string                    `position:"body" name:"second" json:"second"`
	Certicates  []ModifyCategoryCerticate `position:"body" name:"certicates" json:"certicates"`
}

type ModifyCategoryCerticate struct {
	Key   string `name:"key" json:"key"`
	Value string `name:"value" json:"value"`
}

type ModifyCategoryResponse struct {
CommonResponse
}
