package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/getweappsupportversion.html

const ApiGetWeAppSupportVersion = "https://api.weixin.qq.com/cgi-bin/wxopen/getweappsupportversion"

func (client *WeChatClient) GetWeAppSupportVersion() {

}

type GetWeAppSupportVersionRequest struct {
	AccessToken string `name:"access_token" json:"-"`
}

type GetWeAppSupportVersionResponse struct {
	ErrorCode    int64  `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
	NowVersion   string `json:"now_version"`
	UVInfo       struct {
		Items []struct {
			Percentage float64 `json:"percentage"`
			Version    string  `json:"version"`
		} `json:"items"`
	} `json:"uv_info"`
}
