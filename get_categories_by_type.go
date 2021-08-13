package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/getcategorybytype.html

const ApiGetCategoriesByType = ""

// GetCategoriesByType 获取不同主体类型的类目
func (client *WeChatClient) GetCategoriesByType() {

}

type GetCategoriesByTypeRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	VerifyType  string `position:"body" name:"verify_type" json:"verify_type"`
}

type GetCategoriesByTypeResponse struct {
	ErrorCode      int64  `json:"errcode"`
	ErrorMessage   string `json:"errmsg"`
	CategoriesList struct {
		Categories []WxaCategory `json:"categories"`
	} `json:"categories_list"`
}
