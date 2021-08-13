package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/getgrayreleaseplan.html

const ApiWxaGetGrayReleasePlan = "https://api.weixin.qq.com/wxa/getgrayreleaseplan"

func (client *WeChatClient) WxaGetGrayReleasePlan() {

}

type WxaGetGrayReleasePlanRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type WxaGetGrayReleasePlanResponse struct {
	ErrorCode       int64  `json:"errcode"`
	ErrorMessage    string `json:"errmsg"`
	GrayReleasePlan struct {
		Status         int64 `json:"status"`
		CreateTime     int64 `json:"create_time"`
		GrayPercentage int64 `json:"gray_percentage"`
	} `json:"gray_release_plan"`
}
