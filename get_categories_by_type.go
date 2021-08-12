package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/getcategorybytype.html

const ApiGetCategoriesByType = ""

func (client *WeChatClient) GetCategoriesByType() {

}

type GetCategoriesByTypeRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	VerifyType string `position:"body" name:"verify_type" json:"verify_type"`
}