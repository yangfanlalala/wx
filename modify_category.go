package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/modifycategory.html

const ApiModifyCategory = "https://api.weixin.qq.com/cgi-bin/wxopen/modifycategory"

func (client *WeChatClient) ModifyCategory() {

}

type ModifyCategoryRequest struct {
	AccessToken string `name:"access_token"`
	First string `name:"first"`
	Second string `name:"second"`
	Certicates []ModifyCategoryCerticate `name:"certicates"`
}

type ModifyCategoryCerticate struct {
	Key string `name:"key"`
	Value string `name:"value"`
}